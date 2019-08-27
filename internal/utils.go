package internal

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	logger "github.com/sirupsen/logrus"
)

const (
	fuel = "electricity"
)

var (
	baseURI    = "api.octopus.energy"
	scheme     = "https"
	rate       float64
	rateMetric = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			// set metric name and help string
			Name: "octopus_agile_kwh_price",
			Help: "Gauge vector in pence per kilowatt hour.",
		},
		// set labels for each metric
		[]string{
			"fuel",
		},
	)
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
		logger.Error(err)
		return rate, err
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logger.Error(err)
		return rate, err
	}

	err = json.Unmarshal(data, &responseJSON)
	if err != nil {
		logger.Error(err)
		return rate, err
	}

	now := time.Now()
	for _, period := range responseJSON.Results {
		if period.ValidFrom.Before(now) && period.ValidTo.After(now) {
			rate = period.ValueIncVat
		}
	}

	rateMetric.WithLabelValues(
		fuel,
	).Set(rate)

	return rate, err
}

// Infinite loop to run the application's functionality
func Loop(productCode, tariffCode string) {
	prometheus.MustRegister(rateMetric)

	go ServeMetrics()

	URI := createURL(productCode, tariffCode)
	for {
		go getCurrentRate(URI)

		// sleep until the next iteration
		time.Sleep(60 * time.Second)
	}
}

// Serves all registered metrics for scraping
func ServeMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
