package moonpay

import (
	"net/url"
	"testing"
)

func TestSignUrl(t *testing.T) {
	tests := []struct {
		secretKey string
		params    *url.Values
		expected  string
	}{
		{
			secretKey: "sk_test_key",
			params: &url.Values{
				"apiKey":        {"pk_test_key"},
				"currencyCode":  {"eth"},
				"walletAddress": {"0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae"},
			},
			expected: "o1eO2arzL16TstEolGakHHFa33Mb61RCTfqToRWg7PA=",
		},
		{
			secretKey: "feeb",
			params: &url.Values{
				"apiKey": {"beef"},
			},
			expected: "MJt8SjhZqhkoPvcEGelchtt76AT78ZWKXeXlbNewxuU=",
		},
		{
			secretKey: "feeb",
			params: &url.Values{
				"apiKey":       {"beef"},
				"currencyCode": {"btc"},
			},
			expected: "v50paGXnMPaywL2j5eai/jffU2CwOf/T0aM3+9Ktqrg=",
		},
		{
			secretKey: "feeb",
			params: &url.Values{
				"apiKey":        {"beef"},
				"currencyCode":  {"btc"},
				"walletAddress": {"0x123"},
			},
			expected: "UoZHrIlDuudOMPlOE+UH1viBLGZSPjJr/uCShljgJ0o=",
		},
		{
			secretKey: "beef",
			params: &url.Values{
				"apiKey":        {"feeb"},
				"currencyCode":  {"eth"},
				"walletAddress": {"0x456"},
			},
			expected: "w2ndGNCr9iz7OZ2tEdoYT2cbtBsCOkQHObJbxZIF0EA=",
		},
	}
	for _, test := range tests {
		mp := &MoonPayClient{
			secretKey: test.secretKey,
		}
		actual := mp.signUrlParams(test.params)
		if actual != test.expected {
			t.Fatalf(
				"signUrlParams(%s, %s), expected %s, got %s",
				test.secretKey,
				*test.params,
				test.expected,
				actual,
			)
		}
	}
}

func TestSignMoonPayURL(t *testing.T) {
	tests := []struct {
		apiKey    string
		secretKey string
		currency  string
		wallet    string
		expected  string
	}{
		{
			apiKey:    "pk_test_key",
			secretKey: "sk_test_key",
			currency:  "eth",
			wallet:    "0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae",
			expected:  "https://buy-sandbox.moonpay.com?apiKey=pk_test_key&currencyCode=eth&walletAddress=0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae&signature=o1eO2arzL16TstEolGakHHFa33Mb61RCTfqToRWg7PA%3D",
		},
		{
			apiKey:    "beef",
			secretKey: "feeb",
			expected:  "https://buy-sandbox.moonpay.com?apiKey=beef&signature=MJt8SjhZqhkoPvcEGelchtt76AT78ZWKXeXlbNewxuU%3D",
		},
		{
			apiKey:    "beef",
			secretKey: "feeb",
			currency:  "btc",
			expected:  "https://buy-sandbox.moonpay.com?apiKey=beef&currencyCode=btc&signature=v50paGXnMPaywL2j5eai%2FjffU2CwOf%2FT0aM3%2B9Ktqrg%3D",
		},
		{
			apiKey:    "beef",
			secretKey: "feeb",
			currency:  "btc",
			wallet:    "0x123",
			expected:  "https://buy-sandbox.moonpay.com?apiKey=beef&currencyCode=btc&walletAddress=0x123&signature=UoZHrIlDuudOMPlOE%2BUH1viBLGZSPjJr%2FuCShljgJ0o%3D",
		},
		{
			apiKey:    "feeb",
			secretKey: "beef",
			currency:  "eth",
			wallet:    "0x456",
			expected:  "https://buy-sandbox.moonpay.com?apiKey=feeb&currencyCode=eth&walletAddress=0x456&signature=w2ndGNCr9iz7OZ2tEdoYT2cbtBsCOkQHObJbxZIF0EA%3D",
		},
	}
	for _, test := range tests {
		mp := &MoonPayClient{
			apiKey:    test.apiKey,
			secretKey: test.secretKey,
			baseUrl:   "https://buy-sandbox.moonpay.com",
		}
		actual := mp.SignURL(test.currency, test.wallet)
		if actual != test.expected {
			t.Fatalf(
				"SignMoonPayURL(%s, %s, %s), expected %s, got %s",
				test.apiKey,
				test.currency,
				test.wallet,
				test.expected,
				actual,
			)
		}
	}
}
