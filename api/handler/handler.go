package handler

import (
	"encoding/json"
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

	if paresedReq.Country == "" || paresedReq.Device == "" || paresedReq.TimeSpent == 0 || paresedReq.WebRequest == 0 {
		w.WriteHeader(http.StatusBadRequest)
		msg := map[string]string{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  "incomplete request",
		}
		m, _ := json.Marshal(msg)
		w.Write(m)
		return
	}

	t := datatree.GetTree()

	t.Insert(paresedReq)

	w.WriteHeader(http.StatusOK)
	msg := map[string]string{
		"status": http.StatusText(http.StatusOK),
	}
	m, _ := json.Marshal(msg)
	w.Write(m)
}

func Query(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var request models.QueryRequest
	err := json.NewDecoder(r.Body).Decode(&request)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := map[string]string{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  "invalid request",
		}
		m, _ := json.Marshal(msg)
		w.Write(m)
		return
	}

	parsedReq := request.ParseQueryRequest()

	t := datatree.GetTree()

	resp, err := t.Query(&parsedReq)

	if err != nil {
		//log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		msg := map[string]string{
			"status": http.StatusText(http.StatusBadRequest),
			"error":  "required country or devices doesn't exist",
		}
		m, _ := json.Marshal(msg)
		w.Write(m)
		return
	}

	respJson, _ := json.Marshal(resp)
	w.Write(respJson)
}
