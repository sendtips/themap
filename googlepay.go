package themap

import (
	"encoding/base64"
)

//  GooglePay method executes Google Pay payment
func (p *Payment) GooglePay(amount int, token []byte) error {
	var err error

	p.GooglePayToken = base64.StdEncoding.EncodeToString(token)
	p.Amount = amount

	err = proceedRequest("POST", "/Pay", p)
	if err != nil {
		return err
	}

	err = p.checkErrors()

	return err

}
