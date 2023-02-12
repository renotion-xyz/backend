package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/renotion-xyz/backend/cloudflare"
	"github.com/renotion-xyz/backend/moralis"
	"github.com/renotion-xyz/backend/web3"
)

type ApiServer struct {
	app *fiber.App
	mc  *moralis.MoralisClient
	cf  *cloudflare.CloudflareClient
	rnt *web3.Renotion
}

func (srv *ApiServer) Start(port string) {
	srv.app.Listen(fmt.Sprintf(":%s", port))
}

func NewServer(
	mc *moralis.MoralisClient,
	cf *cloudflare.CloudflareClient,
	rnt *web3.Renotion,
	token string,
) *ApiServer {
	app := fiber.New()
	app.Use(cors.New())

	api := app.Group("/api")
	api.Get("/tokens/:owner", getListTokensHandler(mc, token))
	api.Get("/domains/:owner", getListDomainsHandler(mc, rnt, cf, token))
	api.Get("/metadata/:tokenID", getDomainMetadataHandler(rnt))
	api.Get("/domain/:tokenID", getDomainStatusHandler(rnt, cf))

	return &ApiServer{app, mc, cf, rnt}
}
