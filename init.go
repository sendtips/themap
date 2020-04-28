package themap

import (
	"fmt"
)

// Init obtain session token from TheMAP payment gateway
func (p *Payment) Init(amount int) error {
	p.Amount = amount
	return p.Send("/Init")
}

func (p *Payment) Send(link string) error {
	var err error

	err = proceedRequest("POST", link, p)
	if err != nil {
		return err
	}

	if p.Reply.ErrCode != "" {
		err = fmt.Errorf("[THEMAP] %w: %s\n", ErrReplyWithError, p.Reply.ErrCode)
	}

	return err
}
