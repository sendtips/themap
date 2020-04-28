package themap

import (
	"fmt"
)

// Payout makes payout to card
// card parameter is a cardno or shadow card representation
// second paramenter is pan flag, if true we treat first prameter as actual card number
func (p *Payment) Payout(card string, pan bool) error {

	var err error

	if pan {
		p.Card = Card{PAN: card}
	} else {
		p.Card = Card{UID: card}
	}

	err = proceedRequest("POST", "/Payout", p)
	if err != nil {
		return err
	}

	if p.Reply.ErrCode != "" {
		err = fmt.Errorf("[THEMAP] %w: %s\n", ErrReplyWithError, p.Reply.ErrCode)
	}

	return err

}
