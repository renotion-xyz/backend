package cloudflare

import "github.com/cloudflare/cloudflare-go"

type OwnerhsipStatus cloudflare.CustomHostnameStatus

// https://github.com/cloudflare/cloudflare-go/issues/131
type SSLVerificationStatus string

const (
	PENDING_VALIDATION SSLVerificationStatus = "pending_validation"
	PENDING_ISSUANCE   SSLVerificationStatus = "pending_issuance"
	PENDING_DEPLOYMENT SSLVerificationStatus = "pending_deployment"
	ACTIVE             SSLVerificationStatus = "active"
	DELETED            SSLVerificationStatus = "deleted"
	INITIALIZING       SSLVerificationStatus = "initializing"
)
