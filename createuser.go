package themap

// CreateUser registers a user at theMAP
func (p *Payment) CreateUser(ip, phone, email string) error {

	var err error

	p.User = User{IP: ip, Phone: phone, Email: email}

	err = proceedRequest("POST", "/createUser", p)
	if err != nil {
		return err
	}

	err = p.checkErrors()

	return err

}
