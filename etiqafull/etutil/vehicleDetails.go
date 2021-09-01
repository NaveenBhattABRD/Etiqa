package etutil

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func VehilceDetails() {

	url := "https://uat.etiqa.com.my:4442/motorcar/api/vehicleInfo?regNumber=BHQ9111"
	method := "POST"

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
