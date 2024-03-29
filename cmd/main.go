package main

import (
	"context"
	"os"

	"github.com/renotion-xyz/backend/api"
	"github.com/renotion-xyz/backend/cloudflare"
	"github.com/renotion-xyz/backend/moonpay"
	"github.com/renotion-xyz/backend/moralis"
	"github.com/renotion-xyz/backend/web3"
)

var (
	PORT               = os.Getenv("PORT")
	CF_API_TOKEN       = os.Getenv("CF_API_TOKEN")
	RENOTION_CONTRACT  = os.Getenv("RENOTION_CONTRACT")
	RPC_URL            = os.Getenv("RPC_URL")
	MORALIS_API_KEY    = os.Getenv("MORALIS_API_KEY")
	TOKEN_ADDRESS      = os.Getenv("TOKEN_ADDRESS")
	ZONE_NAME          = os.Getenv("ZONE_NAME")
	CHAIN              = os.Getenv("CHAIN")
	MOONPAY_API_KEY    = os.Getenv("MOONPAY_API_KEY")
	MOONPAY_SECRET_KEY = os.Getenv("MOONPAY_SECRET_KEY")
	DEFAULT_CURRENCY   = os.Getenv("DEFAULT_CURRENCY")
)

func init() {
	if PORT == "" {
		PORT = "3000"
	}
}

func main() {
	ctx := context.Background()

	mc := moralis.NewClient(MORALIS_API_KEY, moralis.Chain(CHAIN))
	mp := moonpay.NewClient(MOONPAY_API_KEY, MOONPAY_SECRET_KEY, DEFAULT_CURRENCY)
	cf, err := cloudflare.NewClient(ctx, CF_API_TOKEN, ZONE_NAME)
	if err != nil {
		panic(err)
	}
	rnt, err := web3.NewRenotion(RPC_URL, RENOTION_CONTRACT)
	if err != nil {
		panic(err)
	}

	srv := api.NewServer(mc, cf, rnt, mp, TOKEN_ADDRESS)
	srv.Start(PORT)
}
