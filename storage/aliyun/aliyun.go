package aliyun

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash"
	"io"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// Aliyun struct
type Aliyun struct {
	AccessKeyID      string // 访问密钥
	AccessKeySecret  string // 密钥Secret
	BucketName       string // 存储空间
	ExternalEndpoint string // 对外访问域名
	InternalEndpoint string // 对内访问域名
	Client           *oss.Client
	Bucket           *oss.Bucket
}

// UploadPolicy 上传授权参数
type UploadPolicy struct {
	AccessKeyID string `json:"OSSAccessKeyId"`
	Policy      string `json:"policy"`
	Signature   string `json:"signature"`
	Host        string `json:"host"`
	Key         string `json:"key"`
	Callback    string `json:"callback"`
}

// Init aliyun storage
func (aliyun *Aliyun) Init() error {
	client, err := oss.New(aliyun.ExternalEndpoint, aliyun.AccessKeyID, aliyun.AccessKeySecret)
	if err != nil {
		fmt.Println("oss New Error:", err)
		return err
	}
	aliyun.Client = client
	bucket, err := client.Bucket(aliyun.BucketName)
	if err != nil {
		fmt.Println("bucket init Error:", err)
		return err
	}
	aliyun.Bucket = bucket
	return nil
}

// UploadLocalFile 上传本地文件
func (aliyun *Aliyun) UploadLocalFile(filePath string, remoteFilePath string) error {
	// 上传本地文件
	if err := aliyun.Bucket.PutObjectFromFile(remoteFilePath, filePath); err != nil {
		return err
	}
	return nil
}

// DeleteSingleFile func
func (aliyun *Aliyun) DeleteSingleFile(remoteFilePath string) error {
	if err := aliyun.Bucket.DeleteObject(remoteFilePath); err != nil {
		return err
	}
	return nil
}

// IsFileExist 判断文件是否存在
func (aliyun *Aliyun) IsFileExist(remoteFilePath string) (bool, error) {
	isExist, err := aliyun.Bucket.IsObjectExist(remoteFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		return false, err
	}

	fmt.Println("isExist:", isExist)
	return isExist, err
}

// CopyFile 复制文件
func (aliyun *Aliyun) CopyFile(srcFileName string, destFileName string) error {
	_, err := aliyun.Bucket.CopyObject(srcFileName, destFileName)
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	return err
}

// GetUploadPolicy 获取上传授权
func (aliyun *Aliyun) GetUploadPolicy(remoteFilePath string, callbackURL string, callbackBody string) (string, error) {
	// callbackBodyEncode := url.QueryEscape(callbackBody)
	callbackMap := map[string]string{
		"callbackUrl":      callbackURL,
		"callbackBody":     "bucket=${bucket}&object=${object}&etag=${etag}&filepath=${object}&size=${size}&mimeType=${mimeType}&imageInfo_height=${imageInfo.height}&imageInfo_width=${imageInfo_width}&imageInfo_format=${imageInfo.format}&" + "data=" + callbackBody,
		"callbackBodyType": "application/x-www-form-urlencoded",
	}
	type Policy struct {
		Expiration string     `json:"expiration"`
		Conditions [][]string `json:"conditions"`
	}
	uploadPolicy := UploadPolicy{}

	// TODO: 获取过期时间
	now := time.Now()
	dt, _ := time.ParseDuration("60s")
	expiredTime := now.Add(dt).Format(time.RFC3339Nano)[:23] + "Z"

	// TODO: 设置条件
	conditions := [][]string{{"starts-with", "$key", remoteFilePath}}
	policy := Policy{Expiration: expiredTime, Conditions: conditions}
	policyJSON, err := json.Marshal(policy)
	if err != nil {
		fmt.Println("Marshal Error:", err)
		return "", err
	}
	policyString := string(policyJSON)
	policyEncode := base64.StdEncoding.EncodeToString([]byte(policyString))

	params, err := json.Marshal(callbackMap)
	if err != nil {
		fmt.Println("Marshal Error:", err)
		return "", err
	}
	callbackString := string(params)
	// TODO: base64 encoding callback params
	callbackEncode := base64.StdEncoding.EncodeToString([]byte(callbackString))

	// TODO: 使用accessKey计算签名
	// signStr := policyEncode
	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(aliyun.AccessKeySecret))
	io.WriteString(h, policyEncode)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))

	uploadPolicy.AccessKeyID = aliyun.AccessKeyID
	uploadPolicy.Policy = policyEncode
	uploadPolicy.Callback = callbackEncode
	uploadPolicy.Host = "https://" + aliyun.BucketName + "." + aliyun.ExternalEndpoint
	uploadPolicy.Key = remoteFilePath
	uploadPolicy.Signature = signedStr

	result, err := json.Marshal(uploadPolicy)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	return string(result), nil
}
