package themap

import (
	// "fmt"
	"testing"
)

func TestHook(t *testing.T) {

	tests := []struct {
		payload []byte
		want    *Notify
	}{
		{[]byte(`Amount=1&AuthCode=777777&CardNumber=411111xxxxxx1112&CardUId=1bM8q4ZaJsrW4RyHLp5MunINfzZ&ErrCode=&MerchantContract=SendtipsTestTerminal&MerchantOrderId=1bM8q9QvSVt54NtIhS9JFtxaH10&Notification=AddCard&OriginalOrderId=CardAdd4&RRN=123456789&State=Voided&Success=true`),
			&Notify{Type: "AddCard", Amount: 1, CardUID: "1bM8q4ZaJsrW4RyHLp5MunINfzZ", CardNumber: "411111xxxxxx1112"}},

		{[]byte(`Amount=20000&AuthCode=777777&CardNumber=411111xxxxxx1112&CardUId=&ErrCode=&MerchantContract=SendtipsTestTerminal&MerchantOrderId=1bM8SIqrl1t8breAOXC1lnykhA4&Notification=Block&OriginalOrderId=TipNo3&RRN=123456789&State=Charged&Success=true`),
			&Notify{Type: "Block", Amount: 20000, CardNumber: "411111xxxxxx1112"}},
	}

	for _, test := range tests {

		notify, err := NewNotify(test.payload)
		if err != nil {
			t.Error("Error occurred: ", err.Error())
		}

		if notify.Type != test.want.Type {
			t.Error("Incorrect type")
		}

		if notify.CardUID != test.want.CardUID {
			t.Error("CardUID empty")
		}

		if notify.Amount != test.want.Amount {
			t.Error("Amount is wrong")
		}

	}

}
