package themap

import (
	"context"
	"fmt"
	"os"
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

	ctx := context.Background()

	trans := New("123", "123")
	trans.SetTerm("123")
	err := trans.Init(ctx, 200)

	if err != nil {
		t.Error("Error occurred", err.Error())
	}

}

// The Init method obtain session token
// from TheMAP payment gateway
func ExampleInit() {
	// check themap hostname env is set, otherwise use default host
	apihost, ok := os.LookupEnv("THEMAPAPIHOST")
	if ok {
		APILink = apihost
	}

	pay := New(os.Getenv("THEMAPTERMID"), "TestOrder123")
	pay.SetTerm(os.Getenv("THEMAPTERMPW"))

	err := pay.Init(context.TODO(), 300) // Create session for 3.00RUB
	if err != nil {
		fmt.Printf("Error occurred: %v", err)
	}

	fmt.Printf("%v", pay.Success)
	// Output: true
}
