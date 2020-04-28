package themap

import (
	"fmt"
)

// StoreCard adds card
func (p *Payment) StoreCard(card, cvv, holder string, month, year int) error {

	var err error

	p.Card = Card{PAN: card, Month: month, Year: year, CVV: cvv, Holder: holder}

	err = proceedRequest("POST", "/storeCard", p)
	if err != nil {
		return err
	}

	if p.Reply.ErrCode != "" {
		err = fmt.Errorf("[THEMAP] %w: %s\n", ErrReplyWithError, p.Reply.ErrCode)
	}

	return err

}

// DeleteCard removes card
func (p *Payment) DeleteCard(card string) error {

	var err error

	p.Card = Card{UID: card}

	err = proceedRequest("POST", "/removeCard", p)
	if err != nil {
		return err
	}

	if p.Reply.ErrCode != "" {
		err = fmt.Errorf("[THEMAP] %w: %s\n", ErrReplyWithError, p.Reply.ErrCode)
	}

	return err

}
