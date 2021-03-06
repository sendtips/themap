package themap

import (
	"context"
)

// Init obtain session token from TheMAP payment gateway
func (p *Payment) Init(ctx context.Context, amount int) error {
	var err error

	p.Amount = amount

	err = proceedRequest(ctx, "POST", "/Init", p)

	return p.checkErrors(err)

}
