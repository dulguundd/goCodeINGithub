package app

import (
	"bytes"
	"net/http"
	"testing"
)

func BenchmarkFuncJson(b *testing.B) {
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
	request.Header.Add("Accept", "application/json")
	for i := 0; i < b.N; i++ {
		_, _ = client.Do(request)
	}
}
