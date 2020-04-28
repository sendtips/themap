package themap

import (
	"fmt"
)

// Init obtain session token from TheMAP payment gateway
func (p *Payment) Init(amount int) error {
	var err error

	p.Amount = amount

	err = proceedRequest("POST", "/Init", p)
	if err != nil {
		return err
	}

	err = p.checkErrors()

	return err

}
