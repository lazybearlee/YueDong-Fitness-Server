package sysservice

import (
	"errors"
	"github.com/lazybearlee/yuedong-fitness/global"
	"github.com/lazybearlee/yuedong-fitness/model/common/request"
	sysmodel "github.com/lazybearlee/yuedong-fitness/model/system"
	"github.com/lazybearlee/yuedong-fitness/service/oss"
	"gorm.io/gorm"
	"mime/multipart"
	"strings"
)

type FileService struct{}

func (e *FileService) FindOrCreateFile(fileMd5 string, fileName string, chunkTotal int) (file sysmodel.ExaFile, err error) {
	var cfile sysmodel.ExaFile
	cfile.FileMd5 = fileMd5
	cfile.FileName = fileName
	cfile.ChunkTotal = chunkTotal

	if errors.Is(global.FitnessDb.Where("file_md5 = ? AND is_finish = ?", fileMd5, true).First(&file).Error, gorm.ErrRecordNotFound) {
		err = global.FitnessDb.Where("file_md5 = ? AND file_name = ?", fileMd5, fileName).Preload("ExaFileChunk").FirstOrCreate(&file, cfile).Error
		return file, err
	}
	cfile.IsFinish = true
	cfile.FilePath = file.FilePath
	err = global.FitnessDb.Create(&cfile).Error
	return cfile, err
}

func (e *FileService) CreateFileChunk(id uint, fileChunkPath string, fileChunkNumber int) error {
	var chunk sysmodel.ExaFileChunk
	chunk.FileChunkPath = fileChunkPath
	chunk.ExaFileID = id
	chunk.FileChunkNumber = fileChunkNumber
	err := global.FitnessDb.Create(&chunk).Error
	return err
}

func (e *FileService) DeleteFileChunk(fileMd5 string, filePath string) error {
	var chunks []sysmodel.ExaFileChunk
	var file sysmodel.ExaFile
	err := global.FitnessDb.Where("file_md5 = ? ", fileMd5).First(&file).
		Updates(map[string]interface{}{
			"IsFinish":  true,
			"file_path": filePath,
		}).Error
	if err != nil {
		return err
	}
	err = global.FitnessDb.Where("exa_file_id = ?", file.ID).Delete(&chunks).Unscoped().Error
	return err
}

func (e *FileService) Upload(file sysmodel.ExaFileUploadAndDownload) error {
	return global.FitnessDb.Create(&file).Error
}

func (e *FileService) FindFile(id uint) (sysmodel.ExaFileUploadAndDownload, error) {
	var file sysmodel.ExaFileUploadAndDownload
	err := global.FitnessDb.Where("id = ?", id).First(&file).Error
	return file, err
}

func (e *FileService) DeleteFile(file sysmodel.ExaFileUploadAndDownload) (err error) {
	var fileFromDb sysmodel.ExaFileUploadAndDownload
	fileFromDb, err = e.FindFile(file.ID)
	if err != nil {
		return
	}
	oss_ := oss.NewOSS()
	if err = oss_.DeleteFile(fileFromDb.Key); err != nil {
		return errors.New("文件删除失败")
	}
	err = global.FitnessDb.Where("id = ?", file.ID).Unscoped().Delete(&file).Error
	return err
}

// EditFileName 编辑文件名或者备注
func (e *FileService) EditFileName(file sysmodel.ExaFileUploadAndDownload) (err error) {
	var fileFromDb sysmodel.ExaFileUploadAndDownload
	return global.FitnessDb.Where("id = ?", file.ID).First(&fileFromDb).Update("name", file.Name).Error
}

func (e *FileService) GetFileRecordInfoList(info request.PageInfo) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	keyword := info.Keyword
	db := global.FitnessDb.Model(&sysmodel.ExaFileUploadAndDownload{})
	var fileLists []sysmodel.ExaFileUploadAndDownload
	if len(keyword) > 0 {
		db = db.Where("name LIKE ?", "%"+keyword+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Order("updated_at desc").Find(&fileLists).Error
	return fileLists, total, err
}

func (e *FileService) UploadFile(header *multipart.FileHeader, noSave string) (file sysmodel.ExaFileUploadAndDownload, err error) {
	oss_ := oss.NewOSS()
	filePath, key, uploadErr := oss_.UploadFile(header)
	if uploadErr != nil {
		panic(uploadErr)
	}
	s := strings.Split(header.Filename, ".")
	f := sysmodel.ExaFileUploadAndDownload{
		Url:  filePath,
		Name: header.Filename,
		Tag:  s[len(s)-1],
		Key:  key,
	}
	if noSave == "0" {
		return f, e.Upload(f)
	}
	return f, nil
}
