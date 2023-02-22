package go_ipay88

type Request struct {
	APIVersion       string             `json:"ApiVersion"`
	MerchantCode     string             `json:"MerchantCode"`
	PaymentID        string             `json:"PaymentId"`
	Currency         string             `json:"Currency"`
	RefNo            string             `json:"RefNo"`
	Amount           string             `json:"Amount"`
	ProdDesc         string             `json:"ProdDesc"`
	RequestType      string             `json:"RequestType"`
	UserName         string             `json:"UserName"`
	UserEmail        string             `json:"UserEmail"`
	UserContact      string             `json:"UserContact"`
	Remark           string             `json:"Remark"`
	Lang             string             `json:"Lang"`
	ResponseURL      string             `json:"ResponseURL"`
	BackendURL       string             `json:"BackendURL"`
	Signature        string             `json:"Signature"`
	ItemTransactions []ItemTransactions `json:"ItemTransactions"`
	ShippingAddress  ShippingAddress    `json:"ShippingAddress"`
	BillingAddress   BillingAddress     `json:"BillingAddress"`
	Sellers          []Sellers          `json:"Sellers"`
	SettingField     []SettingField     `json:"SettingField"`
}
type ItemTransactions struct {
	ID         string `json:"Id"`
	Name       string `json:"Name"`
	Quantity   string `json:"Quantity"`
	Amount     string `json:"Amount"`
	ParentType string `json:"ParentType,omitempty"`
	ParentID   string `json:"ParentId,omitempty"`
}

type ShippingAddress struct {
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	Address     string `json:"Address"`
	City        string `json:"City"`
	State       string `json:"State"`
	PostalCode  string `json:"PostalCode"`
	Phone       string `json:"Phone"`
	CountryCode string `json:"CountryCode"`
}

type BillingAddress struct {
	FirstName   string `json:"FirstName"`
	LastName    string `json:"LastName"`
	Address     string `json:"Address"`
	City        string `json:"City"`
	State       string `json:"State"`
	PostalCode  string `json:"PostalCode"`
	Phone       string `json:"Phone"`
	CountryCode string `json:"CountryCode"`
}

type Sellers struct {
	ID             string  `json:"Id"`
	Name           string  `json:"Name"`
	LegalID        string  `json:"LegalId"`
	SellerIDNumber string  `json:"SellerIdNumber"`
	Email          string  `json:"Email"`
	URL            string  `json:"Url"`
	Address        Address `json:"address"`
}

type Address struct {
	FirstName   string `json:"FirstName" validate:"required"`
	LastName    string `json:"LastName"`
	Address     string `json:"Address"`
	City        string `json:"City"`
	State       string `json:"State"`
	PostalCode  string `json:"PostalCode"`
	Phone       string `json:"Phone"`
	CountryCode string `json:"CountryCode"`
}

type SettingField struct {
	Name  string `json:"Name"`
	Value string `json:"Value"`
}
type GetSignatureV2Request struct {
	MerchantKey  string
	MerchantCode string
	RefNo        string
	Amount       string
	Currency     string
}

type PaymentStatusRequest struct {
	RefNo  string `json:"refno"`
	Amount string `json:"amount"`
}

type NotificationRequest struct {
	MerchantCode       string `json:"MerchantCode"`
	PaymentId          string `json:"PaymentId"`
	RefNo              string `json:"RefNo"`
	Amount             string `json:"Amount"`
	Currency           string `json:"Currency"`
	Remark             string `json:"Remark"`
	TransId            string `json:"TransId"`
	AuthCode           string `json:"AuthCode"`
	TransactionStatus  string `json:"TransactionStatus"`
	ErrDesc            string `json:"ErrDesc"`
	Signature          string `json:"Signature"`
	IssuerBank         string `json:"IssuerBank"`
	PaymentDate        string `json:"PaymentDate"`
	Xfield1            string `json:"Xfield1"`
	DCCConversionRate  string `json:"DCCConversionRate"`
	OriginalAmount     string `json:"OriginalAmount"`
	OriginalCurrency   string `json:"OriginalCurrency"`
	SettlementAmount   string `json:"SettlementAmount"`
	SettlementCurrency string `json:"SettlementCurrency"`
	Binbank            string `json:"Binbank"`
}

type NotificationResponse struct {
	Code    string `json:"Code"`
	Message struct {
		English    string `json:"English"`
		Indonesian string `json:"Indonesian"`
	} `json:"Message"`
}
