package main

import (
	"encoding/json"
	"fmt"
	gojson "github.com/goccy/go-json"
	jsoniter "github.com/json-iterator/go"
	"log"
	"os"
	"time"
)

var jsonjoniter = jsoniter.ConfigCompatibleWithStandardLibrary

type Data2 struct {
	INACTENDD    string `json:"INACTENDD"`
	Error        string `json:"Error"`
	RETAILER     string `json:"RETAILER"`
	CLASS        string `json:"CLASS"`
	ACTENDD      string `json:"ACTENDD"`
	ADMINST      string `json:"ADMINST"`
	CREDITIVIOCE string `json:"CREDITIVIOCE"`
	CODE         string `json:"CODE"`
	PHONE        string `json:"PHONE"`
	RBAL         string `json:"RBAL"`
}

var TestData2 = []byte(`
{
	"INACTENDD": "01/01/2038",
	"Error": "0",
	"RETAILER": "2",
	"CLASS": "PRE_Hybrid_14900_N",
	"ACTENDD": "01/01/2038",
	"ADMINST": "1",
	"CREDITIVIOCE": "5490500",
	"CODE": "0",
	"PHONE": "94300048",
	"RBAL": "5490500"
}`)

var comData = Data2{
	INACTENDD:    "01/01/2038",
	Error:        "0",
	RETAILER:     "2",
	CLASS:        "PRE_Hybrid_14900_N",
	ACTENDD:      "01/01/2038",
	ADMINST:      "1",
	CREDITIVIOCE: "5490500",
	CODE:         "0",
	PHONE:        "94300048",
	RBAL:         "5490500",
}

func Funcjsoniter(in []byte) *Data2 {
	//start := time.Now()
	var data Data2
	if err := jsonjoniter.Unmarshal(in, &data); err != nil {
		fmt.Print(err)
		os.Exit(1)

	}
	//ServiceLatencyLogger(start)
	//fmt.Println(data)
	return &data
}

func Funcjson(in []byte) *Data2 {
	//start := time.Now()
	var data Data2
	if err := json.Unmarshal(in, &data); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	//ServiceLatencyLogger(start)
	//fmt.Println(data)
	return &data
}

func Funcgojson(in []byte) *Data2 {
	//start := time.Now()
	var data Data2
	if err := gojson.Unmarshal(in, &data); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	//ServiceLatencyLogger(start)
	//fmt.Println(data)
	return &data
}

func ServiceLatencyLogger(start time.Time) {
	elapsed := time.Since(start)
	logMessage := fmt.Sprintf("response latencie %s", elapsed)
	log.Println(logMessage)
}

func FuncjsonMarshal(in Data2) *[]byte {
	if marshaladData, err := json.Marshal(comData); err != nil {
		fmt.Print(err)
		os.Exit(1)
	} else {
		return &marshaladData
	}
	return nil
}

func FuncGojsonMarshal(in Data2) *[]byte {
	if marshaladData, err := gojson.Marshal(comData); err != nil {
		fmt.Print(err)
		os.Exit(1)
	} else {
		return &marshaladData
	}
	return nil
}

func FuncjsoniterMarshal(in Data2) *[]byte {
	if marshaladData, err := jsoniter.Marshal(comData); err != nil {
		fmt.Print(err)
		os.Exit(1)
	} else {
		return &marshaladData
	}
	return nil
}
