package clients

import (
	"github.com/cloudflare/cloudflare-go"
)

func GetCloudflareClient(email string, token string) (*cloudflare.API, error) {
	return cloudflare.New(token, email)
}
