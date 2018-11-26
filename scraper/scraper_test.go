package scraper_test

import (
	"github.com/hiko1129/macrotrends-scraper/scraper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scraper", func() {

	var client *scraper.Scraper
	BeforeEach(func() {
		client, _ = scraper.New()
	})

	Describe("#GetPERHistoricalData", func() {
		It("expects no error", func() {
			data, err := client.GetPERHistoricalData("FB", "facebook")
			Expect(data).NotTo(BeEmpty())
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("#GetPSRHistoricalData", func() {
		It("expects no error", func() {
			_, err := client.GetPSRHistoricalData("FB", "facebook")
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("#GetPBRHistoricalData", func() {
		It("expects no error", func() {
			_, err := client.GetPBRHistoricalData("FB", "facebook")
			Expect(err).NotTo(HaveOccurred())
		})
	})

	Describe("#GetPriceFCFRatioHistoricalData", func() {
		It("expects no error", func() {
			_, err := client.GetPFCFRHistoricalData("FB", "facebook")
			Expect(err).NotTo(HaveOccurred())
		})
	})
})
