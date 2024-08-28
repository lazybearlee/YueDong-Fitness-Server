package system

import (
	"github.com/lazybearlee/yuedong-fitness/global"
	sysmodel "github.com/lazybearlee/yuedong-fitness/model/system"
)

type FileInitializer struct{}

func (i *FileInitializer) MigrateTable() error {
	return global.FitnessDb.AutoMigrate(
		&sysmodel.ExaFileUploadAndDownload{},
		&sysmodel.ExaFile{},
		&sysmodel.ExaFileChunk{},
	)
}

func (i *FileInitializer) TableCreated() bool {
	return global.FitnessDb.Migrator().HasTable(&sysmodel.ExaFileUploadAndDownload{})
}

func (i *FileInitializer) Name() string {
	return sysmodel.ExaFileUploadAndDownload{}.TableName()
}

func (i *FileInitializer) InitializeData() error {
	return nil
}

func (i *FileInitializer) DataInitialized() bool {
	return true
}
