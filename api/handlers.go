package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/renotion-xyz/backend/cloudflare"
	"github.com/renotion-xyz/backend/moonpay"
	"github.com/renotion-xyz/backend/moralis"
	"github.com/renotion-xyz/backend/web3"
)

func getListTokensHandler(mc *moralis.MoralisClient, token string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		owner := c.Params("owner")
		res, err := mc.GetNFTsByOwner(owner, token)
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

		res, err := getDomainStatus(cf, tokenID, metadata.Hostname, metadata.Page)
		if err != nil {
			return err
		}

		return c.JSON(res)
	}
}

func getListDomainsHandler(mc *moralis.MoralisClient, rnt *web3.Renotion, cf *cloudflare.CloudflareClient, token string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		owner := c.Params("owner")
		res, err := mc.GetNFTsByOwner(owner, token)
		if err != nil {
			return err
		}
		domains := make([]DomainStatus, 0, len(res.Result))
		ch := make(chan DomainStatus)
		for _, nft := range res.Result {
			go func(tokenID string) {
				metadata, _ := rnt.GetDomainMetadata(tokenID)
				info, _ := getDomainStatus(cf, tokenID, metadata.Hostname, metadata.Page)
				ch <- *info
			}(nft.TokenID)
		}
		for i := 0; i < len(res.Result); i++ {
			domains = append(domains, <-ch)
		}

		return c.JSON(DomainsListResponse{domains})
	}
}

func signMoonPayURLHandler(mp *moonpay.MoonPayClient) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request SignMoonPayURLRequest
		if err := c.BodyParser(&request); err != nil {
			return err
		}
		return c.JSON(SignMoonPayURLResponse{
			URL: mp.SignURL("", request.Wallet),
		})
	}
}
