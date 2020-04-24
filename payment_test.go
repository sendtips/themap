package themap

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	trans := New("some", "123")
	dummy := &Payment{}

	if reflect.TypeOf(trans) != reflect.TypeOf(dummy) {
		t.Errorf("New() wrong return %T must be %T type", reflect.TypeOf(trans), reflect.TypeOf(dummy))
	}
}

func ExampleNew() {
	trans := New("some", "123")
	trans.SetAuthUser("user_login", "pass123")
	fmt.Printf("%+v", trans.Credential.Login)
	// Output: user_login
}
