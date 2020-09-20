package themap

import (
	"errors"
	// "strconv"
	"testing"
)

func TestHook(t *testing.T) {

	tests := []struct {
		payload string
		want    *Notify
		err     error
		sig     string
	}{
		{`Amount=1&AuthCode=777777&CardNumber=411111xxxxxx1112&CardUId=1bM8q4ZaJsrW4RyHLp5MunINfzZ&ErrCode=&MerchantContract=SendtipsTestTerminal&MerchantOrderId=1bM8q9QvSVt54NtIhS9JFtxaH10&Notification=AddCard&OriginalOrderId=CardAdd4&RRN=123456789&State=Voided&Success=true`,
			&Notify{Type: "AddCard", Amount: 1, CardUID: "1bM8q4ZaJsrW4RyHLp5MunINfzZ", CardNumber: "411111xxxxxx1112"}, nil, ""},

		{`Amount=20000&AuthCode=777777&CardNumber=411111xxxxxx1112&CardUId=&ErrCode=&MerchantContract=SendtipsTestTerminal&MerchantOrderId=1bM8SIqrl1t8breAOXC1lnykhA4&Notification=Block&OriginalOrderId=TipNo3&RRN=123456789&State=Charged&Success=true`,
			&Notify{Type: "Block", Amount: 20000, CardNumber: "411111xxxxxx1112"}, nil, ""},

		// Bad amount
		{`Amount=z&AuthCode=777777&CardNumber=411111xxxxxx1112&CardUId=&ErrCode=&MerchantContract=SendtipsTestTerminal&MerchantOrderId=1bM8SIqrl1t8breAOXC1lnykhA4&Notification=Block&OriginalOrderId=TipNo3&RRN=123456789&State=Charged&Success=true`,
			&Notify{Type: "Block", Amount: 0, CardNumber: "411111xxxxxx1112"}, nil, ""},

		// Bad success
		// {`Amount=2000&AuthCode=777777&CardNumber=411111xxxxxx1112&CardUId=&ErrCode=&MerchantContract=SendtipsTestTerminal&MerchantOrderId=1bM8SIqrl1t8breAOXC1lnykhA4&Notification=Block&OriginalOrderId=TipNo3&RRN=123456789&State=Charged&Success=-1`,
		//   &Notify{Type: "Block", Amount: 2000, CardNumber: "411111xxxxxx1112"}, strconv.ErrSyntax, ""},

		// Bad fee
		{`FeePercent=z&Amount=2000&AuthCode=777777&CardNumber=411111xxxxxx1112&CardUId=&ErrCode=&MerchantContract=SendtipsTestTerminal&MerchantOrderId=1bM8SIqrl1t8breAOXC1lnykhA4&Notification=Block&OriginalOrderId=TipNo3&RRN=123456789&State=Charged&Success=true`,
			&Notify{Type: "Block", Amount: 2000, FeePercent: 1, CardNumber: "411111xxxxxx1112"}, nil, ""},

		// Bad terminal is
		// {`TerminalID=a&Amount=3000&AuthCode=777777&CardNumber=411111xxxxxx1112&CardUId=&ErrCode=&MerchantContract=SendtipsTestTerminal&MerchantOrderId=1bM8SIqrl1t8breAOXC1lnykhA4&Notification=Block&OriginalOrderId=TipNo3&RRN=123456789&State=Charged&Success=true`,
		//   &Notify{Type: "Block", Amount: 3000, CardNumber: "411111xxxxxx1112"}, strconv.ErrSyntax, "ssx"},

		// Bad signature
		{`Amount=2000&AuthCode=777777&CardNumber=411111xxxxxx1112&CardUId=&ErrCode=&MerchantContract=SendtipsTestTerminal&MerchantOrderId=1bM8SIqrl1t8breAOXC1lnykhA4&Signature=BadSignature&Notification=Block&OriginalOrderId=TipNo3&RRN=123456789&State=Charged&Success=true`,
			&Notify{Type: "Block", Amount: 2000, CardNumber: "411111xxxxxx1112"}, ErrBadSignature, "dd"},
	}

	for _, test := range tests {

		notify, err := NewNotify(test.payload, test.sig)

		if !errors.Is(err, test.err) {
			t.Error("Error occurred: ", err, test.err)
		}

		// if test.err == strconv.ErrSyntax {
		//   if !errors.Is(err, test.err) {
		//     t.Error("Error occurred: ", err, test.err)
		//   }
		// } else if err != test.err {
		//   t.Error("Error occurred: ", err, test.err)
		// }

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
