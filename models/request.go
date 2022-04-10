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
		if v.Key == Country {
			parsedRequest.Country = v.Value
		} else if v.Key == Device {
			parsedRequest.Device = v.Value
		}
	}
	parsedRequest.TimeSpent = 0
	parsedRequest.WebRequest = 0

	for _, v := range req.Metrics {

		if v.Key == Webreq {
			parsedRequest.WebRequest += v.Value
		} else if v.Key == Timespent {
			parsedRequest.TimeSpent += v.Value
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
		if v.Key == Country {
			parsedRequest.Country = v.Value
		}

		if v.Key == Device {
			parsedRequest.Device = v.Value
		}
	}
	return parsedRequest
}
