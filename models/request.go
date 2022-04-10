package models

type InsertRequest struct {
	Dim []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"dim"`
	Metrics []struct {
		Key   string `json:"key"`
		Value int64  `json:"value"`
	}
}

type ParsedInsertRequest struct {
	Country    string
	Device     string
	WebRequest int64
	TimeSpent  int64
}

func (req *InsertRequest) ParseInsertRequest() ParsedInsertRequest {
	var parsedRequest ParsedInsertRequest

	for _, v := range req.Dim {
		if v.Key == "country" {
			parsedRequest.Country = v.Value
		} else if v.Key == "device" {
			parsedRequest.Device = v.Value
		}
	}

	for _, v := range req.Metrics {

		if v.Key == "webreq" {
			parsedRequest.WebRequest = v.Value
		} else if v.Key == "timespent" {
			parsedRequest.TimeSpent = v.Value
		}
	}

	return parsedRequest
}

type QueryRequest struct {
	Dim []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"dim"`
}

type QueryResponse struct {
	Dimensions []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"dim"`
	Metrics []struct {
		Key   string `json:"key"`
		Value int64  `json:"value"`
	} `json:"metrics"`
}

type ParsedQueryRequest struct {
	Country string
	Device  string
}

func (req *QueryRequest) ParseQueryRequest() ParsedQueryRequest {

	var parsedRequest ParsedQueryRequest

	for _, v := range req.Dim {
		if v.Key == "country" {
			parsedRequest.Country = v.Value
		}

		if v.Key == "device" {
			parsedRequest.Device = v.Value
		}
	}
	return parsedRequest
}
