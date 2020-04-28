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
	Success        bool   `json:"Success"`
	ReplyOrderID   string `json:"OrderId"`
	ReplyAmount    int    `json:"Amount"`
	ErrCode        string `json:"ErrCode"`
	ReplyType      string `json:"Type"`
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

// Card represents card at TheMAP
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
