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
	p := New("SendtipsTestTerminal", "OrderID1")
	// p.SetAuthUser("", "")
	// p.SetMerch("", "")
	p.SetTerm("123")

	err := p.Init(300) // Create session for 3.00RUB
	if err != nil {
		fmt.Printf("Error ocurred: %v", err)
	}

	fmt.Printf("%v", p.Reply.Success) // Will have a theMAP session identifier
	// Output: true
}
