package go_ipay88

import (
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func GetSignature(source string) string {
	hash := sha1.New()
	hash.Write([]byte(source))
	signature := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return signature
}
func GetSignatureV2(request GetSignatureV2Request) string {
	//format Signature => || MerchantKey || MerchantCode || RefNo || Amount || Currency ||
	source := "||" + request.MerchantKey + "||" + request.MerchantCode + "||" + request.RefNo + "||" + request.Amount + "||" + request.Currency + "||"
	hash := sha256.New()
	hash.Write([]byte(source))
	res := fmt.Sprintf("%x", hash.Sum(nil))
	return string(res)
}
