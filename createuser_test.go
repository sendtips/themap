package themap

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestCreateUser(t *testing.T) {

	reply := `{
    "Success": true,
    "UserId": 11,
    "AlreadyCreated": false,
    "ErrCode": ""
}`

	reply_err := `{
        "Success": false,
        "UserId": 11,
        "AlreadyCreated": true,
        "Success": false,
        "ErrCode": "WRONG_PARAMS"
    }`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "https://api-stage.mapisacard.com/createUser",
		httpmock.NewStringResponder(200, reply))

	trans := New("123", "123")
	trans.SetAuthUser("login", "123")
	trans.SetTerm("123")
	err := trans.CreateUser("123.123.123.123", "18005000000", "username@example.com")

	if err != nil {
		t.Error("Error occurred", err.Error())
	}

	if trans.AlreadyCreated != false {
		t.Error("User alredy creaded flag is wrong")
	}

	httpmock.RegisterResponder("POST", "https://api-stage.mapisacard.com/createUser",
		httpmock.NewStringResponder(200, reply_err))

	trans_err := New("123", "123")
	trans_err.SetAuthUser("login", "123")
	trans_err.SetTerm("123")
	err2 := trans_err.CreateUser("123.123.123.123", "18005000000", "username@example.com")

	if trans_err.AlreadyCreated != true {
		t.Error("User alredy creaded flag is wrong")
	}

	if err2 == nil {
		t.Error("Error not returned")
	}

}
