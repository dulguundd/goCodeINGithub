package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	body := []byte(`
<XMLData>
    <channel>
        <item>
            <INACTENDD>01/01/2038</INACTENDD>
            <Error></Error>
            <RETAILER>2</RETAILER>
            <CLASS>PRE_Hybrid_14900_N</CLASS>
            <ACTENDD>01/01/2038</ACTENDD>
            <ADMINST>1</ADMINST>
            <CREDITVIOCE>5490500</CREDITVIOCE>
            <CODE>0</CODE>
            <PHONE>94300048</PHONE>
            <RBAL>5490500</RBAL>
        </item>
    </channel>
</XMLData>`)
	client := &http.Client{}
	request, _ := http.NewRequest(http.MethodGet, "http://172.22.2.215:10000/xmltojson", bytes.NewBuffer(body))

	resp, err := client.Do(request)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))
}
