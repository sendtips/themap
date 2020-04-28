package themap

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestPayout(t *testing.T) {

	reply := `{
  "Status": "True",
  "OrderId": "TestOrder123",
  "Amount": 300,
  "ErrCode": "",
  "BankName": "TestBank"
}`

	//     reply_err := `{
	//     "Success": false,
	//     "ErrCode": "INVALID_AUTHENTICATION"
	// }`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api-stage.mapisacard.com/Payout",
		httpmock.NewStringResponder(200, reply))

	trans := New("123", "123")
	//trans.SetAuthUser("login", "123")
	trans.SetTerm("123")
	err := trans.Payout("123CARDSHADOW", false)

	if err != nil {
		t.Error("Error occurred", err.Error())
	}

	if trans.BankName != "TestBank" {
		t.Error("Wrong bankname BankName")
	}

}
