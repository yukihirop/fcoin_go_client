package validator_test

import (
	"fcoin_go_client/fcoin/api/endpoint/validator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OrdersValidator", func() {
	SetValidationSetting()

	Describe("IsValid", func() {
		Context("when CreateOrderLimit", func() {
			v := validator.NewOrdersValidator(validator.MethodName("CreateOrderLimit"), validator.Symbol("ethusdt"), validator.Type("limit"), validator.Side("buy"), validator.VSetting(fixedViper, customViper, nil))
			subject := v.IsValid()

			It("should be false", func() {
				Expect(subject).To(Equal(false))
			})
		})

		Context("when CreateOrderMarket", func() {
			v := validator.NewOrdersValidator(validator.MethodName("CreateOrderMarket"), validator.Symbol("fiusdt"), validator.Type("market"), validator.Side("buy"), validator.VSetting(fixedViper, customViper, nil))
			subject := v.IsValid()

			It("should be false", func() {
				Expect(subject).To(Equal(false))
			})
		})

		Context("when OrderList", func() {
			v := validator.NewOrdersValidator(validator.MethodName("OrderList"), validator.Symbol("ethusdt"), validator.States("invalid_states"), validator.VSetting(fixedViper, customViper, nil))
			subject := v.IsValid()

			It("should be false", func() {
				Expect(subject).To(Equal(false))
			})
		})
	})

	Describe("Messages", func() {
		Context("when CreateOrderLimit", func() {
			v := validator.NewOrdersValidator(validator.MethodName("CreateOrderLimit"), validator.Symbol("ethusdt"), validator.Type("limit"), validator.Side("buy"), validator.VSetting(fixedViper, customViper, nil))
			subject := v.Messages()

			It("should return error messages", func() {
				Expect(subject).To(Equal(map[string]string{
					"price":  "price is 0. price is not betweeen 1 and 10000.",
					"amount": "amount is 0. amount is not betweeen 0.001 and 10000.",
				}))
			})
		})

		Context("when CreateOrderMarket", func() {
			v := validator.NewOrdersValidator(validator.MethodName("CreateOrderMarket"), validator.Symbol("fiusdt"), validator.Type("market"), validator.Side("buy"), validator.VSetting(fixedViper, customViper, nil))
			subject := v.Messages()

			It("should return error messages", func() {
				Expect(subject).To(Equal(map[string]string{
					"symbol": "symbol is fiusdt. symbol board is not adapted on-going order.",
				}))
			})
		})

		Context("when OrderList", func() {
			v := validator.NewOrdersValidator(validator.MethodName("OrderList"), validator.Symbol("ethusdt"), validator.States("invalid_states"), validator.VSetting(fixedViper, customViper, nil))
			subject := v.Messages()

			It("should be false", func() {
				Expect(subject).To(Equal(map[string]string{
					"states": "states is invalid_states. states is not included in the [submitted partial_filled canceled partial_canceled filled pending_cancel].",
				}))
			})
		})
	})
})
