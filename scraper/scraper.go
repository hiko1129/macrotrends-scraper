package scraper

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// Scraper struct
type Scraper struct{}

// New func
func New() (*Scraper, error) {
	return &Scraper{}, nil
}

type datum interface {
	p()
}

// PERHistoricalDatum struct
type PERHistoricalDatum struct {
	Date       *time.Time
	StockPrice float32
	TTMNetEPS  float32
	PER        float32
}

func (p *PERHistoricalDatum) p() {}

// PSRHistoricalDatum struct
type PSRHistoricalDatum struct {
	Date             *time.Time
	StockPrice       float32
	TTMSalesPerShare float32
	PSR              float32
}

func (p *PSRHistoricalDatum) p() {}

// PBRHistoricalDatum struct
type PBRHistoricalDatum struct {
	Date              *time.Time
	StockPrice        float32
	BookValuePerShare float32
	PBR               float32
}

func (p *PBRHistoricalDatum) p() {}

// PFCFRHistoricalDatum struct
type PFCFRHistoricalDatum struct {
	Date           *time.Time
	StockPrice     float32
	TTMFCFPerShare float32
	PFCFR          float32
}

func (p *PFCFRHistoricalDatum) p() {}

// GetPERHistoricalData func
func (s *Scraper) GetPERHistoricalData(tickerSymbol string, companyName string) ([]PERHistoricalDatum, error) {
	url := fmt.Sprintf("https://www.macrotrends.net/stocks/charts/%v/%v/pe-ratio", tickerSymbol, companyName)
	return s.getHistoricalData(url)
}

// GetPSRHistoricalData func
func (s *Scraper) GetPSRHistoricalData(tickerSymbol string, companyName string) ([]PSRHistoricalDatum, error) {
	return []PSRHistoricalDatum{}, nil
}

// GetPBRHistoricalData func
func (s *Scraper) GetPBRHistoricalData(tickerSymbol string, companyName string) ([]PBRHistoricalDatum, error) {
	return []PBRHistoricalDatum{}, nil
}

// GetPFCFRHistoricalData func
func (s *Scraper) GetPFCFRHistoricalData(tickerSymbol string, companyName string) ([]PFCFRHistoricalDatum, error) {
	return []PFCFRHistoricalDatum{}, nil
}

func (s *Scraper) getHistoricalData(url string) ([]PERHistoricalDatum, error) {
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

	data := []PERHistoricalDatum{}
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

		data = append(data, PERHistoricalDatum{
			Date:       &date,
			StockPrice: float32(stockPrice),
			TTMNetEPS:  float32(ttmNetEPS),
			PER:        float32(per),
		})
	})
	if len(errs) > 0 {
		return nil, errors.New("scraping failed")
	}

	return data, nil

}
