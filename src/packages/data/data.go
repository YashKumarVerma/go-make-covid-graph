package data

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type responseStructure struct {
	country     string
	countryCode string
	province    string
	city        string
	cityCode    string
	lat         string
	lon         string
	Cases       int
	status      int
	Date        string
}

// Fetch latest data from COVID api
func Fetch() []float64 {
	var cases []float64
	resp, err := http.Get("https://api.covid19api.com/country/india/status/confirmed")

	if err != nil {
		return fmt.Errorf("Unable to connect to server")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Invalid body received")
	}

	var response []responseStructure
	json.Unmarshal([]byte(body), &response)

	for _, item := range response {
		cases = append(cases, float64(item.Cases))
	}

	return cases
}
