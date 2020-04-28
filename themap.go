package themap

// Payment holds query to initialize payment session
type Payment struct {
	Key             string     `json:"key"`
	MerchantID      string     `json:"merchant_order_id"`
	Amount          int        `json:"amount"`
	AddCard         bool       `json:"add_card"`
	Type            string     `json:"type"`
	PaymentType     string     `json:"payment_type"`
	Lifetime        int        `json:"lifetime"`
	Credential      Credential `json:"credential"`
	CustomParamsRow string     `json:"custom_params_raw"`
	CardUID         string     `json:"card_uid"`
	Action          string     `json:"action"`
	Recurrent       bool       `json:"recurrent"`
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
	Success      bool   `json:"Success,omitempty"`
	ReplyOrderID string `json:"OrderId,omitempty"`
	ReplyAmount  int    `json:"Amount,omitempty"`
	ErrCode      string `json:"ErrCode,omitempty"`
	ReplyType    string `json:"Type,omitempty"`
	SessionGUID  string `json:"SessionGUID,omitempty"`
}

// Card represents card at TheMAP
type Card struct {
	// Card number
	PAN string `json:"pan"`
	// Card identifier
	UID string `json:"uid"`
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
