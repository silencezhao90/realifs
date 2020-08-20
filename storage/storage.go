/*
 * @Author: zhaobo
 * @Date: 2020-07-24 14:32:00
 * @Last Modified by: zhaobo
 * @Last Modified time: 2020-07-27 17:18:59
 */

package storage

// Storage is filesystem interface
type Storage interface {
	UploadLocalFile(filePath string, remoteFilePath string) error                                   // 上传本地文件
	DeleteSingleFile(remoteFilePath string) error                                                   // 删除单个文件
	CopyFile(srcFileName string, destFileName string) error                                         // 拷贝文件
	GetUploadPolicy(remoteFilePath string, callbackURL string, callbackBody string) (string, error) // 上传授权
	// GetFileDetailedMeta(remoteFilePath string) (interface{}, error) // 获取文件元信息
}
