package themap

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestAddCard(t *testing.T) {

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
	reply_err := `{
    "Success": false,
    "ErrCode": "INVALID_AUTHENTICATION"
}`

	httpmock.RegisterResponder("POST", "https://api-stage.mapisacard.com/storeCard",
		httpmock.NewStringResponder(200, reply_err))

	trans_err := New("123", "123")
	trans_err.SetAuthUser("login", "")
	trans_err.SetTerm("123")
	err2 := trans.StoreCard("4300000000000777", "123", "Ivan Ivanov", 12, 21)

	if trans_err.Success != false {
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
	reply_err := `{
    "Success": false,
    "ErrCode": "INVALID_AUTHENTICATION"
}`

	httpmock.RegisterResponder("POST", "https://api-stage.mapisacard.com/removeCard",
		httpmock.NewStringResponder(200, reply_err))

	trans_err := New("123", "123")
	trans_err.SetAuthUser("login", "")
	trans_err.SetTerm("123")
	err2 := trans_err.DeleteCard("456ceFOFYXmjlZraP12nfP")

	if trans_err.Success != false {
		t.Error("Success delete wrong card")
	}

	if err2 == nil {
		t.Error("Error not returned")
	}

}
