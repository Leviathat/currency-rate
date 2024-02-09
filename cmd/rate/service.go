package rate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RateService struct{}

func (s *RateService) GetRateData(rateName string) (interface{}, error) {
	// Construct the API URL
	apiURL := fmt.Sprintf("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1&ids=%s", rateName)

	// Make the request to the API
	response, err := http.Get(apiURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON response
	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
