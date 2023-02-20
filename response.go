package go_ipay88

type Response struct {
	RefNo      string `json:"RefNo"`
	Signature  string `json:"Signature"`
	CheckoutID string `json:"CheckoutID"`
	Code       string `json:"Code"`
	Message    string `json:"Message"`
}
type ResponseStatus struct {
	Description string `json:"Description"`
}
