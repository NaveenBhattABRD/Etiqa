package etutil

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func InitiateTransaction() {

	url := "https://uat.etiqa.com.my:4442/motorcar/api/initiate"
	method := "POST"

	payload := strings.NewReader(`{` +
		"" +
		`
		"vehicleRegistrationNumber":"AJT3309",` +
		"" + `
		"productCode":"MT",` +
		"" +
		`"nric":"900818035263",` +
		"" +
		`"postcode":"09400",` +
		"" + `"agentCode":"V0003926"` +
		"" + `}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
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
