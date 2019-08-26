package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

var (
	baseURI = "api.octopus.energy"
	scheme  = "https"
	rate    float64
)

func createURL(productCode, tariffCode string) string {
	fullURL := url.URL{
		Scheme: scheme,
		Host:   baseURI,
		Path:   "/v1/products/" + productCode + "/electricity-tariffs/" + tariffCode + "/standard-unit-rates/",
	}
	return fullURL.String()
}

func getCurrentRate(URI string) float64 {
	var responseJSON APIResponse

	response, err := http.Get(URI)
	if err != nil {
		fmt.Println("failed")
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("failed")
	}

	err = json.Unmarshal(data, &responseJSON)
	if err != nil {
		fmt.Println("failed")
	}

	now := time.Now()
	for _, period := range responseJSON.Results {
		if period.ValidFrom.Before(now) && period.ValidTo.After(now) {
			rate = period.ValueIncVat
		}
	}
	fmt.Println(rate)
	return rate
}

func CollectorLoop(productCode, tariffCode string) {
	URI := createURL(productCode, tariffCode)
	for {
		go getCurrentRate(URI)

		// sleep until the next iteration
		time.Sleep(60 * time.Second)
	}
}
