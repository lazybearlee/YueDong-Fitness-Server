package sysmodel

import "github.com/lazybearlee/yuedong-fitness/global"

// ExaFile file struct, 文件结构体
type ExaFile struct {
	global.BaseModel
	FileName     string
	FileMd5      string
	FilePath     string
	ExaFileChunk []ExaFileChunk
	ChunkTotal   int
	IsFinish     bool
}

// ExaFileChunk file chunk struct, 切片结构体
type ExaFileChunk struct {
	global.BaseModel
	ExaFileID       uint
	FileChunkNumber int
	FileChunkPath   string
}

// ExaFileUploadAndDownload 文件上传下载
type ExaFileUploadAndDownload struct {
	global.BaseModel
	Name string `json:"name" gorm:"comment:文件名"` // 文件名
	Url  string `json:"url" gorm:"comment:文件地址"` // 文件地址
	Tag  string `json:"tag" gorm:"comment:文件标签"` // 文件标签
	Key  string `json:"key" gorm:"comment:编号"`   // 编号
}

func (ExaFileUploadAndDownload) TableName() string {
	return "exa_file_upload_and_downloads"
}
