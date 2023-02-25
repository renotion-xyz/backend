package moralis

const (
	MORALIS_BASE_URL = "https://deep-index.moralis.io/api/v2"
)

type MoralisClient struct {
	baseUrl string
	apiKey  string
	chain   Chain
}

func NewClient(apiKey string, chain Chain) *MoralisClient {
	return &MoralisClient{
		apiKey:  apiKey,
		baseUrl: MORALIS_BASE_URL,
		chain:   chain,
	}
}
