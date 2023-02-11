package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/renotion-xyz/backend/cloudflare"
	"github.com/renotion-xyz/backend/moralis"
	"github.com/renotion-xyz/backend/web3"
)

func getListTokensHandler(mc *moralis.MoralisClient, token string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		owner := c.Params("owner")
		res, err := mc.GetNFTsByOwner(owner, token, moralis.POLYGON)
		if err != nil {
			return err
		}
		return c.JSON(ListTokensResponse{res.Result})
	}
}

func getDomainMetadataHandler(rnt *web3.Renotion) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenID := c.Params("tokenID")
		metadata, err := rnt.GetDomainMetadata(tokenID)
		if err != nil {
			return err
		}
		res := DomainMetadataResponse{
			TokenID: tokenID,
		}
		if metadata.Hostname != "" {
			res.Hostname = &metadata.Hostname
		}
		if metadata.Page != "" {
			res.Page = &metadata.Page
		}
		return c.JSON(res)
	}
}

func getDomainStatusHandler(rnt *web3.Renotion, cf *cloudflare.CloudflareClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenID := c.Params("tokenID")
		metadata, err := rnt.GetDomainMetadata(tokenID)
		if err != nil {
			return err
		}

		if metadata.Hostname == "" {
			return fiber.ErrNotFound
		}

		ctx := context.Background()

		info, err := cf.GetHostnameInfo(ctx, metadata.Hostname)
		if _, ok := err.(*cloudflare.NotFoundError); ok {
			info, err = cf.RegisterHostname(ctx, metadata.Hostname)
		}
		if err != nil {
			return err
		}

		var txtRecord *TXTRecordDetails
		if len(info.SSL.ValidationRecords) > 0 {
			rec := info.SSL.ValidationRecords[0]
			txtRecord = &TXTRecordDetails{
				Name:  rec.TxtName,
				Value: rec.TxtValue,
			}
		}
		res := CustomHostnameInfoResponse{
			Hostname:         metadata.Hostname,
			OwnershipStatus:  cloudflare.OwnerhsipStatus(info.Status),
			SSLStatus:        cloudflare.SSLVerificationStatus(info.SSL.Status),
			TXTRecordDetails: txtRecord,
		}

		return c.JSON(res)
	}
}
