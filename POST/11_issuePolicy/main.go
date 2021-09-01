package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://uat.etiqa.com.my:4442/motorcar/api/getPolicy"
	method := "POST"

	payload := strings.NewReader(`{` +
		"" + `"qqId":"209916",` +
		"" + `"transactionId":"317131",` +
		"" + `"paymentGatewayCode":"FPX",` +
		"" + `"amountPaid":"520",` +
		"" + `"referenceNo": "NO0123455NO",` +
		"" + `"paymentStatus":"SUCCESS",` +
		"" + `"transactionDateTime":"12/07/2021",` +
		"" + `"preInsuranceCode": "098"` +
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
