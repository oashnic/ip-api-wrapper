package ipapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// DNS response struct
type DNS struct {
	DNS struct {
		IP  string `json:"ip"`
		Geo string `json:"geo"`
	} `json:"dns"`
	Edns struct {
		IP  string `json:"ip"`
		Geo string `json:"geo"`
	} `json:"edns"`
}

// GetDNS get current dns servers
func GetDNS() (DNS, error) {
	resp, err := http.DefaultClient.Get("http://edns.ip-api.com/json")
	if err != nil {
		return DNS{}, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return DNS{}, err
	}

	var dns DNS

	if err = json.Unmarshal(data, &dns); err != nil {
		return dns, err
	}

	return dns, nil
}
