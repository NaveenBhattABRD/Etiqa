package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://uat.etiqa.com.my:4442/motorcar/api/generateInvoiceNo"
	method := "POST"

	payload := strings.NewReader(`{` +
		"" +
		`"qqId":"209344",` +
		"" +
		`
		"paymentGatewayCode":"FPX",` +
		"" +
		`"amountPaid":"500",` +
		"" +
		`"productCode":"MT"` +
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
