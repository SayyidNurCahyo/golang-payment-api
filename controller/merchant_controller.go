package controller

import (
	"fmt"
	// "merchant-payment-api/model"
	"merchant-payment-api/service"
)

type MerchantController struct {
	merchantService service.MerchantService
}

// func (m *MerchantController) insertMerchant(){
// 	err := m.merchantService.Create(model.Merchant{
// 		Name: "merchant1",
// 		PhoneNumber: "085708813281",
// 		Address: "indonesia",
// 	})
// 	if err!=nil{
// 		fmt.Println(err)
// 		return
// 	}
// }

func (m *MerchantController) getAllMerchant(){
	merchants, err := m.merchantService.FindAll()
	if err!= nil{
		fmt.Println(err)
		return
	}
	if len(merchants)==0{
		fmt.Println("data kosong")
		return
	}
	for _, merchant := range merchants{
		fmt.Printf("ID: %s, Name: %s, Phone Number: %s, Address: %s", merchant.Id, merchant.Name, merchant.PhoneNumber, merchant.Address)
	}
}

func (m *MerchantController) getMerchantById() {

}

func NewMerchantController(merchantService service.MerchantService) *MerchantController {
	return &MerchantController{merchantService: merchantService}
}