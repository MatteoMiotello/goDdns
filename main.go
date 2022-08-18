package main

import (
	"context"
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"github.com/joho/godotenv"
	"goRdns/clients"
	"os"
)

func main() {
	godotenv.Load()

	currentIp, err := clients.GetCurrentPublicIp()

	if err != nil {
		fmt.Printf("Not able to get public IP: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Current IP: %s\n", currentIp)

	api, err := clients.GetCloudflareClient()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	zoneId, err := api.ZoneIDByName(os.Getenv("ZONE_NAME"))

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("zoneID: %s\n", zoneId)

	ctx := context.Background()

	record := cloudflare.DNSRecord{
		Name: os.Getenv("DNS_RECORD"),
	}

	recordsDetails, err := api.DNSRecords(ctx, zoneId, record)
	recordDetails := recordsDetails[0]

	fmt.Printf("id: %s\n", recordDetails.ID)

	recordDetails.Content = currentIp

	updateErr := api.UpdateDNSRecord(ctx, zoneId, recordDetails.ID, recordDetails)

	if updateErr != nil {
		fmt.Printf("Error updating record: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("DNS updated successfully")
}
