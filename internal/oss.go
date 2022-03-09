package internal

import (
	"github.com/CommercialManagementSystem/back/internal/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func InitOSS() (client *oss.Client, err error) {
	c := config.C.OSS
	client, err = oss.New(c.Endpoint, c.AccessKeyId, c.AccessKeySecret, oss.EnableCRC(c.CRC))
	if err != nil {
		return nil, err
	}
	return client, nil
}
