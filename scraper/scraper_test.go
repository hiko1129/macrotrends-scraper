package scraper_test

import (
	"time"

	"github.com/hiko1129/macrotrends-scraper/scraper"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Scraper", func() {
	Describe("#FetchPERHistoricalData", func() {
		Context("when valid ticker symbol", func() {
			It("expects no error", func() {
				data, err := scraper.FetchPERHistoricalData("FB", "facebook")
				Expect(data).NotTo(BeEmpty())
				Expect(err).NotTo(HaveOccurred())
				time.Sleep(1 * time.Second)
			})
		})

		Context("when invalid ticker symbol", func() {
			It("expects error", func() {
				data, err := scraper.FetchPERHistoricalData("hogefuga", "hogefuga")
				Expect(data).To(BeEmpty())
				Expect(err).To(HaveOccurred())
				time.Sleep(1 * time.Second)
			})
		})
	})

	// Describe("#FetchPBRHistoricalData", func() {
	// 	Context("when valid ticker symbol", func() {
	// 		It("expects no error", func() {
	// 			data, err := scraper.FetchPBRHistoricalData("FB", "facebook")
	// 			Expect(data).NotTo(BeEmpty())
	// 			Expect(err).NotTo(HaveOccurred())
	// 			time.Sleep(1 * time.Second)
	// 		})
	// 	})

	// 	Context("when invalid ticker symbol", func() {
	// 		It("expects error", func() {
	// 			data, err := scraper.FetchPBRHistoricalData("hogefuga", "hogefuga")
	// 			Expect(data).To(BeEmpty())
	// 			Expect(err).To(HaveOccurred())
	// 			time.Sleep(1 * time.Second)
	// 		})
	// 	})
	// })

	// Describe("#FetchPSRHistoricalData", func() {
	// 	Context("when valid ticker symbol", func() {
	// 		It("expects no error", func() {
	// 			data, err := scraper.FetchPSRHistoricalData("FB", "facebook")
	// 			Expect(data).NotTo(BeEmpty())
	// 			Expect(err).NotTo(HaveOccurred())
	// 			time.Sleep(1 * time.Second)
	// 		})
	// 	})

	// 	Context("when invalid ticker symbol", func() {
	// 		It("expects error", func() {
	// 			data, err := scraper.FetchPSRHistoricalData("hogefuga", "hogefuga")
	// 			Expect(data).To(BeEmpty())
	// 			Expect(err).To(HaveOccurred())
	// 			time.Sleep(1 * time.Second)
	// 		})
	// 	})
	// })

	// Describe("#FetchPFCFRHistoricalData", func() {
	// 	Context("when valid ticker symbol", func() {
	// 		It("expects no error", func() {
	// 			data, err := scraper.FetchPFCFRHistoricalData("FB", "facebook")
	// 			Expect(data).NotTo(BeEmpty())
	// 			Expect(err).NotTo(HaveOccurred())
	// 			time.Sleep(1 * time.Second)
	// 		})
	// 	})

	// 	Context("when invalid ticker symbol", func() {
	// 		It("expects error", func() {
	// 			data, err := scraper.FetchPFCFRHistoricalData("hogefuga", "hogefuga")
	// 			Expect(data).To(BeEmpty())
	// 			Expect(err).To(HaveOccurred())
	// 			time.Sleep(1 * time.Second)
	// 		})
	// 	})
	// })
})
