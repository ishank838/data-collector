package datatree

import (
	"fmt"

	"github.com/ishank838/data-collection/models"
)

var Root = initRoot()

func initRoot() *Node {
	return &Node{
		Type:       RootType,
		Name:       "root",
		Children:   make(map[string]*Node),
		TimeSpent:  0,
		WebRequest: 0,
	}
}

func Insert(req models.ParsedInsertRequest) {

	//update root metrics
	UpdateMetrics(Root, req.WebRequest, req.TimeSpent)

	//update country metrics
	countries := Root.Children
	reqCountry, ok := countries[req.Country]

	if !ok {
		node := NewNode(CountryType)
		node.Name = req.Country
		UpdateMetrics(node, req.WebRequest, req.TimeSpent)
		countries[req.Country] = node
	} else {
		UpdateMetrics(reqCountry, req.WebRequest, req.TimeSpent)
	}

	//update device metrics

	devices := countries[req.Country]
	reqDevice, ok := devices.Children[req.Device]

	if !ok {
		node := NewNode(DeviceType)
		node.Name = req.Device
		UpdateMetrics(node, req.WebRequest, req.TimeSpent)
		devices.Children[req.Device] = node
	} else {
		UpdateMetrics(reqDevice, req.WebRequest, req.TimeSpent)
	}
}

func Query(req *models.ParsedQueryRequest) (*models.QueryResponse, error) {

	countries := Root.Children

	reqCountry, ok := countries[req.Country]
	if !ok {
		return nil, fmt.Errorf("country not found")
	}

	var response models.QueryResponse

	response.Dimensions = append(response.Dimensions,
		struct {
			Key   string "json:\"key\""
			Value string "json:\"value\""
		}{Key: models.Country, Value: reqCountry.Name})

	if req.Device == "" {

		if reqCountry.WebRequest != 0 {
			response.Metrics = append(response.Metrics,
				struct {
					Key   string "json:\"key\""
					Value int64  "json:\"value\""
				}{Key: models.Webreq, Value: reqCountry.WebRequest},
			)
		}
		if reqCountry.TimeSpent != 0 {
			response.Metrics = append(response.Metrics,
				struct {
					Key   string "json:\"key\""
					Value int64  "json:\"value\""
				}{Key: models.Timespent, Value: reqCountry.TimeSpent},
			)
		}

		return &response, nil
	}

	devices := reqCountry.Children
	d, ok := devices[req.Device]

	if !ok {
		return nil, fmt.Errorf("required device doesn't exist")
	}
	response.Dimensions = append(response.Dimensions,
		struct {
			Key   string "json:\"key\""
			Value string "json:\"value\""
		}{Key: models.Device, Value: d.Name})

	if d.WebRequest != 0 {
		response.Metrics = append(response.Metrics,
			struct {
				Key   string "json:\"key\""
				Value int64  "json:\"value\""
			}{
				Key:   models.Webreq,
				Value: d.WebRequest,
			},
		)
	}
	if d.TimeSpent != 0 {
		response.Metrics = append(response.Metrics,
			struct {
				Key   string "json:\"key\""
				Value int64  "json:\"value\""
			}{
				Key:   models.Timespent,
				Value: d.TimeSpent,
			},
		)
	}

	return &response, nil
}

func UpdateMetrics(node *Node, webReq int64, timeSpent int64) {
	node.TimeSpent += timeSpent
	node.WebRequest += webReq
}
