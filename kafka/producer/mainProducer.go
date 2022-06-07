package main

import (
	"encoding/json"
	"fmt"
	"github.com/dulguundd/logError-lib/logger"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Comment struct {
	Text string `form:"text" json:"text"`
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/comment", createComment).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("0.0.0.0:10051", router))
}

func createComment(w http.ResponseWriter, r *http.Request) {
	cmt := new(Comment)

	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		cmtInBytes, err := json.Marshal(cmt)

		PushCommentToQueue("go-comments", cmtInBytes)

		if err != nil {
			writeResponse(w, http.StatusInternalServerError, err.Error())
		} else {
			writeResponse(w, http.StatusCreated, nil)
		}
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func serviceLatencyLogger(start time.Time) {
	elapsed := time.Since(start)
	logMessage := fmt.Sprintf("response latencie %s", elapsed)
	logger.Info(logMessage)
}
