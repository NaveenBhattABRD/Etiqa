package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://staging.api.maybank.com/api/my/retail/insurance/v1/motorcar/txn/sum-insured?bizcode=etq"
	method := "POST"

	payload := strings.NewReader(`{` +
		"" + `
    	"sumInsured":"23000",` +
		"" + `
    	"qqId":"204360"` +
		"" +
		`}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("X-MB-Signed-Headers", "X-MB-Client-Id;Authorization;XMB-Timestamp")
	req.Header.Add("X-MB-Signature-Alg", "RSA-SHA256")
	req.Header.Add("X-MB-Timestamp", "1621210235938")
	req.Header.Add("X-MB-Signature-Value", "wFclVrkmBUVb7bSPXwZ4QPkYZVB17eRvb5438zUCML5qsJDk/YU14m7g6eVy+mYNvEFHz09TTXlxOUtjsxXpjLG009MBOMS0DBet3jchtjQSOilt7WQ1xjKRih2XemteO3LqnzalMOP8jJdqakNxB+W83mmEYW7YXtxvSzIJMhxExp6cUF1XAldbYRCp8kglrcE9lpDg464wVC4OE2k93qPcWgUKaD+dwHy74Tsy1QPm0W/Csb+0C/+5joyY2LXxjlArzTvMjgu3NlN7XA+31N8hgjBaVPH4C2C5pux50hFgtvyXR10lDhp+sZi699RLxBcnh9gwDupsK4D6Li4F6g==")
	req.Header.Add("X-MB-E2E-Id", "APIGW1621210235938")
	req.Header.Add("X-MB-Client-Id", "EDCAF1A98B894A51AA1CA8C7348BCA1B")

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
