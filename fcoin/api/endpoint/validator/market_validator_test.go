package validator_test

import (
	"fcoin_go_client/fcoin/api/endpoint/validator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MarketValidator", func() {
	SetValidationSetting()

	Describe("IsValid", func() {
		Context("when MarketDepth", func() {
			Context("when valid", func() {
				v := validator.NewMarketValidator(validator.MethodName("MarketDepth"), validator.Symbol("ethusdt"), validator.Level("L20"), validator.VSetting(fixedViper, customViper, nil))
				subject := v.IsValid()

				It("should be true", func() {
					Expect(subject).To(Equal(true))
				})
			})

			Context("when invalid", func() {
				v := validator.NewMarketValidator(validator.MethodName("MarketDepth"), validator.VSetting(fixedViper, customViper, nil))
				subject := v.IsValid()

				It("should be false", func() {
					Expect(subject).To(Equal(false))
				})
			})
		})

		Context("when MarketCandles", func() {
			Context("when valid", func() {
				v := validator.NewMarketValidator(validator.MethodName("MarketCandles"), validator.Symbol("ethusdt"), validator.Resolution("MN"), validator.VSetting(fixedViper, customViper, nil))
				subject := v.IsValid()

				It("should be true", func() {
					Expect(subject).To(Equal(true))
				})
			})

			Context("when invalid", func() {
				v := validator.NewMarketValidator(validator.MethodName("MarketCandles"), validator.VSetting(fixedViper, customViper, nil))
				subject := v.IsValid()

				It("should be false", func() {
					Expect(subject).To(Equal(false))
				})
			})
		})
	})

	Describe("Messages", func() {
		Context("when MarketDepth", func() {
			Context("when valid", func() {
				v := validator.NewMarketValidator(validator.MethodName("MarketDepth"), validator.Symbol("ethusdt"), validator.Level("L20"), validator.VSetting(fixedViper, customViper, nil))
				subject := v.Messages()

				It("should be blank", func() {
					Expect(subject).To(Equal(map[string]string{}))
				})
			})

			Context("when invalid", func() {
				v := validator.NewMarketValidator(validator.MethodName("MarketDepth"), validator.VSetting(fixedViper, customViper, nil))
				subject := v.Messages()

				It("should return error message", func() {
					Expect(subject).To(Equal(map[string]string{
						"symbol": "symbol is . symbol can't be blank.",
						"level":  "level is . level is not included in the [L20 L100 full].",
					}))
				})
			})
		})

		Context("when MarketCandles", func() {
			Context("when valid", func() {
				v := validator.NewMarketValidator(validator.MethodName("MarketCandles"), validator.Symbol("ethusdt"), validator.Resolution("MN"), validator.VSetting(fixedViper, customViper, nil))
				subject := v.Messages()

				It("should be blank", func() {
					Expect(subject).To(Equal(map[string]string{}))
				})
			})

			Context("when invalid", func() {
				v := validator.NewMarketValidator(validator.MethodName("MarketCandles"), validator.VSetting(fixedViper, customViper, nil))
				subject := v.Messages()

				It("should return error message", func() {
					Expect(subject).To(Equal(map[string]string{
						"symbol":     "symbol is . symbol can't be blank.",
						"resolution": "resolution is . resolution is not included in the [M1 M3 M5 M15 M30 H1 H4 H6 D1 W1 MN].",
					}))
				})
			})
		})
	})
})
