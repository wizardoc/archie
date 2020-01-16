package services

import (
	"archie/utils/configer"
	"context"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

func genToken() string {
	qiniuConfig := configer.LoadQiNiuConfig()

	putPolicy := storage.PutPolicy{
		Scope: qiniuConfig.Bucket,
	}
	mac := qbox.NewMac(qiniuConfig.AK, qiniuConfig.SK)

	return putPolicy.UploadToken(mac)
}

func uploadByForm(key string) {
	token := genToken()

	config := storage.Config{
		Zone:          &storage.ZoneHuanan,
		UseHTTPS:      false,
		UseCdnDomains: false,
	}

	formUploader := storage.NewFormUploader(&config)
	ret := storage.PutRet{}

	formUploader.PutFile(context.Background(), &ret, token, key)
}
