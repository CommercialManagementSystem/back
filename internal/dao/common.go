package dao

import (
	"fmt"
	"gorm.io/gorm"
)

const (
	NormalAuth  = 1 << 3
	LeaderAuth  = 1 << 2
	CompanyAuth = 1 << 1
	AdminAuth   = 1
)

// CheckAuth 权限校验
func CheckAuth(auth, target int) bool {
	return auth&target != 0
}

// QueryOption dao层查询结构体
type QueryOption struct {
	Offset   int
	PageSize int
	Order    []string
}

// Paginate 分页排序操作
func Paginate(co *QueryOption) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := co.Offset
		if page == 0 {
			page = 1
		}

		pageSize := co.PageSize
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 1
		}

		offset := (page - 1) * pageSize
		db.Offset(offset).Limit(pageSize)

		orderSize := len(co.Order)
		for i := 0; i < orderSize; i++ {
			db = db.Order(co.Order[i])
		}

		return db
	}
}

// getOption 根据传入参数判断是否含有option，有则返回第一个，无则返回默认
func getOption(options ...QueryOption) *QueryOption {
	if len(options) == 0 {
		order := make([]string, 0)
		return &QueryOption{
			Offset:   0,
			PageSize: 10,
			Order:    order,
		}
	}

	return &options[0]
}

type ResOption struct {
	PageSize    int
	CurrentPage int
}

func GetResOption(options ...QueryOption) *ResOption {
	if len(options) == 0 {
		return &ResOption{
			PageSize:    10,
			CurrentPage: 1,
		}
	}
	var option = options[0]
	var pageSize = option.PageSize
	switch {
	case pageSize < 1:
		pageSize = 1
	case pageSize > 100:
		pageSize = 100
	}
	return &ResOption{
		PageSize:    pageSize,
		CurrentPage: option.Offset,
	}
}

func GetLikeSQL(prefix string, suffix string) string {
	return fmt.Sprintf("%s like %% %s %%", prefix, suffix)
}
