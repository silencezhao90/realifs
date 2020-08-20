/*
 * @Author: zhaobo
 * @Date: 2020-07-27 12:32:44
 * @Last Modified by: zhaobo
 * @Last Modified time: 2020-07-27 15:26:12
 */
package aliyun

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/spf13/viper"
)

func InitConf() *Aliyun {
	// 初始化配置
	viper.SetConfigFile("../../config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("config file error: %s\n", err)
	}
	ali := Aliyun{
		AccessKeyID:      viper.GetString("aliyun.accessKeyID"),
		AccessKeySecret:  viper.GetString("aliyun.accessKeySecret"),
		BucketName:       viper.GetString("aliyun.bucketName"),
		ExternalEndpoint: viper.GetString("aliyun.externalEndpoint"),
		InternalEndpoint: viper.GetString("aliyun.InternalEndpoint"),
	}
	ali.Init()
	return &ali
}

func TestUploadLocalFile(t *testing.T) {
	ali := InitConf()
	path, _ := os.Getwd()
	filePath := path + "/aliyun_test.go"
	remoteFilePath := "tests/filesystem/aliyun_test/aliyun_test.go"
	err := ali.UploadLocalFile(filePath, remoteFilePath)
	assert.Equal(t, nil, err)

	// 验证是否上传成功
}

func TestIsFileExist(t *testing.T) {
	ali := InitConf()
	remoteFilePath := "tests/filesystem/aliyun_test/aliyun_test.go"
	exist, _ := ali.IsFileExist(remoteFilePath)
	assert.Equal(t, true, exist)

	remoteFilePath = "tests/filesystem/aliyun_test/aliyun_test2.go"
	exist, _ = ali.IsFileExist(remoteFilePath)
	ali.IsFileExist(remoteFilePath)
	assert.Equal(t, false, exist)
}

func TestCopyFile(t *testing.T) {
	ali := InitConf()
	remoteFilePath := "tests/filesystem/aliyun_test/aliyun_test.go"
	destFileName := "tests/filesystem/aliyun_test/copy/aliyun_test.go"
	err := ali.CopyFile(remoteFilePath, destFileName)
	assert.Equal(t, nil, err)
}

func TestDeleteSingleFile(t *testing.T) {
	ali := InitConf()
	remoteFilePath := "tests/filesystem/aliyun_test/aliyun_test.go"
	err := ali.DeleteSingleFile(remoteFilePath)
	assert.Equal(t, nil, err)
}

func TestGetUploadPolicy(t *testing.T) {
	ali := InitConf()
	callback := map[string]string{"name": "test"}
	callbackJSON, err := json.Marshal(callback)
	if err != nil {
		fmt.Println("Error:", err)
	}

	callbackString := string(callbackJSON)
	policyString, err := ali.GetUploadPolicy("tests/filesystem", "http://127.0.0.1", callbackString)
	if err != nil {
		fmt.Println("Policy error:", err)
	}
	fmt.Println(policyString)
}
