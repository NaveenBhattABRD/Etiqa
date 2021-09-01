package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://uat.etiqa.com.my:4442/motorcar/api/personal"
	method := "POST"

	payload := strings.NewReader(`{` + "" + `"name":"Adam Arif",` + "" + `"genderCode":"001",` + "" + `"maritalStatusCode":"001",` + "" + `"email":"adamarif@gmail.com",` + "" + `"mobileNumber":"0123456789",` + "" + `"address1":"Level 8B, Dataran Maybank,",` + "" + `"address2":"Jalan Maarof",` + "" + `"address3":"",` + "" + `"postcode":"59000",` + "" + `"qqId":"197255"` + "" + `}`)

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
