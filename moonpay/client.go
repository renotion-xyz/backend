package moonpay

const (
	MOONPAY_BASE_URL = "https://buy-sandbox.moonpay.com"
)

type MoonPayClient struct {
	baseUrl         string
	apiKey          string
	secretKey       string
	defaultCurrency string
}

func NewClient(apiKey string, secretKey string, defaultCurrency string) *MoonPayClient {
	return &MoonPayClient{
		apiKey:          apiKey,
		secretKey:       secretKey,
		baseUrl:         MOONPAY_BASE_URL,
		defaultCurrency: defaultCurrency,
	}
}
