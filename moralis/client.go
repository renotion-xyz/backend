package moralis

const (
	MORALIS_BASE_URL = "https://deep-index.moralis.io/api/v2"
)

type MoralisClient struct {
	baseUrl string
	apiKey  string
}

func NewClient(apiKey string) *MoralisClient {
	return &MoralisClient{
		apiKey:  apiKey,
		baseUrl: MORALIS_BASE_URL,
	}
}
