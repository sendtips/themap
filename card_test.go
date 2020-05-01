package themap

import (
	"fmt"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestAddCardSession(t *testing.T) {

	reply := `{
    "Success": true,
    "OrderId": "TestOrder123",
    "Amount": 300,
    "ErrCode": "",
    "Type": "add",
    "SessionGUID": "1ILZMU42Zs8YivEsYXOA67ijRYs"
}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api-stage.mapisacard.com/Init",
		httpmock.NewStringResponder(200, reply))

	trans := New("123", "123")
	trans.SetTerm("123")
	err := trans.AddCardSession()

	if err != nil {
		t.Error("Error occurred", err.Error())
	}

	if trans.Type != "add" {
		t.Error("Wrong session type:", trans.Type)
	}

	if trans.AddCard == false {
		t.Error("AddCard flag is false")
	}

	// if trans.CardUID == "" {
	//     t.Error("Empty card UID")
	// }
}

func TestStoreCard(t *testing.T) {

	reply := `{
    "Success": true,
    "CardUId": "7sTwecksRSs1fIpUQw8su",
    "PANMask": "411111xxxxxx1111",
    "IsActive": true,
    "ErrCode": ""
}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api-stage.mapisacard.com/storeCard",
		httpmock.NewStringResponder(200, reply))

	trans := New("123", "123")
	trans.SetAuthUser("login", "123")
	trans.SetTerm("123")
	err := trans.StoreCard("4300000000000777", "123", "Ivan Ivanov", 12, 21)

	if err != nil {
		t.Error("Error occurred", err.Error())
	}

	if trans.Reply.CardUID != "7sTwecksRSs1fIpUQw8su" {
		t.Errorf("Card UID mismatch want %s returned %s", "7sTwecksRSs1fIpUQw8su", trans.Reply.CardUID)
	}

	if trans.Reply.PANMask != "411111xxxxxx1111" {
		t.Errorf("Card PANMask mismatch want %s returned %s", "411111xxxxxx1111", trans.Reply.PANMask)
	}

	// Error reply
	replyErr := `{
    "Success": false,
    "ErrCode": "INVALID_AUTHENTICATION"
}`

	httpmock.RegisterResponder("POST", "https://api-stage.mapisacard.com/storeCard",
		httpmock.NewStringResponder(200, replyErr))

	transErr := New("123", "123")
	transErr.SetAuthUser("login", "")
	transErr.SetTerm("123")
	err2 := trans.StoreCard("4300000000000777", "123", "Ivan Ivanov", 12, 21)

	if transErr.Success != false {
		t.Error("Successfully added bad card")
	}

	if err2 == nil {
		t.Error("Error not returned")
	}

}

func TestDeleteCard(t *testing.T) {

	reply := `{
    "Success": true,
    "CardUId": "7sTwecksRSs1fIpUQw8su"
}`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api-stage.mapisacard.com/removeCard",
		httpmock.NewStringResponder(200, reply))

	trans := New("123", "123")
	trans.SetAuthUser("login", "123")
	trans.SetTerm("123")
	err := trans.DeleteCard("7sTwecksRSs1fIpUQw8su")

	if err != nil {
		t.Error("Error occurred", err.Error())
	}

	if trans.Reply.CardUID != "7sTwecksRSs1fIpUQw8su" {
		t.Errorf("Card UID mismatch want %s returned %s", "7sTwecksRSs1fIpUQw8su", trans.Reply.CardUID)
	}

	// Error reply
	replyErr := `{
    "Success": false,
    "ErrCode": "INVALID_AUTHENTICATION"
}`

	httpmock.RegisterResponder("POST", "https://api-stage.mapisacard.com/removeCard",
		httpmock.NewStringResponder(200, replyErr))

	transErr := New("123", "123")
	transErr.SetAuthUser("login", "")
	transErr.SetTerm("123")
	err2 := transErr.DeleteCard("456ceFOFYXmjlZraP12nfP")

	if transErr.Success != false {
		t.Error("Success delete wrong card")
	}

	if err2 == nil {
		t.Error("Error not returned")
	}

}

// The Init method obtain session token
// from TheMAP payment gateway for card manage
func ExampleAddCardSession() {
	pay := New("SendtipsTestTerminal", "TestOrder123")
	pay.SetAuthUser("login", "password123")
	pay.SetTerm("123")

	err := pay.AddCardSession() // Create add_card session
	if err != nil {
		fmt.Printf("Error occurred: %v", err)
	}

	fmt.Printf("%v", pay.Success) // Will have a theMAP reply success flag
	// Output: true
}
