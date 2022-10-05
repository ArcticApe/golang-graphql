package transportforlondon

import (
	"fmt"
	"os"

	"github.com/golang-graphql/datasource/http"
	"github.com/golang-graphql/graph/model"
)

var (
	url = "https://api.tfl.gov.uk/"
)

func httpQueryBuilder(authtoken string, params TflAPIParams) (header map[string]string, query map[string]string) {
	header = make(map[string]string)
	query = make(map[string]string)

	if params.appkey != "" {
		query["app_key"] = params.appkey
	}

	return

}

func GetStopPointFares(params TflAPIParams) (model.StopPointFares, error) {
	url := fmt.Sprintf("%v/StopPoint/940GZZLUASL/FareTo/940GZZLUVIC", url)
	var data model.StopPointFares

	header, query := httpQueryBuilder(os.Getenv("APP_KEY"), params)

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

type TflAPIParams struct {
	appkey string
}
