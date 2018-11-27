package main

import (
	"fmt"
	"log"

	"github.com/hiko1129/macrotrends-scraper/scraper"
)

func main() {
	// https://www.macrotrends.net/stocks/charts/FB/facebook/price-fcf
	// FB facebook
	pbrData, err := scraper.FetchPBRHistoricalData("FB", "facebook")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PBRData")
	fmt.Println(pbrData)
	fmt.Println()

	perData, err := scraper.FetchPERHistoricalData("FB", "facebook")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PERData")
	fmt.Println(perData)
	fmt.Println()

	psrData, err := scraper.FetchPSRHistoricalData("FB", "facebook")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PSRData")
	fmt.Println(psrData)
	fmt.Println()

	pfcfrData, err := scraper.FetchPFCFRHistoricalData("FB", "facebook")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PFCFRData")
	fmt.Println(pfcfrData)
	fmt.Println()
}
