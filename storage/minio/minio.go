/*
 * @Author: zhaobo
 * @Date: 2020-07-24 14:39:42
 * @Last Modified by: zhaobo
 * @Last Modified time: 2020-07-27 16:48:42
 */

package minio

import (
	"fmt"
	"os"

	"github.com/minio/minio-go"
)

// Minio struct
type Minio struct {
	Endpoint        string // 对象存储服务的URL
	AccessKeyID     string // Access key是唯一标识你的账户的用户ID。
	SecretAccessKey string // Secret key是你账户的密码。
	Secure          bool   // true代表使用HTTPS
	BucketName      string
	Client          *minio.Client
}

// Init minio初始化
func (newMinio *Minio) Init() error {
	minioClient, err := minio.New(newMinio.Endpoint, newMinio.AccessKeyID, newMinio.SecretAccessKey, newMinio.Secure)
	if err != nil {
		fmt.Println("minio init Error:", err)
		return err
	}
	newMinio.Client = minioClient
	return nil
}

// UploadLocalFile 上传本地文件
func (newMinio *Minio) UploadLocalFile(filePath string, remoteFilePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Open file error:", err)
		return err
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = newMinio.Client.PutObject(newMinio.BucketName, remoteFilePath, file, fileStat.Size(), minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// CopyFile 复制文件
func (newMinio *Minio) CopyFile(srcFileName string, destFileName string) error {
	src := minio.NewSourceInfo(newMinio.BucketName, srcFileName, nil)
	dst, err := minio.NewDestinationInfo(newMinio.BucketName, destFileName, nil, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if err = newMinio.Client.CopyObject(dst, src); err != nil {
		return err
	}
	return nil
}

// DeleteSingleFile 删除单个文件
func (newMinio *Minio) DeleteSingleFile(remoteFilePath string) error {
	if err := newMinio.Client.RemoveObject(newMinio.BucketName, remoteFilePath); err != nil {
		return err
	}
	return nil
}

// GetUploadPolicy 获取上传授权
func (newMinio *Minio) GetUploadPolicy(remoteFilePath string, callbackURL string, callbackBody string) (string, error) {
	// TODO: minio upload policy
	return "", nil
}
