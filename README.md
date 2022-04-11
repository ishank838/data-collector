# data-collector

A Data Aggregation API service which aggregates metrics based on country and device.

## API
It contains two API's as follows:

```
Insert URL: /v1/insert Method: POST
Query URL: /v1/query Method: POST
```

Insert accepts dimensions and metrics and insert them to the aggregator tree.

Query accepts the dimensions and output the metrics based on the queries dimensions.


## Build
### Docker Build
```
docker-compose up --build
```

### Go run

```
go run main.go
```

## Sample Requests
### Insert
```
{
    "dim":[
        {"key":"device","value":"mobile"},{"key":"country","value":"JP"}
    ],
    "metrics": [
        {"key":"webreq","value":40},{"key":"timespent","value":50}
    ]
}
```

### Query
```
{
    "dim":[
        {"key":"country","value":"JP"},{"key":"device","value":"mobile"}
    ]
}
```
