package themap

import (
	"context"
	"fmt"
	// "net"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNewRequest(t *testing.T) {
	req, _ := newRequest(context.Background(), "POST", "/Init", nil, nil)
	dummy := &http.Request{}

	if reflect.TypeOf(req) != reflect.TypeOf(dummy) {
		t.Errorf("newRequest() wrong return %T must be %T type", reflect.TypeOf(req), reflect.TypeOf(dummy))
	}
}

func TestProceedRequest(t *testing.T) {
	var err error
	payload := New("TestTerminal", "TestOrder123")
	//listner, _ := net.Listen("tcp", APILink+":8060")

	// ErrBadJSON
	serv_badjson := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{{{{bad json")
	}))
	//serv_badjson.Listener = listner
	serv_badjson.Start()
	defer serv_badjson.Close()

	// Route request to mocked http server
	APILink = serv_badjson.URL

	err = proceedRequest(context.Background(), "POST", "/Init", payload)

	if err != ErrBadJSON {
		t.Errorf("Wrong error for bad JSON return")
	}

	// ErrBadStatusReply
	serv_errcode := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Something wrong", http.StatusInternalServerError)
	}))
	serv_errcode.Start()
	defer serv_errcode.Close()

	// Route request to mocked http server
	APILink = serv_errcode.URL

	err = proceedRequest(context.Background(), "POST", "/Init", payload)

	if err != ErrBadStatusReply {
		t.Errorf("Wrong error for error HTTP error code response")
	}

	// GoodRequest
	serv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		reply := `{
      "Success": true,
      "OrderId": "TestOrder123",
      "Amount": 300,
      "ErrCode": "",
      "Type": "pay",
      "SessionGUID": "1ILZMU42Zs8YivEsYXOA67ijRYs"
  }`
		fmt.Fprintln(w, reply)
	}))
	defer serv.Close()

	// Route request to mocked http server
	APILink = serv.URL

	err = proceedRequest(context.Background(), "POST", "/Init", payload)

	if payload.SessionGUID != "1ILZMU42Zs8YivEsYXOA67ijRYs" {
		t.Errorf("Wrong return in HTTP response")
	}

	if err != nil {
		t.Errorf("Error shoud be empty")
	}
}
