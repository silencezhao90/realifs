/*
 * @Author: zhaobo
 * @Date: 2020-07-27 12:32:26
 * @Last Modified by: zhaobo
 * @Last Modified time: 2020-07-27 17:12:23
 */

package minio

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/spf13/viper"
)

func InitConf() *Minio {
	// 初始化配置
	viper.SetConfigFile("../../config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
	}
	newMinio := Minio{
		Endpoint:        viper.GetString("minio.endpoint"),
		AccessKeyID:     viper.GetString("minio.accessKeyID"),
		SecretAccessKey: viper.GetString("minio.secretAccessKey"),
		Secure:          viper.GetBool("minio.secure"),
		BucketName:      viper.GetString("minio.bucketName"),
	}
	newMinio.Init()
	return &newMinio
}

func TestUploadLocalFile(t *testing.T) {
	newMinio := InitConf()
	path, _ := os.Getwd()
	filePath := path + "/minio_test.go"
	remoteFilePath := "tests/filesystem/minio_test/minio_test.go"
	err := newMinio.UploadLocalFile(filePath, remoteFilePath)
	assert.Equal(t, nil, err)
}

func TestCopyFile(t *testing.T) {
	newMinio := InitConf()
	remoteFilePath := "tests/filesystem/minio_test/minio_test.go"
	destFileName := "tests/filesystem/minio_test/copy/minio_test.go"
	err := newMinio.CopyFile(remoteFilePath, destFileName)
	assert.Equal(t, nil, err)
}

func TestDeleteSingleFile(t *testing.T) {
	newMinio := InitConf()
	remoteFilePath := "tests/filesystem/minio_test/minio_test.go"
	err := newMinio.DeleteSingleFile(remoteFilePath)
	assert.Equal(t, nil, err)
}
