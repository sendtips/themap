// Package themap implements client library to work with TheMAP payment gateway.
package themap

import (
	"fmt"
)

// Payment holds query to initialize payment session
type Payment struct {
	Key             string     `json:"key"`
	MerchantID      string     `json:"merchant_order_id"`
	Amount          int        `json:"amount"`
	AddCard         bool       `json:"add_card,omitempty"`
	Type            string     `json:"type,omitempty"`
	PaymentType     string     `json:"payment_type,omitempty"`
	Lifetime        int        `json:"lifetime,omitempty"`
	Credential      Credential `json:"credential,omitempty"`
	CustomParamsRow string     `json:"custom_params_raw,omitempty"`
	CardUID         string     `json:"card_uid,omitempty"`
	Action          string     `json:"action,omitempty"`
	ApplePayToken   string     `json:"apple_pay_token,omitempty"`
	GooglePayToken  string     `json:"google_pay_token,omitempty"`
	Recurrent       bool       `json:"recurrent,omitempty"`
	Card            Card       `json:"card,omitempty"`
	User

	Reply
}

// Credential holds auth credentials
type Credential struct {
	Login            string `json:"login"`
	Password         string `json:"password"`
	MerchantName     string `json:"merchant_name"`
	MerchantPassword string `json:"merchant_password"`
	TerminalPassword string `json:"terminal_password"`
}

// Reply carriers reply from gateway
type Reply struct {
	Success        bool   `json:"Success,omitempty"`
	ReplyOrderID   string `json:"OrderId,omitempty"`
	ReplyAmount    int    `json:"Amount,omitempty"`
	ErrMessage     string `json:"ErrMessage,omitempty"`
	ErrCode        string `json:"ErrCode,omitempty"`
	ReplyType      string `json:"Type,omitempty"`
	SessionGUID    string `json:"SessionGUID,omitempty"`
	BankName       string `json:"BankName,omitempty"`
	UserID         int    `json:"UserId,omitempty"`
	AlreadyCreated bool   `json:"AlreadyCreated,omitempty"`
	CardUID        string `json:"CardUId,omitempty"`
	PANMask        string `json:"PANMask,omitempty"`
	CardActive     bool   `json:"IsActive,omitempty"`
}

// Card represents card at TheMAP
type Card struct {
	// Card number
	PAN string `json:"pan,omitempty"`
	// Card identifier
	UID string `json:"uid,omitempty"`
	//  Month expirity of the card
	Month int `json:"emonth,omitempty"`
	// Year expirity of the card
	Year int `json:"eyear,omitempty"`
	// CVV card code
	CVV string `json:"cvv,omitempty"`
	// Card holder name
	Holder string `json:"holder,omitempty"`
}

// User represents user at TheMAP
type User struct {
	// Remote IP  address
	IP string `json:"ip"`
	// User phone number
	Phone string `json:"phone"`
	// User email address
	Email string `json:"email"`
}

// New constructs new query to initialize payment session
func New(key, merchid string) *Payment {
	return &Payment{Key: key, MerchantID: merchid, Type: "pay", Lifetime: 10000, PaymentType: "OneStep"}
}

// SetAuthUser sets user credentials
func (p *Payment) SetAuthUser(login, passwd string) {
	p.Credential.Login = login
	p.Credential.Password = passwd
}

// SetMerch sets Merchant credentials
func (p *Payment) SetMerch(name, passwd string) {
	p.Credential.MerchantName = name
	p.Credential.MerchantPassword = passwd
}

// SetTerm sets Terminal password
func (p *Payment) SetTerm(passwd string) {
	p.Credential.TerminalPassword = passwd
}

// checkErrors checks if errors is presented in reply
func (p *Payment) checkErrors() error {
	var err error
	if p.Reply.ErrCode != "" {
		err = fmt.Errorf("[THEMAP] %w: %s (%s)\n", ErrReplyWithError, p.Reply.ErrCode, p.Reply.ErrMessage)
	}

	return err
}
