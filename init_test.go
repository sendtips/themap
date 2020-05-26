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

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", APILink+"/Init",
		httpmock.NewStringResponder(200, reply))

	trans := New("123", "123")
	trans.SetTerm("123")
	err := trans.Init(200)

	if err != nil {
		t.Error("Error occurred", err.Error())
	}

}

// The Init method obtain session token
// from TheMAP payment gateway
func ExampleInit() {
	pay := New("SendtipsTestTerminal", "TestOrder123")
	pay.SetTerm("123")

	err := pay.Init(300) // Create session for 3.00RUB
	if err != nil {
		fmt.Printf("Error occurred: %v", err)
	}

	fmt.Printf("%v", pay.Success)
	// Output: true
}
