package dto

type SaveProductRequest struct {
	MerchantId string `json:"merchantId"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
}

type UpdateProductRequest struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type GetProductResponse struct {
	Id           string `json:"id"`
	MerchantId   string `json:"merchantId"`
	MerchantName string `json:"merchantName"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
}
