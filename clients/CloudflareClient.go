package clients

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/joho/godotenv"
	"os"
)

func GetCloudflareClient() (*cloudflare.API, error) {
	godotenv.Load()
	token := os.Getenv("CLOUDFLARE_TOKEN")
	email := os.Getenv("CLOUDFLARE_EMAIL")

	return cloudflare.New(token, email)
}
