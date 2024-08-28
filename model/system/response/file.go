package sysresponse

import sysmodel "github.com/lazybearlee/yuedong-fitness/model/system"

type FilePathResponse struct {
	FilePath string `json:"filePath"`
}

type FileResponse struct {
	File sysmodel.ExaFile `json:"file"`
}

type ExaFileResponse struct {
	File sysmodel.ExaFileUploadAndDownload `json:"file"`
}
