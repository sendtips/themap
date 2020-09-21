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

	return p.checkErrors(err)

}

// ApplePayBlock blocks
// Only blocks amount, without charge
// This method currentry uset for some internal tests
// To actually charge use ApplePay() method insread
func (p *Payment) ApplePayBlock(ctx context.Context, amount int, token []byte) error {
	var err error

	p.ApplePayToken = base64.StdEncoding.EncodeToString(token)
	p.Amount = amount

	err = proceedRequest(ctx, "POST", "/Block", p)

	return p.checkErrors(err)

}
