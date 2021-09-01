package etutil

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func GetPremium() {

	url := "https://uat.etiqa.com.my:4442/motorcar/api/premium"
	method := "POST"

	payload := strings.NewReader(`{` +
		"" + `"vehicleRegistrationNumber":"WXB4857",` +
		"" + `"productCode":"MT",` +
		"" + `"nric":"760725075088",` +
		"" + `"postcode":"56000",` +
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
