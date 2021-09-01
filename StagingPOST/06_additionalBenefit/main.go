package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "https://staging.api.maybank.com/api/my/retail/insurance/v1/motorcar/txn/additional-benefit?bizcode=etq"
	method := "POST"

	payload := strings.NewReader(`{` +
		"" + `
	  "qqId":"204343",` +
		"" + `
	  "code":"C57A,P10A",` +
		"" + `
  	"codeWithSumInsured":"B089",` +
		"" + `
  	"sumInsuredAdditionalBenefit":"1000"` +
		"" + `}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("X-MB-Signed-Headers", "X-MB-Client-Id;Authorization;XMB-Timestamp")
	req.Header.Add("X-MB-Signature-Alg", "RSA-SHA256")
	req.Header.Add("X-MB-Timestamp", "1621160114837")
	req.Header.Add("X-MB-Signature-Value", "V6YtbFjUCnpsoNfzsby8ZO23UBOHZCCq4IT/iztYLvJbTgOGvXnDh5Tx1JcH6IJoGAwdzRT5JRG4WhM+Sc6sJKICXttEo7H68WrAbGEg+2cpYCHoEOJ+/1gYvI9KJzNHGcp4xIh81VznfNUzkWjBlJMOGbHVHDYcJvZud/Fn5Vl6yyPdwz8Fj2ZSERxhiLrXfiqnvt5F47bCtuxr2t7Wx6d0DykdXm4cbx3yV+xxsCmyQKRnfI9fE82Ud7gzeqiZp0+ndLnjRp9W6TGlrSvwnZyguYvWPHZJEshpBKyoczaVP51SY45Ovdi0RTYzuoThVHAOugV0Vl2/cWCoDYtv7A==")
	req.Header.Add("X-MB-E2E-Id", "APIGW1621160114837")
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
