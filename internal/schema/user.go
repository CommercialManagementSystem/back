package schema

type UpdateUserRequestBody struct {
	ID        int    ` json:"id"`
	UID       string `json:"uid"`
	Password  string `json:"password"`
	Authority int    `json:"authority"`

	Name  string `json:"Name"`
	Sex   int    `json:"Sex"`
	Phone string `json:"phone"`
	Email string `json:"email"`

	Role int `json:"role"`
}

type DeleteUserRequestBody struct {
	IDs []int `json:"ids"`
}

type CreateUserRequestBody struct {
	UID       string `json:"uid"`
	Password  string `json:"password"`
	Authority int    `json:"authority"`

	Name  string `json:"Name"`
	Sex   int    `json:"Sex"`
	Phone string `json:"phone"`
	Email string `json:"email"`

	Role int `json:"role"`
}

type QueryUserRequestBody struct {
	QueryOption
	Name  string `json:"name"`
	Phone string `json:"phone"`
}
