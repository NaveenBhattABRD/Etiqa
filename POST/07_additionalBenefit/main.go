package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://uat.etiqa.com.my:4442/motorcar/api/additional-benefit"
	method := "POST"

	payload := strings.NewReader(`{` + "" + `"qqId":"206119",` + "" + `"code":"C57A,P10A",` + "" + `"codeWithSumInsured":"B089",` + "" + `"sumInsuredAdditionalBenefit":"1000"` + "" + `}`)

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
