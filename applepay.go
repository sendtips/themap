package themap

//  ApplePay method executes ApplePay payment
func (p *Payment) ApplePay(amount int, token string) error {
	var err error

	p.ApplePayToken = token
	p.Amount = amount

	err = proceedRequest("POST", "/Pay", p)
	if err != nil {
		return err
	}

	err = p.checkErrors()

	return err

}
