package themap

// Payout makes payout to card
// orderid original orderid used in Charge or Pay methods
// card parameter is a cardno or shadow card representation
// second paramenter is pan flag, if true we treat first prameter as actual card number
func (p *Payment) Payout(orderid, card string, ispan bool) error {

	var err error

	if ispan {
		p.Card = Card{PAN: card}
	} else {
		p.Card = Card{UID: card}
	}

	p.CustomParamsRDY = CParamsRDY{OriginalOrderID: orderid}

	err = proceedRequest("POST", "/Payout", p)
	if err != nil {
		return err
	}

	err = p.checkErrors()

	return err

}
