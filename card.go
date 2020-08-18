package themap

// AddCardSession makes session request to store user card
func (p *Payment) AddCardSession() error {

	var err error

	// skip validation at bank
	p.Amount = 1
	p.Type = "add"
	p.AddCard = true

	err = proceedRequest("POST", "/Init", p)
	if err != nil {
		return err
	}

	err = p.checkErrors()

	return err

}

// StoreCard adds card
func (p *Payment) StoreCard(card, cvv, holder string, month, year int) error {

	var err error

	p.Card = Card{PAN: card, Month: month, Year: year, CVV: cvv, Holder: holder}

	err = proceedRequest("POST", "/storeCard", p)
	if err != nil {
		return err
	}

	err = p.checkErrors()

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

	err = p.checkErrors()

	return err

}
