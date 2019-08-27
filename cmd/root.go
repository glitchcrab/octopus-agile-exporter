/*
Copyright © 2019 Simon Weald

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/glitchcrab/octopus-agile-exporter/internal"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var productCode string
var tariffCode string

var rootCmd = &cobra.Command{
	Use:   "octopus-agile-exporter",
	Short: "Octopus Energy Prometheus exporter",
	Long: `Prometheus exporter for Octopus Energy's Agile tarrif.

Daily half-hourly unit prices are published the preceeding day, but this
exporter will only provide the unit price for the current time period.

Unit prices are in pence per kilowatt hour.

Product and tariff codes can be provided as flags or environment variables.
If provided, environment variables take precedence.`,

	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Exporter initialised.")
		internal.Loop(productCode, tariffCode)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Initialise config and parse flags.
func init() {
	logger.Info("Starting exporter.")
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&productCode, "product-code", "", "Product code (env OCTOPUS_PRODUCT_CODE).")
	rootCmd.PersistentFlags().StringVar(&tariffCode, "tarrif-code", "", "Tariff code (env OCTOPUS_TARIFF_CODE).")
}

// initConfig reads ENV variables if set. ENV variables take precedence over flags.
func initConfig() {
	// get product code if it isn't passed as a flag
	if os.Getenv("OCTOPUS_PRODUCT_CODE") != "" {
		productCode = os.Getenv("OCTOPUS_PRODUCT_CODE")
	} else {
		if productCode == "" {
			logger.Fatal("Product code must be provided.")
		}
	}
	// get tariff code if it isn't passed as a flag
	if os.Getenv("OCTOPUS_TARIFF_CODE") != "" {
		tariffCode = os.Getenv("OCTOPUS_TARIFF_CODE")
	} else {
		if tariffCode == "" {
			logger.Fatal("Tariff code must be provided.")
		}
	}
}
