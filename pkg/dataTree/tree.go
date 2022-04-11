package datatree

import (
	"fmt"
	"sync"

	"github.com/ishank838/data-collection/models"
)

type Tree struct {
	Root  *Node
	Mutex sync.Mutex
}

var treeInstance = InitRoot()

func GetTree() *Tree {
	return treeInstance
}

func InitRoot() *Tree {
	node := &Node{
		Type:       RootType,
		Name:       "root",
		Children:   make(map[string]*Node),
		TimeSpent:  0,
		WebRequest: 0,
	}

	return &Tree{Root: node}
}

func (t *Tree) Insert(req models.ParsedInsertRequest) {

	var updateNodes []*Node

	//update root metrics
	updateNodes = append(updateNodes, t.Root)
	//UpdateMetrics(t.Root, req.WebRequest, req.TimeSpent)

	//update country metrics
	countries := t.Root.Children
	reqCountry, ok := countries[req.Country]

	if !ok {
		node := NewNode(CountryType)
		node.Name = req.Country
		updateNodes = append(updateNodes, node)
		//UpdateMetrics(node, req.WebRequest, req.TimeSpent)
		countries[req.Country] = node
	} else {
		updateNodes = append(updateNodes, reqCountry)
		//UpdateMetrics(reqCountry, req.WebRequest, req.TimeSpent)
	}

	//update device metrics

	devices := countries[req.Country]
	reqDevice, ok := devices.Children[req.Device]

	if !ok {
		node := NewNode(DeviceType)
		node.Name = req.Device
		updateNodes = append(updateNodes, node)
		//UpdateMetrics(node, req.WebRequest, req.TimeSpent)
		devices.Children[req.Device] = node
	} else {
		updateNodes = append(updateNodes, reqDevice)

		//UpdateMetrics(reqDevice, req.WebRequest, req.TimeSpent)
	}

	t.UpdateMetrics(updateNodes, req.WebRequest, req.TimeSpent)
}

func (t *Tree) Query(req *models.ParsedQueryRequest) (*models.QueryResponse, error) {

	countries := t.Root.Children

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

func (t *Tree) UpdateMetrics(nodes []*Node, webReq int64, timeSpent int64) {

	t.Mutex.Lock()
	defer t.Mutex.Unlock()

	for _, node := range nodes {
		node.TimeSpent += timeSpent
		node.WebRequest += webReq
	}
}
