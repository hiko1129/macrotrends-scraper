package scraper_test

import (
	"time"

	"github.com/hiko1129/macrotrends-scraper/scraper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scraper", func() {
	Describe("#FetchPERHistoricalData", func() {
		It("expects no error", func() {
			data, err := scraper.FetchPERHistoricalData("FB", "facebook")
			Expect(data).NotTo(BeEmpty())
			Expect(err).NotTo(HaveOccurred())
			time.Sleep(1 * time.Second)
		})
	})

	Describe("#FetchPSRHistoricalData", func() {
		It("expects no error", func() {
			data, err := scraper.FetchPSRHistoricalData("FB", "facebook")
			Expect(data).NotTo(BeEmpty())
			Expect(err).NotTo(HaveOccurred())
			time.Sleep(1 * time.Second)
		})
	})

	Describe("#FetchPBRHistoricalData", func() {
		It("expects no error", func() {
			data, err := scraper.FetchPBRHistoricalData("FB", "facebook")
			Expect(data).NotTo(BeEmpty())
			Expect(err).NotTo(HaveOccurred())
			time.Sleep(1 * time.Second)
		})
	})

	Describe("#FetchPriceFCFRatioHistoricalData", func() {
		It("expects no error", func() {
			data, err := scraper.FetchPFCFRHistoricalData("FB", "facebook")
			Expect(data).NotTo(BeEmpty())
			Expect(err).NotTo(HaveOccurred())
			time.Sleep(1 * time.Second)
		})
	})
})
