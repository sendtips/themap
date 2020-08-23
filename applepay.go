package themap

import (
	"encoding/base64"
)

// ApplePay method executes ApplePay payment
func (p *Payment) ApplePay(amount int, token []byte) error {
	var err error

	p.ApplePayToken = base64.StdEncoding.EncodeToString(token)
	p.Amount = amount

	err = proceedRequest("POST", "/Pay", p)
	if err != nil {
		return err
	}

	err = p.checkErrors()

	return err

}
