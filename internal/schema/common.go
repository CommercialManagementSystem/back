package schema

// QueryOption 查询分页共同参数
type QueryOption struct {
	PageSize int      `form:"pg"`
	Offset   int      `form:"offset"`
	Order    []string `form:"order"`
}

// QueryResponse query 共同返回分页参数
type QueryResponse struct {
	PageSize    int `json:"pageSize"`
	CurrentPage int `json:"current"`
	Count       int `json:"count"`
}

// QueryResponseBody query 统一返回接口
type QueryResponseBody struct {
	QueryResponse
	Data interface{} `json:"data"`
}
