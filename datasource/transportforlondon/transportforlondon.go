package transportforlondon

import (
	"fmt"
	"os"

	"github.com/golang-graphql/datasource/http"
)

var (
	url = "https://api.tfl.gov.uk/"
)

func httpQueryBuilder(params TypAPIParams) (header map[string]string, query map[string]string) {
	header = make(map[string]string)
	query = make(map[string]string)

	if params.Appkey != "" {
		query["app_key"] = params.Appkey
	}

	return

}

func GetStopPointFares() ([]StopPointFares, error) {
	url := fmt.Sprintf("%vStopPoint/940GZZLUASL/FareTo/940GZZLUVIC", url)
	var data []StopPointFares
	params := TypAPIParams{Appkey: os.Getenv("APP_KEY")}

	header, query := httpQueryBuilder(params)

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
	Appkey string
}

type FromStation struct {
	Type         *string `json:"$type"`
	AtcoCode     *string `json:"atcoCode"`
	CommonName   *string `json:"commonName"`
	FareCategory *string `json:"fareCategory"`
}

type Journey struct {
	Type        *string      `json:"$type"`
	FromStation *FromStation `json:"fromStation"`
	ToStation   *ToStation   `json:"toStation"`
}

type Messages struct {
	Type        *string `json:"$type"`
	BulletOrder *int    `json:"bulletOrder"`
	MessageText *string `json:"messageText"`
}

type Rows struct {
	Type                   *string             `json:"$type"`
	StartDate              *string             `json:"startDate"`
	EndDate                *string             `json:"endDate"`
	PassengerType          *string             `json:"passengerType"`
	ContactlessPAYOnlyFare *bool               `json:"contactlessPAYOnlyFare"`
	From                   *string             `json:"from"`
	To                     *string             `json:"to"`
	FromStation            *string             `json:"fromStation"`
	ToStation              *string             `json:"toStation"`
	DisplayName            *string             `json:"displayName"`
	DisplayOrder           *int                `json:"displayOrder"`
	RouteDescription       *string             `json:"routeDescription"`
	SpecialFare            *bool               `json:"specialFare"`
	ThroughFare            *bool               `json:"throughFare"`
	IsTour                 *bool               `json:"isTour"`
	TicketsAvailable       []*TicketsAvailable `json:"ticketsAvailable"`
	Messages               []*string           `json:"messages"`
}

type StopPointFares struct {
	Header   *string     `json:"header"`
	Index    *int        `json:"index"`
	Journey  *Journey    `json:"journey"`
	Rows     []*Rows     `json:"rows"`
	Messages []*Messages `json:"messages"`
}

type TicketTime struct {
	Type        *string `json:"$type"`
	Description *string `json:"description"`
}

type TicketType struct {
	Type        *string `json:"$type"`
	Description *string `json:"description"`
}

type TicketsAvailable struct {
	Type          *string     `json:"$type"`
	PassengerType *string     `json:"passengerType"`
	TicketType    *TicketType `json:"ticketType"`
	TicketTime    *TicketTime `json:"ticketTime"`
	Cost          *string     `json:"cost"`
	Description   *string     `json:"description"`
	Mode          *string     `json:"mode"`
	DisplayOrder  *int        `json:"displayOrder"`
	Messages      []*string   `json:"messages"`
}

type ToStation struct {
	Type         *string `json:"$type"`
	AtcoCode     *string `json:"atcoCode"`
	CommonName   *string `json:"commonName"`
	FareCategory *string `json:"fareCategory"`
}
