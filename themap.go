package themap

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
	Reply           Reply
}

type Credential struct {
	Login            string `json:"login"`
	Password         string `json:"password"`
	MerchantName     string `json:"merchant_name"`
	MerchantPassword string `json:"merchant_password"`
	TerminalPassword string `json:"terminal_password"`
}

type Reply struct {
	Success      bool   `json:"Success"`
	ReplyOrderID string `json:"OrderId"`
	ReplyAmount  int    `json:"Amount"`
	ErrCode      string `json:"ErrCode"`
	ReplyType    string `json:"Type"`
	SessionGUID  string `json:"SessionGUID"`
}

func New(key, merchid string) *Payment {
	return &Payment{Key: key, MerchantID: merchid, Type: "pay", Lifetime: 10000, PaymentType: "OneStep"}
}

func (p *Payment) SetAuthUser(login, passwd string) {
	p.Credential.Login = login
	p.Credential.Password = passwd
}

func (p *Payment) SetMerch(name, passwd string) {
	p.Credential.MerchantName = name
	p.Credential.MerchantPassword = passwd
}

func (p *Payment) SetTerm(passwd string) {
	p.Credential.TerminalPassword = passwd
}
