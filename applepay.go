package themap

import (
	"context"
	"encoding/base64"
)

// ApplePay method executes ApplePay payment
func (p *Payment) ApplePay(ctx context.Context, amount int, token []byte) error {
	var err error

	p.ApplePayToken = base64.StdEncoding.EncodeToString(token)
	p.Amount = amount

	err = proceedRequest(ctx, "POST", "/Pay", p)
	if err != nil {
		return err
	}

	err = p.checkErrors()

	return err

}
