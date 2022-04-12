package main

import "fmt"

func main() {
	fmt.Println("Encoding/json decoding time:")
	stdjson()
	fmt.Println("Gojson decoding time:")
	funcgojson()
	fmt.Println("Jsoniter decoding time:")
	funcjsoniter()
}
