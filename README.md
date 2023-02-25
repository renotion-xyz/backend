# Renotion API backend

The backend which server the [Renotion web3 app](https://github.com/renotion-xyz/web3-app).

It checks on-chain state (from Polygon), and then fetches relevant metadata, and asserts Cloudflare Custom Domains registration.

Successful registrations are then forwarded to the [Renotion CF Worker](https://github.com/renotion-xyz/cf-worker) which proxies Notion pages based on blockchain state.

## Development

It relies on the deployed [Renotion Smart Contract](https://github.com/renotion-xyz/contracts).

### Setup

1. First `cp .env.example .env` – and then setup your environment
2. `go run cmd/main.go` to run the fiber webapp

## Deployment

### Secrets

First, you need to setup the secrets for the fly.io app:
```sh
flyctl secrets set \
  CF_API_TOKEN=$CF_API_TOKEN \
  MORALIS_API_KEY=$MORALIS_API_KEY \
  ZONE_NAME=$ZONE_NAME \
  CHAIN=$CHAIN \
  RPC_URL=$RPC_URL \
  RENOTION_CONTRACT=$RENOTION_CONTRACT \
  TOKEN_ADDRESS=$TOKEN_ADDRESS
```

### Test

```sh
flyctl deploy -a renotion-test
```

### Production

```sh
flyctl deploy
```

## API

- `GET /api/tokens/:owner`  
Fetches page registrations owned by an account
- `GET /api/domains/:owner`  
Fetches page registrations from chain, and then asserts Cloudflare domain associations.
- `GET /api/metadata/:tokenID`  
Reads metadata from chain, currently only hostname and an associated page
- `GET /api/domain/:tokenID`  
Fetches page registrations from chain, and then asserts Cloudflare domain association for the given registration.