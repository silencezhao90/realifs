/*
 * @Author: zhaobo
 * @Date: 2020-07-24 14:33:37
 * @Last Modified by: zhaobo
 * @Last Modified time: 2020-07-27 17:19:12
 */

package config

import (
	"filesystem/storage"
	"filesystem/storage/aliyun"
	"filesystem/storage/minio"
	"os"

	"github.com/spf13/viper"
)

// LoadStorage 初始化配置
func LoadStorage() (storage.Storage, error) {
	// TODO: 读取配置文件，判断使用哪个云端存储，返回对应云端配置信息
	c := viper.GetString("default.storage")
	switch c {
	case "aliyun":
		aliStorage := aliyun.Aliyun{
			AccessKeyID:      os.Getenv(viper.GetString("aliyun.accessKeyID")),
			AccessKeySecret:  os.Getenv(viper.GetString("aliyun.accessKeySecret")),
			BucketName:       os.Getenv(viper.GetString("aliyun.bucketName")),
			ExternalEndpoint: os.Getenv(viper.GetString("aliyun.externalEndpoint")),
			InternalEndpoint: os.Getenv(viper.GetString("aliyun.InternalEndpoint")),
		}
		if err := aliStorage.Init(); err != nil {
			return &aliStorage, err
		}
		return &aliStorage, nil
	case "minio":
		minioStorage := minio.Minio{
			Endpoint:        viper.GetString("minio.endpoint"),
			AccessKeyID:     viper.GetString("minio.accessKeyID"),
			SecretAccessKey: viper.GetString("minio.secretAccessKey"),
			Secure:          viper.GetBool("minio.secure"),
			BucketName:      viper.GetString("minio.bucketName"),
		}
		if err := minioStorage.Init(); err != nil {
			return &minioStorage, err
		}
		return &minioStorage, nil
	default:
		return nil, nil
	}

}
