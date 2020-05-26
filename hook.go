package themap

import (
	"net/url"
	"strconv"
)

// type NotifyType uint8
//
// const (
//     AddCard NotifyType = iota
//     Block
//     Charged
//     Void
// )

// Notify holds notification payload
type Notify struct {
	MerchantContract string
	OriginalOrderId  string
	MerchantOrderId  string
	Amount           int
	AuthCode         string
	RRN              string
	Success          bool
	CardNumber       string
	BankName         string
	ErrCode          string
	State            string
	Type             string
	CardUID          string
	CustomParams     string
	FeePercent       float32
	TerminalID       int
	Signature        string
}

// NewNotify returns Notify data from bytes
func NewNotify(s, signkey string) (*Notify, error) {

	var err error

	p, err := url.ParseQuery(s)
	if err != nil {
		return nil, err
	}

	notify := &Notify{MerchantContract: p.Get("MerchantContract"), OriginalOrderId: p.Get("OriginalOrderId"), MerchantOrderId: p.Get("MerchantOrderId"),
		AuthCode: p.Get("AuthCode"), RRN: p.Get("RRN"), CardNumber: p.Get("CardNumber"), BankName: p.Get("BankName"), ErrCode: p.Get("ErrCode"),
		State: p.Get("State"), Type: p.Get("Notification"), CardUID: p.Get("CardUId"), CustomParams: p.Get("CustomParams"), Signature: p.Get("Signature")}

	if p.Get("Amount") != "" {
		notify.Amount, err = strconv.Atoi(p.Get("Amount"))
		if err != nil {
			notify.Amount = 0
		}
	}

	if p.Get("Success") != "" {
		notify.Success, err = strconv.ParseBool(p.Get("Success"))
		if err != nil {
			notify.Success = false
		}
	}

	if p.Get("FeePercent") != "" {
		fee, err := strconv.ParseFloat(p.Get("FeePercent"), 32)
		notify.FeePercent = float32(fee)
		if err != nil {
			notify.FeePercent = 0
		}
	}

	if p.Get("TerminalID") != "" {
		notify.TerminalID, err = strconv.Atoi(p.Get("TerminalID"))
		if err != nil {
			notify.TerminalID = 0
		}
	}

	// Check signature
	if notify.Signature != "" {
		sig := NewSignature(signkey, notify.Signature)
		if !sig.Verify(s) {
			err = ErrBadSignature
		}
	}

	return notify, err
}
