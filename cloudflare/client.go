package cloudflare

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
)

type CloudflareClient struct {
	api    *cloudflare.API
	zoneID string
}

func NewClient(ctx context.Context, apiToken string, zoneName string) (*CloudflareClient, error) {
	api, err := cloudflare.NewWithAPIToken(apiToken)
	if err != nil {
		return nil, err
	}
	id, err := api.ZoneIDByName(zoneName)
	if err != nil {
		return nil, err
	}
	return &CloudflareClient{api, id}, nil
}
