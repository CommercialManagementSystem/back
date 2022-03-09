package schema

// LoginReqBodySchema 登录结构体
type LoginReqBodySchema struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginResBodySchema 需要返回给前端的数据
type LoginResBodySchema struct {
	UID       string `json:"uid"`
	Authority int    `json:"authority"`
	Token     string `json:"token"`
}
