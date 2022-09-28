package electricitymap

import (
	"fmt"
	"os"

	"github.com/golang-graphql/datasource/http"

	"strconv"
)

var (
	url = "https://api.electricitymap.org/v3"
)

func httpQueryBuilder(zoneKey string, params TypAPIParams) (header map[string]string, query map[string]string) {
	header = make(map[string]string)
	query = make(map[string]string)

	header["auth-token"] = zoneKey

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
	if params.EstimationFallback == true {
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
func GetLiveCarbonIntensity(params TypAPIParams) (TypCI, error) {
	url := fmt.Sprintf("%v/carbon-intensity/latest", url)
	var data TypCI

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

type TypCI struct {
	Zone            string `json:"zone"`
	CarbonIntensity int    `json:"carbonIntensity"`
	Datetime        string `json:"datetime"`
	UpdatedAt       string `json:"updatedAt"`
	CreatedAt       string `json:"createdAt"`
}

type TypPB struct {
	Zone                      string                       `json:"zone"`
	Datetime                  string                       `json:"datetime"`
	PowerProductionBreakdown  TypPowerProductionBreakdown  `json:"powerProductionBreakdown"`
	PowerProductionTotal      int                          `json:"powerProductionTotal"`
	PowerConsumptionBreakdown TypPowerConsumptionBreakdown `json:"powerConsumptionBreakdown"`
	PowerConsumptionTotal     int                          `json:"powerConsumptionTotal"`
	PowerImportBreakdown      TypPowerImpExpBreakdown      `json:"powerImportBreakdown"`
	PowerImportTotal          int                          `json:"powerImportTotal"`
	PowerExportBreakdown      TypPowerImpExpBreakdown      `json:"powerExportBreakdown"`
	PowerExportTotal          int                          `json:"powerExportTotal"`
	FossilFreePercentage      int                          `json:"fossilFreePercentage"`
	RenewablePercentage       int                          `json:"renewablePercentage"`
	UpdatedAt                 string                       `json:"updatedAt"`
	CreatedAt                 string                       `json:"createdAt"`
}

type TypPowerConsumptionBreakdown struct {
	BatteryDischarge string // battery discharge `json:"batteryDischarge"`
	Biomass          int    `json:"biomass"`
	Coal             int    `json:"coal"`
	Gas              int    `json:"gas"`
	Geothermal       int    `json:"geothermal"`
	Hydro            int    `json:"hydro"`
	HydroDischarge   int    //hydro discharge `json:"hydroDischarge"`
	Nuclear          int    `json:"nuclear"`
	Oil              int    `json:"oil"`
	Solar            int    `json:"solar"`
	Unknown          int    `json:"unknown"`
	Wind             int    `json:"wind"`
}

type TypPowerImpExpBreakdown struct {
	DE     int `json:"DE"`
	DK_DK1 int //DK-DK1 `json:"DK_DK1"`
	SE     int `json:"SE"`
}

type TypPowerProductionBreakdown struct {
	Biomass    int `json:"biomass"`
	Coal       int `json:"coal"`
	Gas        int `json:"gas"`
	Geothermal int `json:"geothermal"`
	Hydro      int `json:"hydro"`
	Nuclear    int `json:"nuclear"`
	Oil        int `json:"oil"`
	Solar      int `json:"solar"`
	Unknown    int `json:"unknown"`
	Wind       int `json:"wind"`
}

type TypZone struct {
	CountryName string   `json:"countryName"`
	ZoneName    string   `json:"zoneName"`
	Access      []string `json:"access"`
}

type TypRecentCI struct {
	Zone    string `json:"zone"`
	History []struct {
		CarbonIntensity int    `json:"carbonIntensity"`
		Datetime        string `json:"datetime"`
		UpdatedAt       string `json:"updatedAt"`
		CreatedAt       string `json:"createdAt"`
	} `json:"history"`
}

type TypRecentPB struct {
	Zone    string `json:"zone"`
	History []struct {
		Datetime                  string                       `json:"datetime"`
		FossilFreePercentage      string                       `json:"fossilFreePercentage"`
		PowerConsumptionBreakdown TypPowerConsumptionBreakdown `json:"powerConsumptionBreakdown"`
		PowerConsumptionTotal     int                          `json:"powerConsumptionTotal"`
		PowerImportBreakdown      TypPowerImpExpBreakdown      `json:"powerImportBreakdown"`
		PowerImportTotal          int                          `json:"powerImportTotal"`
		PowerExportBreakdown      TypPowerImpExpBreakdown      `json:"powerExportBreakdown"`
		PowerExportTotal          int                          `json:"powerExportTotal"`
		PowerProductionBreakdown  TypPowerProductionBreakdown  `json:"powerProductionBreakdown"`
		PowerProductionTotal      int                          `json:"powerProductionTotal"`
		RenewablePercentage       int                          `json:"renewablePercentage"`
	} `json:"history"`
}
