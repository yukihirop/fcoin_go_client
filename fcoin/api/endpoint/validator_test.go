package endpoint_test

import (
	"errors"
	"fcoin_go_client/fcoin/api/endpoint"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Validator", func() {
	SetValidationSetting()

	Describe("IsValid", func() {
		Describe("when MarketDepth", func() {
			validator := endpoint.NewValidator(endpoint.MethodName("MarketDepth"), endpoint.VSetting(fixedViper, customViper, nil))
			subject := validator.IsValid()

			It("should be false", func() {
				Expect(subject).To(Equal(false))
			})
		})

		Describe("when MarketCandles", func() {
			validator := endpoint.NewValidator(endpoint.MethodName("MarketCandles"), endpoint.VSetting(fixedViper, customViper, nil))
			subject := validator.IsValid()

			It("should be false", func() {
				Expect(subject).To(Equal(false))
			})
		})

		Describe("when CreateOrderLimit", func() {
			validator := endpoint.NewValidator(endpoint.MethodName("CreateOrderLimit"), endpoint.VSetting(fixedViper, customViper, nil))
			subject := validator.IsValid()

			It("should be false", func() {
				Expect(subject).To(Equal(false))
			})
		})

		Describe("when CreateOrderMarket", func() {
			validator := endpoint.NewValidator(endpoint.MethodName("CreateOrderMarket"), endpoint.VSetting(fixedViper, customViper, nil))
			subject := validator.IsValid()

			It("should be false", func() {
				Expect(subject).To(Equal(false))
			})
		})

		Describe("when OrderList", func() {
			validator := endpoint.NewValidator(endpoint.MethodName("OrderList"), endpoint.VSetting(fixedViper, customViper, nil))
			subject := validator.IsValid()

			It("should be false", func() {
				Expect(subject).To(Equal(false))
			})
		})
	})

	Describe("Messages", func() {
		Describe("when MarketDepth", func() {
			validator := endpoint.NewValidator(endpoint.MethodName("MarketDepth"), endpoint.VSetting(fixedViper, customViper, nil))
			subject := validator.Messages()

			It("should return error message", func() {
				Expect(subject).To(Equal([]error{
					errors.New("{symbol: symbol is . symbol can't be blank.}"),
					errors.New("{level: level is . level is not included in the [L20 L100 full].}"),
				}))
			})
		})

		Describe("when MarketCandles", func() {
			validator := endpoint.NewValidator(endpoint.MethodName("MarketCandles"), endpoint.VSetting(fixedViper, customViper, nil))
			subject := validator.Messages()

			It("should return error message", func() {
				Expect(subject).To(Equal([]error{
					errors.New("{symbol: symbol is . symbol can't be blank.}"),
					errors.New("{resolution: resolution is . resolution is not included in the [M1 M3 M5 M15 M30 H1 H4 H6 D1 W1 MN].}"),
				}))
			})
		})

		Describe("when CreateOrderLimit", func() {
			validator := endpoint.NewValidator(endpoint.MethodName("CreateOrderLimit"), endpoint.VSetting(fixedViper, customViper, nil))
			subject := validator.Messages()

			It("should return error message", func() {
				Expect(subject).To(Equal([]error{
					errors.New("{symbol: symbol is . symbol can't be blank.}"),
				}))
			})
		})

		Describe("when CreateOrderMarket", func() {
			validator := endpoint.NewValidator(endpoint.MethodName("CreateOrderMarket"), endpoint.VSetting(fixedViper, customViper, nil))
			subject := validator.Messages()

			It("should be false", func() {
				Expect(subject).To(Equal([]error{
					errors.New("{symbol: symbol is . symbol can't be blank.}"),
				}))
			})
		})

		Describe("when OrderList", func() {
			validator := endpoint.NewValidator(endpoint.MethodName("OrderList"), endpoint.VSetting(fixedViper, customViper, nil))
			subject := validator.Messages()

			It("should be false", func() {
				Expect(subject).To(Equal([]error{
					errors.New("{symbol: symbol is . symbol can't be blank.}"),
				}))
			})
		})
	})
})
