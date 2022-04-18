package app

import (
	"encoding/xml"
	//"encoding/json"
	json "github.com/goccy/go-json"
	"goCodeINGithub/jsonToXML/dto"
	"goCodeINGithub/jsonToXML/service"
	"net/http"
)

type ConvertHandler struct {
	service service.ConvertService
}

func (ch *ConvertHandler) XMLToJson(w http.ResponseWriter, r *http.Request) {
	inputXMLData := dto.XMLData{}
	_ = xml.NewDecoder(r.Body).Decode(&inputXMLData)
	//response , _ := ch.service.XMLToJson(inputXMLData)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
	if err := json.NewEncoder(w).Encode(inputXMLData); err != nil {
		panic(err)
	}
}

func (ch *ConvertHandler) JsonToXML(w http.ResponseWriter, r *http.Request) {
	inputJsonData := dto.JsonData{}
	_ = json.NewDecoder(r.Body).Decode(&inputJsonData)
	//response, _ := ch.service.JsonToXML(inputJsonData)
	w.Header().Add("Content-Type", "application/xml")
	w.WriteHeader(200)
	if err := xml.NewEncoder(w).Encode(inputJsonData); err != nil {
		panic(err)
	}
}
