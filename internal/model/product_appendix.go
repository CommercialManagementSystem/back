package model

import (
	"bytes"
	"context"
	"github.com/CommercialManagementSystem/back/internal/dao"
	"github.com/CommercialManagementSystem/back/internal/entity"
	"github.com/CommercialManagementSystem/back/internal/schema"
	"github.com/CommercialManagementSystem/back/pkg/compress"
	"github.com/CommercialManagementSystem/back/pkg/oss"
	aliyun "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/wire"
	"github.com/jinzhu/copier"
)

var ProductAppendixModelSet = wire.NewSet(wire.Struct(new(ProductAppendixModel), "*"))

type ProductAppendixModel struct {
	OSSClient          *aliyun.Client
	ProductAppendixDao *dao.ProductAppendixDao
}

func (p *ProductAppendixModel) DownloadProductAppendix(ctx context.Context,
	params *schema.DownloadProductAppendixRequestBody) (*bytes.Buffer, error) {
	appendixes, err := p.ProductAppendixDao.QueryByProduct(ctx,
		&entity.ProductAppendix{
			ProductID: uint(params.ProductID),
		},
	)
	if err != nil {
		return nil, err
	}
	type appendixUrl struct {
		Name string
		Url  string
	}
	appendixUrls := new([]appendixUrl)
	err = copier.Copy(appendixUrls, appendixes)
	if err != nil {
		return nil, err
	}

	data := make([]compress.File, 0)
	for _, appendixUrl := range *appendixUrls {
		appendixData, err := oss.DownloadFromOSS(p.OSSClient, "product", appendixUrl.Name)
		if err != nil {
			return nil, err
		}
		data = append(data, compress.File{
			Name: appendixUrl.Name,
			Data: appendixData,
		})
	}

	buffer, err := compress.GenerateZipFromByte(&data)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

func (p *ProductAppendixModel) AddProductAppendix(ctx context.Context, params *schema.AddProductAppendixRequestBody) error {
	data := new([]entity.ProductAppendix)
	err := copier.Copy(data, params.Items)
	if err != nil {
		return err
	}
	err = p.ProductAppendixDao.Create(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProductAppendixModel) DeleteProductAppendix(ctx context.Context, params *schema.DeleteProductAppendixRequestBody) error {
	return p.ProductAppendixDao.Delete(ctx, &params.ProductID)
}
