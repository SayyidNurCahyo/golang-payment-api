package dto

type SaveCustomerRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

type UpdateCustomerRequest struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}

type GetCustomerResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
	Username    string `json:"username"`
}
