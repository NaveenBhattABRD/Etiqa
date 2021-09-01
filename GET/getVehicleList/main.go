package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	url := "https://uat.etiqa.com.my:4442/motorcar/api/vehicle?classCode=02&makeCode=33&modelCode=06&year=2017"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

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
