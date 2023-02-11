package moralis

import "time"

type Status string

const (
	SYNCING Status = "SYNCING"
	SYNCED  Status = "SYNCED"
)

type NFTResult struct {
	TokenAddress      string     `json:"token_address"`
	TokenID           string     `json:"token_id"`
	OwnerOf           string     `json:"owner_of"`
	BlockNumber       string     `json:"block_number"`
	BlockNumberMinted string     `json:"block_number_minted"`
	TokenHash         string     `json:"token_hash"`
	Amount            string     `json:"amount"`
	ContractType      string     `json:"contract_type"`
	Name              string     `json:"name"`
	Symbol            string     `json:"symbol"`
	TokenURI          *string    `json:"token_uri"`
	Metadata          *string    `json:"metadata"`
	LastTokenURISync  *time.Time `json:"last_token_uri_sync"`
	LastMetadataSync  *time.Time `json:"last_metadata_sync"`
	MinterAddress     string     `json:"minter_address"`
}

type NFTsByOwnerResponse struct {
	Result []NFTResult `json:"result"`
	Status Status      `json:"status"`
}
