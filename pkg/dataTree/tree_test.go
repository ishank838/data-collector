package datatree_test

import (
	"fmt"
	"testing"

	"github.com/ishank838/data-collection/models"
	datatree "github.com/ishank838/data-collection/pkg/dataTree"
	"github.com/stretchr/testify/assert"
)

func getMetrics(resp *models.QueryResponse) (int64, int64) {

	Webreq := int64(0)
	timespent := int64(0)

	for _, v := range resp.Metrics {
		if v.Key == models.Webreq {
			Webreq += v.Value
		}

		if v.Key == models.Timespent {
			timespent += v.Value
		}
	}

	return Webreq, timespent
}

func TestInsert(t *testing.T) {

	tree := datatree.InitRoot()

	req := models.ParsedInsertRequest{
		Country:    "IN",
		Device:     "mobile",
		WebRequest: 50,
		TimeSpent:  80,
	}

	tree.Insert(req)

	queryReq := models.ParsedQueryRequest{
		Country: "IN",
		Device:  "mobile",
	}
	resp, _ := tree.Query(&queryReq)

	Webreq, timespent := getMetrics(resp)

	assert.Equal(t, req.WebRequest, Webreq)
	assert.Equal(t, req.TimeSpent, timespent)
}

func TestInsertTwo(t *testing.T) {

	tree := datatree.InitRoot()

	req := models.ParsedInsertRequest{
		Country:    "IN",
		Device:     "mobile",
		WebRequest: 50,
		TimeSpent:  80,
	}

	req2 := models.ParsedInsertRequest{
		Country:    "IN",
		Device:     "mobile",
		WebRequest: 40,
		TimeSpent:  40,
	}

	tree.Insert(req)
	tree.Insert(req2)

	queryReq := models.ParsedQueryRequest{
		Country: "IN",
		Device:  "mobile",
	}
	resp, _ := tree.Query(&queryReq)

	Webreq, timespent := getMetrics(resp)

	assert.Equal(t, req.WebRequest+req2.WebRequest, Webreq)
	assert.Equal(t, req.TimeSpent+req2.TimeSpent, timespent)
}

func TestInsertMultipleDevices(t *testing.T) {

	tree := datatree.InitRoot()

	req := models.ParsedInsertRequest{
		Country:    "IN",
		Device:     "mobile",
		WebRequest: 50,
		TimeSpent:  80,
	}

	req2 := models.ParsedInsertRequest{
		Country:    "IN",
		Device:     "tablet",
		WebRequest: 40,
		TimeSpent:  40,
	}

	tree.Insert(req)
	tree.Insert(req2)

	queryReq := models.ParsedQueryRequest{
		Country: "IN",
		Device:  "mobile",
	}
	resp, _ := tree.Query(&queryReq)

	Webreq := int64(0)
	timespent := int64(0)

	for _, v := range resp.Metrics {
		if v.Key == models.Webreq {
			Webreq += v.Value
		}

		if v.Key == models.Timespent {
			timespent += v.Value
		}
	}

	assert.Equal(t, req.WebRequest, Webreq)
	assert.Equal(t, req.TimeSpent, timespent)
}

func TestInsertMultipleDevices2(t *testing.T) {

	tree := datatree.InitRoot()

	req := models.ParsedInsertRequest{
		Country:    "IN",
		Device:     "mobile",
		WebRequest: 50,
		TimeSpent:  80,
	}

	req2 := models.ParsedInsertRequest{
		Country:    "IN",
		Device:     "tablet",
		WebRequest: 40,
		TimeSpent:  40,
	}

	tree.Insert(req)
	tree.Insert(req2)

	queryReq := models.ParsedQueryRequest{
		Country: "IN",
		Device:  "tablet",
	}
	resp, _ := tree.Query(&queryReq)

	Webreq := int64(0)
	timespent := int64(0)

	for _, v := range resp.Metrics {
		if v.Key == models.Webreq {
			Webreq += v.Value
		}

		if v.Key == models.Timespent {
			timespent += v.Value
		}
	}

	assert.Equal(t, req2.WebRequest, Webreq)
	assert.Equal(t, req2.TimeSpent, timespent)
}

func TestInsertMultipleCountries(t *testing.T) {

	tree := datatree.InitRoot()

	req := models.ParsedInsertRequest{
		Country:    "IN",
		Device:     "mobile",
		WebRequest: 50,
		TimeSpent:  80,
	}

	req2 := models.ParsedInsertRequest{
		Country:    "US",
		Device:     "mobile",
		WebRequest: 40,
		TimeSpent:  40,
	}

	tree.Insert(req)
	tree.Insert(req2)

	queryReq := models.ParsedQueryRequest{
		Country: "IN",
		Device:  "mobile",
	}
	resp, _ := tree.Query(&queryReq)

	Webreq, timespent := getMetrics(resp)

	assert.Equal(t, req.WebRequest, Webreq)
	assert.Equal(t, req.TimeSpent, timespent)

	queryReqUS := models.ParsedQueryRequest{
		Country: "US",
		Device:  "mobile",
	}
	resp2, _ := tree.Query(&queryReqUS)

	Webreq2, timespent2 := getMetrics(resp2)

	assert.Equal(t, req2.WebRequest, Webreq2)
	assert.Equal(t, req2.TimeSpent, timespent2)
}

func TestErrCountryNotFound(t *testing.T) {

	tree := datatree.InitRoot()

	req := models.ParsedInsertRequest{
		Country:    "IN",
		Device:     "mobile",
		WebRequest: 50,
		TimeSpent:  80,
	}

	tree.Insert(req)

	queryReq := models.ParsedQueryRequest{}

	_, err := tree.Query(&queryReq)

	assert.Equal(t, err, fmt.Errorf("country not found"))
}

func TestQueryCountry(t *testing.T) {

	tree := datatree.InitRoot()

	req := models.ParsedInsertRequest{
		Country:    "IN",
		Device:     "mobile",
		WebRequest: 50,
		TimeSpent:  80,
	}

	tree.Insert(req)

	queryReq := models.ParsedQueryRequest{
		Country: "IN",
	}

	resp, _ := tree.Query(&queryReq)

	web, timespent := getMetrics(resp)

	assert.Equal(t, req.WebRequest, web)
	assert.Equal(t, req.TimeSpent, timespent)
}
