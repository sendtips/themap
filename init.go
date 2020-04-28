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

	if p.Reply.ErrCode != "" {
		err = fmt.Errorf("[THEMAP] %w: %s (%s)\n", ErrReplyWithError, p.Reply.ErrCode, p.Reply.ErrMessage)
	}

	return err

}
