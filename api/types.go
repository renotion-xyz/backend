package api

import (
	"github.com/renotion-xyz/backend/cloudflare"
	"github.com/renotion-xyz/backend/moralis"
)

type ListTokensResponse struct {
	Tokens []moralis.NFTResult `json:"tokens"`
}

type DomainMetadataResponse struct {
	TokenID  string  `json:"tokenID"`
	Hostname *string `json:"hostname"`
	Page     *string `json:"page"`
}

type TXTRecordDetails struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type DomainStatus struct {
	Page             string                           `json:"page"`
	TokenID          string                           `json:"tokenID"`
	Hostname         string                           `json:"hostname"`
	OwnershipStatus  cloudflare.OwnerhsipStatus       `json:"ownershipStatus"`
	SSLStatus        cloudflare.SSLVerificationStatus `json:"sslStatus"`
	TXTRecordDetails *TXTRecordDetails                `json:"txtRecordDetails"`
}

type DomainsListResponse struct {
	Domains []DomainStatus `json:"domains"`
}
