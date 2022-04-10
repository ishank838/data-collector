package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ishank838/data-collection/models"
	datatree "github.com/ishank838/data-collection/pkg/dataTree"
)

func Insert(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var request models.InsertRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	paresedReq := request.ParseInsertRequest()
	log.Println(paresedReq)

	datatree.Insert(paresedReq)

	w.WriteHeader(http.StatusOK)
}

func Query(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var request models.QueryRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	parsedReq := request.ParseQueryRequest()

	resp, err := datatree.Query(&parsedReq)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//type,name := request.GetQueryName()
	respJson, _ := json.Marshal(resp)
	w.Write(respJson)
}
