package cloudflare

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
)

func (client *CloudflareClient) GetHostnameInfo(ctx context.Context, hostname string) (*cloudflare.CustomHostname, error) {
	items, _, err := client.api.CustomHostnames(ctx, client.zoneID, 0, cloudflare.CustomHostname{Hostname: hostname})
	if err != nil {
		return nil, err
	}
	if len(items) > 0 {
		item := items[0]
		return &item, nil
	}
	return nil, &NotFoundError{hostname}
}

func (client *CloudflareClient) RegisterHostname(ctx context.Context, hostname string) (*cloudflare.CustomHostname, error) {
	ch := cloudflare.CustomHostname{
		Hostname: hostname,
		SSL: &cloudflare.CustomHostnameSSL{
			Method: "txt",
			Type:   "dv",
			Settings: cloudflare.CustomHostnameSSLSettings{
				MinTLSVersion: "1.2",
				TLS13:         "on",
				HTTP2:         "on",
			},
		},
	}
	response, err := client.api.CreateCustomHostname(ctx, client.zoneID, ch)
	if err != nil {
		return nil, err
	}
	item := response.Result
	return &item, nil
}
