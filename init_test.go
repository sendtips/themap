package themap

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestInit(t *testing.T) {

	reply := `{
    "Success": true,
    "OrderId": "TestOrder123",
    "Amount": 300,
    "ErrCode": "",
    "Type": "pay",
    "SessionGUID": "1ILZMU42Zs8YivEsYXOA67ijRYs"
}`

	// reply_err := `{
	//         "Success": false,
	//         "ErrCode": "WRONG_PARAMS"
	//     }`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api-stage.mapisacard.com/Init",
		httpmock.NewStringResponder(200, reply))

	trans := New("123", "123")
	trans.SetAuthUser("login", "123")
	err := trans.Init(200)

	if err != nil {
		t.Error("Error occurred", err.Error())
	}

	// if ident.User.PayerID != "WDJJHEBZ4X2LY" {
	//     t.Error("Wrong payer id", ident.User.PayerID)
	// }

}
