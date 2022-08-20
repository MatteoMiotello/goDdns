package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"goRdns/clients"
	"os"
)

func main() {
	name := flag.String("zone", "test.cm", "zone name")
	dnsName := flag.String("record", "test", "record name")
	email := flag.String("email", "test@gmail.com", "cloudflare email")
	token := flag.String("token", "", "cloudflare token")

	flag.Parse()

	var zoneName string = *name
	var dnsRecordName string = *dnsName

	currentIp, err := clients.GetCurrentPublicIp()

	if err != nil {
		fmt.Printf("Not able to get public IP: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Current IP: %s\n", currentIp)

	api, err := clients.GetCloudflareClient(*email, *token)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	zoneId, err := api.ZoneIDByName(zoneName)

	if err != nil {
		os.Exit(1)
	}

	fmt.Printf("zoneID: %s\n", zoneId)

	ctx := context.Background()

	record := cloudflare.DNSRecord{
		Name: dnsRecordName + "." + zoneName,
		Type: "A",
	}

	recordsDetails, err := api.DNSRecords(ctx, zoneId, record)
	recordDetails := recordsDetails[0]

	fmt.Printf("id: %s\n", recordDetails.ID)

	if recordDetails.Content == currentIp {
		fmt.Printf("same ip of current\n")
		os.Exit(0)
	}

	recordDetails.Content = currentIp

	updateErr := api.UpdateDNSRecord(ctx, zoneId, recordDetails.ID, recordDetails)

	if updateErr != nil {
		fmt.Printf("Error updating record: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("DNS updated successfully\n")
}
