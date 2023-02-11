package cloudflare

import "fmt"

type NotFoundError struct {
	hostname string
}

func (err *NotFoundError) Error() string {
	return fmt.Sprintf("Hostname %s not found", err.hostname)
}
