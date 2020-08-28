package themap

import (
	"context"
)

// NewPayout create Payout object
func NewPayout(key, merchid string) *Payment {
	return &Payment{Key: key, MerchantID: merchid}
}

// Payout makes payout to card.
// Orderid patameter holds orderid used in Charge or Pay methods.
// card parameter is a cardno or shadow card representation.
// ispan paramenter if true we treat card prameter as pan, if false - card holds a shadow pan
func (p *Payment) Payout(ctx context.Context, amount int, orderid, card string, ispan bool) error {

	var err error

	p.Amount = amount

	if ispan {
		p.Card = Card{PAN: card}
	} else {
		p.Card = Card{UID: card}
	}

	p.CustomParamsRDY = CParamsRDY{OriginalOrderID: orderid}

	err = proceedRequest(ctx, "POST", "/Payout", p)
	if err != nil {
		return err
	}

	err = p.checkErrors()

	return err

}
