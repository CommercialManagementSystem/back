package schema

type DownloadProductAppendixRequestBody struct {
	ProductID int `json:"id"`
}

type DeleteProductAppendixRequestBody struct {
	ProductID []int `json:"ids"`
}

type AddProductAppendixRequestBody struct {
	Items []AddProductAppendixRequestItem `json:"items"`
}

type AddProductAppendixRequestItem struct {
	ProductID int    `json:"id"`
	Name      string `json:"name"`
	Url       string `json:"url"`
}
