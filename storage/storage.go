/*
 * @Author: zhaobo
 * @Date: 2020-07-24 14:32:00
 * @Last Modified by: zhaobo
 * @Last Modified time: 2020-07-27 17:18:59
 */

package storage

import (
	"filesystem/storage/aliyun"
	"filesystem/storage/minio"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var FS Storage

// Storage is filesystem interface
type Storage interface {
	UploadLocalFile(filePath string, remoteFilePath string) error                                   // 上传本地文件
	DeleteSingleFile(remoteFilePath string) error                                                   // 删除单个文件
	CopyFile(srcFileName string, destFileName string) error                                         // 拷贝文件
	GetUploadPolicy(remoteFilePath string, callbackURL string, callbackBody string) (string, error) // 上传授权
	// GetFileDetailedMeta(remoteFilePath string) (interface{}, error) // 获取文件元信息
}

func Load() {
	// 读取配置文件，判断使用哪个云端存储，返回对应云端配置信息
	c := viper.GetString("default.storage")
	var err error
	switch c {
	case "aliyun":
		FS, err = aliyun.New()
	case "minio":
		FS, err = minio.New()
	default:
		fmt.Printf("%s storage nonsupport", c)
		os.Exit(1)
	}

	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("load with %s storage\n", c)
}
