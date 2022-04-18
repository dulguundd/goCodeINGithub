package app

import (
	"github.com/gorilla/mux"
	"goCodeINGithub/jsonToXML/domain"
	"goCodeINGithub/jsonToXML/service"
	"log"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	ch := ConvertHandler{service.NewConvertService(domain.NewConvertRepository)}

	router.HandleFunc("/xmltojson", ch.XMLToJson).Methods(http.MethodGet)
	router.HandleFunc("/jsontoxml", ch.JsonToXML).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":10000", router))
}
