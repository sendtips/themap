package themap

import (
	// "fmt"
	"testing"
)

// func TestHookPayment(t *testing.T) {
//
//     payload_charge := `Amount=20000&AuthCode=777777&CardNumber=411111xxxxxx1112&CardUId=&ErrCode=&MerchantContract=SendtipsTestTerminal&MerchantOrderId=1bM8SIqrl1t8breAOXC1lnykhA4&Notification=Block&OriginalOrderId=TipNo3&RRN=123456789&State=Charged&Success=true`
//     payload_block := `Amount=20000&AuthCode=777777&CardNumber=411111xxxxxx1112&CardUId=&ErrCode=&MerchantContract=SendtipsTestTerminal&MerchantOrderId=1bM8SIqrl1t8breAOXC1lnykhA4&Notification=Block&OriginalOrderId=TipNo3&RRN=123456789&State=Charged&Success=true`
//
// }

func TestHookAddCard(t *testing.T) {

	payload := []byte(`Amount=1&AuthCode=777777&CardNumber=411111xxxxxx1112&CardUId=1bM8q4ZaJsrW4RyHLp5MunINfzZ&ErrCode=&MerchantContract=SendtipsTestTerminal&MerchantOrderId=1bM8q9QvSVt54NtIhS9JFtxaH10&Notification=AddCard&OriginalOrderId=CardAdd4&RRN=123456789&State=Voided&Success=true`)

	notify, err := NewNotify(payload)
	if err != nil {
		t.Error("Error occurred: ", err.Error())
	}

	if notify.CardUID == "" {
		t.Error("CardUID empty")
	}
}
