package oss

import (
	"bytes"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"io/ioutil"
	"path"
)

func DownloadFromOSS(client *oss.Client, bucketName string, fileName string) (*[]byte, error) {
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return nil, err
	}

	// 下载文件到流。
	body, err := bucket.GetObject(fileName)
	if err != nil {
		return nil, err
	}

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			return
		}
	}(body)

	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func JudgeOSSFileExists(client *oss.Client, bucketName string, fileName string) (bool, error) {
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return false, err
	}

	isExist, err := bucket.IsObjectExist(fileName)
	if err != nil {
		return false, err
	}
	return isExist, nil
}

func UploadFileToOSS(client *oss.Client, bucketName string, filepath string, fileName string, file *[]byte) error {
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		return err
	}

	err = bucket.PutObject(path.Join(filepath, fileName), bytes.NewReader(*file))
	if err != nil {
		return err
	}
	return nil
}
