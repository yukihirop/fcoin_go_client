package orders_test

import (
	"errors"
	"fcoin_go_client/fcoin/api/endpoint/validator/orders"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CreateOrderMarketValidator", func() {
	SetValidationSetting()

	Describe("IsValid", func() {
		Context("when valid", func() {
			validator := orders.NewCreateOrderMarketValidator(orders.Symbol("ethusdt"), orders.Side("buy"), orders.Total(100), orders.VSetting(fixedViper, customViper, nil))
			subject := validator.IsValid()

			It("should be true", func() {
				Expect(subject).To(Equal(true))
			})
		})

		Context("when invalid", func() {
			validator := orders.NewCreateOrderMarketValidator(orders.Symbol("fiusdt"), orders.Side("sell"), orders.Amount(0), orders.VSetting(fixedViper, customViper, nil))
			subject := validator.IsValid()

			It("should be false", func() {
				Expect(subject).To(Equal(false))
			})
		})
	})

	Describe("Messages", func() {
		Context("when valid", func() {
			validator := orders.NewCreateOrderMarketValidator(orders.Symbol("ethusdt"), orders.Side("buy"), orders.Total(100), orders.VSetting(fixedViper, customViper, nil))
			subject := validator.Messages()

			It("should be blank", func() {
				Expect(subject).To(Equal([]error{}))
			})
		})

		Context("when invalid", func() {
			Context("when symbol board is adapted on-going order", func() {
				validator := orders.NewCreateOrderMarketValidator(orders.Symbol("ethusdt"), orders.Side("sell"), orders.Amount(0), orders.VSetting(fixedViper, customViper, nil))
				subject := validator.Messages()

				It("shold return error mesasges", func() {
					Expect(subject).To(Equal([]error{
						errors.New("{amount: amount is 0. amount is not betweeen 0.001 and 10000.}"),
					}))
				})
			})

			Context("when symbol board is not adapted on-going order", func() {
				validator := orders.NewCreateOrderMarketValidator(orders.Symbol("fiusdt"), orders.Side("sell"), orders.Amount(0), orders.VSetting(fixedViper, customViper, nil))
				subject := validator.Messages()

				It("shold return error mesasges", func() {
					Expect(subject).To(Equal([]error{
						errors.New("{symbol: symbol is fiusdt. symbol board is not adapted on-going order.}"),
					}))
				})
			})
		})
	})
})
