package schema

type QueryProductRequestBody struct {
	QueryOption
	Name string `form:"name"`
}

type UpdateProductRequestBody struct {
	BaseProductRequestBody
	ID string `json:"id"`
}

type BaseProductRequestBody struct {
	Name         string `json:"name"`
	ExpenseRatio string `json:"expense_ratio"`
	CompanyID    int    `json:"company_id"`
	Plan         string `json:"plan"`
	Rule         string `json:"rule"`
	Steps        string `json:"steps"`
}

type CreateProductRequestBody struct {
	BaseProductRequestBody
}

type DeleteProductRequestBody struct {
	IDs []int `json:"ids"`
}
