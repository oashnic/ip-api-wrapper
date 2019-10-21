package ipapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// Geolocation response struct
type Geolocation struct {
	Query         string  `json:"query"`
	Status        string  `json:"status"`
	Continent     string  `json:"continent"`
	ContinentCode string  `json:"continentCode"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"countryCode"`
	Region        string  `json:"region"`
	RegionName    string  `json:"regionName"`
	City          string  `json:"city"`
	District      string  `json:"district"`
	Zip           string  `json:"zip"`
	Lat           float64 `json:"lat"`
	Lon           float64 `json:"lon"`
	Timezone      string  `json:"timezone"`
	Currency      string  `json:"currency"`
	Isp           string  `json:"isp"`
	Org           string  `json:"org"`
	AS            string  `json:"as"`
	ASname        string  `json:"asname"`
	Reverse       string  `json:"reverse"`
	Mobile        bool    `json:"mobile"`
	Proxy         bool    `json:"proxy"`
}

// GetLocation get geolocation by IP in localization
func GetLocation(ip string, localization string) (Geolocation, error) {
	resp, err := http.DefaultClient.Get("http://ip-api.com/json/" + ip + "?fields=status,message,continent,continentCode,country,countryCode,region,regionName,city,district,zip,lat,lon,timezone,currency,isp,org,as,asname,reverse,mobile,proxy,query&lang=" + localization)
	if err != nil {
		return Geolocation{}, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Geolocation{}, err
	}

	var geo Geolocation

	if err = json.Unmarshal(data, &geo); err != nil {
		return Geolocation{}, err
	}

	return geo, nil
}
