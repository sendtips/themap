package themap

import (
	"testing"
)

func TestSignature(t *testing.T) {

	key := "haha"
	payload := `Amount=1&AuthCode=777777&CardNumber=411111xxxxxx0031&CardUId=&ErrCode=&MerchantContract=TestTerminal&MerchantOrderId=1bj7rhNW2spNCQ289Rh0wFsfdm8&Notification=Unblock&OriginalOrderId=CardAdd4&RRN=123456789&Signature=8d02e2ef0a1565552da96bface33425fdc7158b1e682ebca9ad137e35f902bd1&State=Voided&Success=true`

	sig := NewSignature(key, "8d02e2ef0a1565552da96bface33425fdc7158b1e682ebca9ad137e35f902bd1")
	isValid := sig.Verify(payload)

	if !isValid {
		t.Error("Wrong signature")
	}

}
