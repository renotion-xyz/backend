package moralis

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func (mc *MoralisClient) newRequest(method string, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("x-api-key", mc.apiKey)

	return req, nil
}

func (mc *MoralisClient) GetNFTsByOwner(owner string, token string, chain Chain) (*NFTsByOwnerResponse, error) {
	query := url.Values{
		"chain":              {string(chain)},
		"token_addresses[0]": {token},
		"limit":              {"100"},
	}

	url := fmt.Sprintf("%s/%s/nft?%s", mc.baseUrl, owner, query.Encode())
	req, err := mc.newRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response NFTsByOwnerResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return nil, err
	}

	return &response, nil
}
