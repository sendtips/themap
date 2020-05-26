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

	replyErr := `{
        "Success": false,
        "UserId": 11,
        "AlreadyCreated": true,
        "Success": false,
        "ErrCode": "WRONG_PARAMS"
    }`

	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", APILink+"/createUser",
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

	if trans.UserID != 11 {
		t.Error("Empty MAP user id")
	}

	httpmock.RegisterResponder("POST", APILink+"/createUser",
		httpmock.NewStringResponder(200, replyErr))

	transErr := New("123", "123")
	transErr.SetAuthUser("login", "123")
	transErr.SetTerm("123")
	err2 := transErr.CreateUser("123.123.123.123", "18005000000", "username@example.com")

	if transErr.AlreadyCreated != true {
		t.Error("User alredy creaded flag is wrong")
	}

	if err2 == nil {
		t.Error("Error not returned")
	}

}
