package themap

import (
	"context"
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestNewPayout(t *testing.T) {
	trans := NewPayout("some", "123")
	dummy := &Payment{}

	if reflect.TypeOf(trans) != reflect.TypeOf(dummy) {
		t.Errorf("NewPayout() wrong return %T must be %T type", reflect.TypeOf(trans), reflect.TypeOf(dummy))
	}
}

func TestPayout(t *testing.T) {

	reply := `{
  "Status": "True",
  "OrderId": "TestOrder123",
  "Amount": 300,
  "ErrCode": "",
  "BankName": "TestBank"
}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", APILink+"/Payout",
		httpmock.NewStringResponder(200, reply))

	ctx := context.Background()

	// Shadow card number
	trans := NewPayout("123", "123")
	trans.SetTerm("123")
	err := trans.Payout(ctx, 300, "TipNo1", "123CARDSHADOW", false)

	if err != nil {
		t.Error("Error occurred", err.Error())
	}

	if trans.BankName != "TestBank" {
		t.Error("Wrong bankname BankName")
	}

	// PAN
	transpan := NewPayout("123", "123")
	transpan.SetTerm("123")
	err2 := transpan.Payout(ctx, 300, "TipNo2", "4111111111111112", true)

	if err2 != nil {
		t.Error("Error occurred", err2.Error())
	}

	if transpan.BankName != "TestBank" {
		t.Error("Wrong bankname BankName")
	}

}
