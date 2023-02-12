package api

import (
	"context"

	"github.com/renotion-xyz/backend/cloudflare"
)

func getDomainStatus(cf *cloudflare.CloudflareClient, tokenID string, hostname string, page string) (*DomainStatus, error) {
	ctx := context.Background()

	info, err := cf.GetHostnameInfo(ctx, hostname)
	if _, ok := err.(*cloudflare.NotFoundError); ok {
		info, err = cf.RegisterHostname(ctx, hostname)
	}
	if err != nil {
		return nil, err
	}

	var txtRecord *TXTRecordDetails
	if len(info.SSL.ValidationRecords) > 0 {
		rec := info.SSL.ValidationRecords[0]
		txtRecord = &TXTRecordDetails{
			Name:  rec.TxtName,
			Value: rec.TxtValue,
		}
	}
	res := &DomainStatus{
		Page:             page,
		TokenID:          tokenID,
		Hostname:         hostname,
		OwnershipStatus:  cloudflare.OwnerhsipStatus(info.Status),
		SSLStatus:        cloudflare.SSLVerificationStatus(info.SSL.Status),
		TXTRecordDetails: txtRecord,
	}
	return res, nil
}
