package clients

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetCurrentPublicIp() (string, error) {
	url := "https://api.ipify.org?format=text"
	fmt.Printf("Getting IP address from  ipify\n")
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(ip), nil
}
