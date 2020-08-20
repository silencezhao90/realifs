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
	"fmt"

	"github.com/spf13/viper"
)

// AliyunConfig config
type AliyunConfig struct {
	AccessKeyID      string // 访问密钥
	AccessKeySecret  string // 密钥Secret
	Bucket           string // 存储空间
	ExternalEndpoint string // 对外访问域名
	InternalEndpoint string // 对内访问域名
}

// MinioCofig config
type MinioCofig struct {
	Endpoint        string // 对象存储服务的URL
	AccessKeyID     string // Access key是唯一标识你的账户的用户ID。
	SecretAccessKey string // Secret key是你账户的密码。
	Secure          bool   // true代表使用HTTPS
}

// ReadConfig 读取配置文件
// func ReadConfig() (*viper.Viper, error) {
// 	viper.SetConfigFile("./config/config.yaml")
// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		fmt.Printf("config file error: %s\n", err)
// 		return nil, err
// 	}
// 	return viper.GetViper(), err
// }

// LoadStorage 初始化配置
func LoadStorage() (storage.Storage, error) {
	// TODO: 读取配置文件，判断使用哪个云端存储，返回对应云端配置信息
	viper.SetConfigFile("./config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
		return nil, err
	}
	c := viper.GetString("default.storage")
	switch c {
	case "aliyun":
		aliStorage := aliyun.Aliyun{
			AccessKeyID:      viper.GetString("aliyun.accessKeyID"),
			AccessKeySecret:  viper.GetString("aliyun.accessKeySecret"),
			BucketName:       viper.GetString("aliyun.bucketName"),
			ExternalEndpoint: viper.GetString("aliyun.externalEndpoint"),
			InternalEndpoint: viper.GetString("aliyun.InternalEndpoint"),
		}
		aliStorage.Init()
		return &aliStorage, nil
	case "minio":
		minioStorage := minio.Minio{
			Endpoint:        viper.GetString("minio.endpoint"),
			AccessKeyID:     viper.GetString("minio.accessKeyID"),
			SecretAccessKey: viper.GetString("minio.secretAccessKey"),
			Secure:          viper.GetBool("minio.secure"),
			BucketName:      viper.GetString("minio.bucketName"),
		}
		minioStorage.Init()
		return &minioStorage, nil
	default:
		return nil, nil
	}

}
