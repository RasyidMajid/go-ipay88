package go_ipay88

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type GatewayIPay88 struct {
	client Client
}

func NewGatewayIPay88(client Client) *GatewayIPay88 {
	return &GatewayIPay88{client: client}
}

func (g *GatewayIPay88) PaymentRequest(request *Request) (*Response, error) {

	signature := GetSignatureV2(GetSignatureV2Request{
		MerchantKey:  g.client.MerchantKey,
		MerchantCode: g.client.MerchantCode,
		RefNo:        request.RefNo,
		Amount:       request.Amount,
		Currency:     request.Currency,
	})
	request.Signature = signature

	var res *Response
	var url = g.client.Env.String() + "/ePayment/WebService/PaymentAPI/Checkout"
	err := g.client.DoProses(http.MethodPost, url, request, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (g *GatewayIPay88) PaymentStatus(request *PaymentStatusRequest) (*ResponseStatus, error) {
	Fa, err := strconv.ParseFloat(request.Amount, 64)
	if err != nil {
		return nil, err
	}
	var resp string
	VFloat := fmt.Sprintf("%.2f", Fa)
	var url = g.client.Env.String() + "/epayment/enquiry.asp?MerchantCode=" + g.client.MerchantCode + "&RefNo=" + request.RefNo + "&Amount=" + strings.ReplaceAll(VFloat, ".", ",")
	err = g.client.DoProses(http.MethodGet, url, nil, &resp)
	if err != nil {
		return nil, err
	}

	out := &ResponseStatus{Description: resp}
	return out, nil
}
