package themap

import (
	"fmt"
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

}

// The Init method obtain session token
// from TheMAP payment gateway
func ExampleInit() {
	p := New("SendtipsTestTerminal", "TestOrder123")
	p.SetTerm("123")

	err := p.Init(300) // Create session for 3.00RUB
	if err != nil {
		fmt.Printf("Error ocurred: %v", err)
	}

	fmt.Printf("%v", p.Reply.Success) // Will have a theMAP reply success flag
	// Output: true
}
