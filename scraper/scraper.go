package scraper

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type historicalDatum struct {
	date       *time.Time
	stockPrice float32
	perShare   float32
	ratio      float32
}

// PERHistoricalDatum struct
type PERHistoricalDatum struct {
	Date       *time.Time
	StockPrice float32
	TTMNetEPS  float32
	PER        float32
}

// PSRHistoricalDatum struct
type PSRHistoricalDatum struct {
	Date             *time.Time
	StockPrice       float32
	TTMSalesPerShare float32
	PSR              float32
}

// PBRHistoricalDatum struct
type PBRHistoricalDatum struct {
	Date              *time.Time
	StockPrice        float32
	BookValuePerShare float32
	PBR               float32
}

// PFCFRHistoricalDatum struct
type PFCFRHistoricalDatum struct {
	Date           *time.Time
	StockPrice     float32
	TTMFCFPerShare float32
	PFCFR          float32
}

// FetchPERHistoricalData func
func FetchPERHistoricalData(tickerSymbol string, companyName string) ([]PERHistoricalDatum, error) {
	url := fmt.Sprintf("https://www.macrotrends.net/stocks/charts/%v/%v/pe-ratio", tickerSymbol, companyName)
	data, err := fetchHistoricalData(url)
	result := []PERHistoricalDatum{}
	if err != nil {
		return result, err
	}

	for _, datum := range data {
		result = append(result, PERHistoricalDatum{
			Date:       datum.date,
			StockPrice: datum.stockPrice,
			TTMNetEPS:  datum.perShare,
			PER:        datum.ratio,
		})
	}

	return result, nil
}

// FetchPSRHistoricalData func
func FetchPSRHistoricalData(tickerSymbol string, companyName string) ([]PSRHistoricalDatum, error) {
	url := fmt.Sprintf("https://www.macrotrends.net/stocks/charts/%v/%v/price-sales", tickerSymbol, companyName)
	data, err := fetchHistoricalData(url)
	result := []PSRHistoricalDatum{}
	if err != nil {
		return result, err
	}

	for _, datum := range data {
		result = append(result, PSRHistoricalDatum{
			Date:             datum.date,
			StockPrice:       datum.stockPrice,
			TTMSalesPerShare: datum.perShare,
			PSR:              datum.ratio,
		})
	}

	return result, nil
}

// FetchPBRHistoricalData func
func FetchPBRHistoricalData(tickerSymbol string, companyName string) ([]PBRHistoricalDatum, error) {
	url := fmt.Sprintf("https://www.macrotrends.net/stocks/charts/%v/%v/price-book", tickerSymbol, companyName)
	data, err := fetchHistoricalData(url)
	result := []PBRHistoricalDatum{}
	if err != nil {
		return result, err
	}

	for _, datum := range data {
		result = append(result, PBRHistoricalDatum{
			Date:              datum.date,
			StockPrice:        datum.stockPrice,
			BookValuePerShare: datum.perShare,
			PBR:               datum.ratio,
		})
	}

	return result, nil
}

// FetchPFCFRHistoricalData func
func FetchPFCFRHistoricalData(tickerSymbol string, companyName string) ([]PFCFRHistoricalDatum, error) {
	url := fmt.Sprintf("https://www.macrotrends.net/stocks/charts/%v/%v/price-fcf", tickerSymbol, companyName)
	data, err := fetchHistoricalData(url)
	result := []PFCFRHistoricalDatum{}
	if err != nil {
		return result, err
	}

	for _, datum := range data {
		result = append(result, PFCFRHistoricalDatum{
			Date:           datum.date,
			StockPrice:     datum.stockPrice,
			TTMFCFPerShare: datum.perShare,
			PFCFR:          datum.ratio,
		})
	}

	return result, nil
}

func fetchHistoricalData(url string) ([]historicalDatum, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	data := []historicalDatum{}
	var errs []error
	doc.Find("#style-1 tbody").Children().Each(func(i int, s *goquery.Selection) {
		date, err := time.Parse("2006-01-02", s.Find("td:nth-child(1)").Text())
		if err != nil {
			errs = append(errs, err)
		}

		stockPrice, err := strconv.ParseFloat(s.Find("td:nth-child(2)").Text(), 32)
		if err != nil {
			errs = append(errs, err)
		}

		ttmNetEPSString := s.Find("td:nth-child(3)").Text()
		if len(ttmNetEPSString) > 0 {
			ttmNetEPSString = ttmNetEPSString[1:]
		} else {
			ttmNetEPSString = "0"
		}

		ttmNetEPS, err := strconv.ParseFloat(ttmNetEPSString, 32)
		if err != nil {
			errs = append(errs, err)
		}

		per, err := strconv.ParseFloat(s.Find("td:nth-child(4)").Text(), 32)
		if err != nil {
			errs = append(errs, err)
		}

		data = append(data, historicalDatum{
			date:       &date,
			stockPrice: float32(stockPrice),
			perShare:   float32(ttmNetEPS),
			ratio:      float32(per),
		})
	})
	if len(errs) > 0 {
		return nil, errors.New("scraping failed")
	}

	return data, nil
}
