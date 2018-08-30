package orders_test

import (
	"errors"
	"fcoin_go_client/fcoin/api/endpoint/validator/orders"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("CreateOrderLimitValidator", func() {
	SetValidationSetting()

	Describe("IsValid", func() {
		Context("when valid", func() {
			validator := orders.NewCreateOrderLimitValidator(orders.Symbol("ethusdt"), orders.Side("sell"), orders.Amount(100), orders.Price(100), orders.VSetting(fixedViper, customViper, nil))
			subject := validator.IsValid()

			It("should be true", func() {
				Expect(subject).To(Equal(true))
			})
		})

		Context("When invalid", func() {
			validator := orders.NewCreateOrderLimitValidator(orders.Symbol("ethusdt"), orders.Side("sell"), orders.Amount(0), orders.Price(0), orders.VSetting(fixedViper, customViper, nil))
			subject := validator.IsValid()

			It("should be false", func() {
				Expect(subject).To(Equal(false))
			})
		})
	})

	Describe("Messages", func() {
		Context("when valid", func() {
			validator := orders.NewCreateOrderLimitValidator(orders.Symbol("ethusdt"), orders.Side("sell"), orders.Amount(100), orders.Price(100), orders.VSetting(fixedViper, customViper, nil))
			subject := validator.Messages()

			It("should be blank", func() {
				Expect(subject).To(Equal([]error{}))
			})
		})

		Context("when invalid", func() {
			validator := orders.NewCreateOrderLimitValidator(orders.Symbol("ethusdt"), orders.Side("sell"), orders.Amount(0), orders.Price(0), orders.VSetting(fixedViper, customViper, nil))
			subject := validator.Messages()

			It("should return error message", func() {
				Expect(subject).To(Equal([]error{
					errors.New("{price: price is 0. price is not betweeen 1 and 10000.}"),
					errors.New("{amount: amount is 0. amount is not betweeen 0.001 and 10000.}"),
				}))
			})
		})
	})
})
