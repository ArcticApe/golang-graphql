package electricitymap

import (
	"fmt"
	"os"

	"github.com/golang-graphql/datasource/http"
	"github.com/golang-graphql/graph/model"

	"strconv"
)

var (
	url = "https://api.electricitymap.org/v3"
)

func httpQueryBuilder(authtoken string, params TypAPIParams) (header map[string]string, query map[string]string) {
	header = make(map[string]string)
	query = make(map[string]string)

	header["auth-token"] = authtoken

	if params.Zone != "" {
		query["zone"] = params.Zone
	}
	if params.Lon != "" && params.Lat != "" {
		query["lon"] = params.Lon
		query["lat"] = params.Lat
	}
	if params.Datetime != "" {
		query["datetime"] = params.Datetime
	}
	if params.Start != "" {
		query["start"] = params.Start
	}
	if params.End != "" {
		query["end"] = params.End
	}
	if params.EstimationFallback {
		query["estimationFallback"] = strconv.FormatBool(params.EstimationFallback)
	}

	return
}

/*
This endpoint retrieves the last known carbon intensity (in gCO2eq/kWh) of electricity consumed in an area. It can either be queried by zone identifier or by geolocation.
QUERY PARAMETERS
Parameter | Description
zone | A string representing the zone identifier
lon | Longitude (if querying with a geolocation)
lat | Latitude (if querying with a geolocation)
*/
func GetLiveCarbonIntensity(params TypAPIParams) (model.CarbonIntensity, error) {
	url := fmt.Sprintf("%v/carbon-intensity/latest", url)
	var data model.CarbonIntensity

	header, query := httpQueryBuilder(os.Getenv("AUTH_TOKEN"), params)

	request := http.Request{
		Url:      url,
		Method:   "GET",
		Header:   header,
		Query:    query,
		Response: &data,
	}

	err := request.Send()
	return data, err
}

func GetLivePowerBreakdown(params TypAPIParams) (model.PowerBreakdown, error) {
	url := fmt.Sprintf("%v/power-breakdown/latest", url)
	var data model.PowerBreakdown

	header, query := httpQueryBuilder(os.Getenv("AUTH_TOKEN"), params)

	request := http.Request{
		Url:      url,
		Method:   "GET",
		Header:   header,
		Query:    query,
		Response: &data,
	}

	err := request.Send()
	return data, err
}

type TypAPIParams struct {
	Zone               string
	Lon                string
	Lat                string
	Datetime           string
	Start              string
	End                string
	EstimationFallback bool
}
