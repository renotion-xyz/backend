package moonpay

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"net/url"
)

func (mp *MoonPayClient) signUrlParams(params *url.Values) string {
	query := "?" + params.Encode()
	hmac := hmac.New(sha256.New, []byte(mp.secretKey))
	hmac.Write([]byte(query))
	signature := base64.StdEncoding.EncodeToString(hmac.Sum(nil))
	return signature
}

func (mp *MoonPayClient) SignURL(currency string, wallet string) string {
	params := url.Values{
		"apiKey": {mp.apiKey},
	}
	if currency != "" {
		params.Add("currencyCode", currency)
	} else if mp.defaultCurrency != "" {
		params.Add("currencyCode", mp.defaultCurrency)
	}
	if wallet != "" {
		params.Add("walletAddress", wallet)
	}
	signature := mp.signUrlParams(&params)
	url := fmt.Sprintf("%s?%s&%s", mp.baseUrl, params.Encode(), url.Values{"signature": {signature}}.Encode())
	return url
}
