package themap

import (
	"context"
)

// CreateUser registers a user at theMAP
func (p *Payment) CreateUser(ctx context.Context, ip, phone, email string) error {

	var err error

	p.User = User{IP: ip, Phone: phone, Email: email}

	err = proceedRequest(ctx, "POST", "/createUser", p)

	err = p.checkErrors(err)

	return err

}
