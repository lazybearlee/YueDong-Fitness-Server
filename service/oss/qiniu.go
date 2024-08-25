package oss

import (
	"context"
	"errors"
	"fmt"
	"github.com/lazybearlee/yuedong-fitness/global"
	"mime/multipart"
	"time"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
	"go.uber.org/zap"
)

type Qiniu struct{}

func (*Qiniu) UploadFile(file *multipart.FileHeader) (string, string, error) {
	putPolicy := storage.PutPolicy{Scope: global.FitnessConfig.Qiniu.Bucket}                       // 设置上传策略
	mac := qbox.NewMac(global.FitnessConfig.Qiniu.AccessKey, global.FitnessConfig.Qiniu.SecretKey) // 设置上传凭证
	upToken := putPolicy.UploadToken(mac)                                                          // 生成上传凭证
	cfg := qiniuConfig()                                                                           // 获取配置，包括地区，是否使用https，是否使用cdn加速
	formUploader := storage.NewFormUploader(cfg)                                                   // 创建表单上传的对象
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{Params: map[string]string{"x:name": "github logo"}} // 设置上传额外选项

	f, openError := file.Open()
	if openError != nil {
		global.FitnessLog.Error("function file.Open() failed", zap.Any("err", openError.Error()))

		return "", "", errors.New("function file.Open() failed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)
	putErr := formUploader.Put(context.Background(), &ret, upToken, fileKey, f, file.Size, &putExtra) // 上传文件！
	if putErr != nil {
		global.FitnessLog.Error("function formUploader.Put() failed", zap.Any("err", putErr.Error()))
		return "", "", errors.New("function formUploader.Put() failed, err:" + putErr.Error())
	}
	return global.FitnessConfig.Qiniu.ImgPath + "/" + ret.Key, ret.Key, nil
}

func (*Qiniu) DeleteFile(key string) error {
	mac := qbox.NewMac(global.FitnessConfig.Qiniu.AccessKey, global.FitnessConfig.Qiniu.SecretKey)
	cfg := qiniuConfig()
	bucketManager := storage.NewBucketManager(mac, cfg)
	if err := bucketManager.Delete(global.FitnessConfig.Qiniu.Bucket, key); err != nil {
		global.FitnessLog.Error("function bucketManager.Delete() failed", zap.Any("err", err.Error()))
		return errors.New("function bucketManager.Delete() failed, err:" + err.Error())
	}
	return nil
}

func qiniuConfig() *storage.Config {
	cfg := storage.Config{
		UseHTTPS:      global.FitnessConfig.Qiniu.UseHTTPS,
		UseCdnDomains: global.FitnessConfig.Qiniu.UseCdnDomains,
	}
	switch global.FitnessConfig.Qiniu.Zone { // 根据配置文件进行初始化空间对应的机房
	case "ZoneHuadong":
		cfg.Zone = &storage.ZoneHuadong
	case "ZoneHuabei":
		cfg.Zone = &storage.ZoneHuabei
	case "ZoneHuanan":
		cfg.Zone = &storage.ZoneHuanan
	case "ZoneBeimei":
		cfg.Zone = &storage.ZoneBeimei
	case "ZoneXinjiapo":
		cfg.Zone = &storage.ZoneXinjiapo
	}
	return &cfg
}
