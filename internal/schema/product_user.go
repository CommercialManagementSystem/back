package schema

type AddProductUserRequestBody struct {
	UID  int       `json:"uid"`
	PIDs []Product `json:"pids"`
}

type Product struct {
	PID int `json:"pid"`
}

type DeleteProductUserRequestBody struct {
	UID int `json:"uid"`
	PID int `json:"pid"`
}

type QueryProductUserRequestBody struct {
	QueryOption
	UserID int
}

type QueryProductUserResponseItem struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
