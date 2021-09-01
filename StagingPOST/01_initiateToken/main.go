package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://staging.api.maybank.com/api/oauth2/v2/clientcred/token"
	method := "POST"

	payload := strings.NewReader(`{` +
		"" + `
    "grant_type": "client_credentials",` +
		"" + `
    "scope": "motors",` +
		"" + `
    "client_id": "EDCAF1A98B894A51AA1CA8C7348BCA1B",` +
		"" + `
    "client_secret": "0a1b77300cd8474ea1dae4401def78d9"` +
		"" + `
	}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("X-MB-Signed-Headers", "X-MB-Timestamp")
	req.Header.Add("X-MB-Signature-Alg", "RSA-SHA256")
	req.Header.Add("X-MB-Timestamp", "1622783254630")
	req.Header.Add("X-MB-Signature-Value", "wT1rWFamE7XKKJhYe5XieprNXzqvu7uWksiB3IitnFTyCJITNm6vVOjLj/VKA2OoGNxoAjBOakJzVd3z2b51JEcSoRPS/IHc8eMXWLd0+MiRVIbOu8TBGPvvWwFE+MABUvdOaNUwUkbUpNXxO3tk4zC68WkpOM6bborVROcbsVkIZAEJugwHW09i/q3/20NvelTt2ygUxdxAxq5e/UTovIm9KLNp5eG0ppqry+AE6u097UkqKVu0KVD/iXDElnETZbCmG+derU823sD+DZXKyolW1eDqeHjnQKVAzra2i6oLtbe2+Gc8oHHv01+tVofvWs00a46WS1gVvCCAX7kWvQ==")
	req.Header.Add("X-MB-E2E-ID", "APIGW1622783254630")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
