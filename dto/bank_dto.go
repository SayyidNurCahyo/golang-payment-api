package dto

type SaveBankRequest struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateBankRequest struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type GetBankResponse struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}
