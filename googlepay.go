package themap

import (
	"context"
	"encoding/base64"
)

// GooglePay method executes Google Pay payment
func (p *Payment) GooglePay(ctx context.Context, amount int, token []byte) error {
	var err error

	p.GooglePayToken = base64.StdEncoding.EncodeToString(token)
	p.Amount = amount

	err = proceedRequest(ctx, "POST", "/Pay", p)

	err = p.checkErrors(err)

	return err

}
