package themap

import (
	"context"
)

// AddCardSession makes session request to store user card
func (p *Payment) AddCardSession(ctx context.Context) error {

	var err error

	// skip validation at bank
	p.Amount = 1
	p.Type = "add"
	p.AddCard = true

	err = proceedRequest(ctx, "POST", "/Init", p)

	return p.checkErrors(err)

}

// StoreCard adds card
func (p *Payment) StoreCard(ctx context.Context, card, cvv, holder string, month, year int) error {

	var err error

	p.Card = Card{PAN: card, Month: month, Year: year, CVV: cvv, Holder: holder}

	err = proceedRequest(ctx, "POST", "/storeCard", p)

	return p.checkErrors(err)

}

// DeleteCard removes card
func (p *Payment) DeleteCard(ctx context.Context, card string) error {

	var err error

	p.Card = Card{UID: card}

	err = proceedRequest(ctx, "POST", "/removeCard", p)

	return p.checkErrors(err)

}
