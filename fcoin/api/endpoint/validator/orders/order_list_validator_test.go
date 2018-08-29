package orders_test

import (
	"errors"
	"fcoin_go_client/fcoin/api/endpoint/validator/orders"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("OrderListValidator", func() {
	SetValidationSetting()

	Describe("IsValid", func() {
		Context("when valid", func() {
			validator := orders.NewOrderListValidator(orders.Symbol("ethusdt"), orders.States("canceled"), orders.VSetting(fixedViper, customViper, nil))
			subject := validator.IsValid()
			It("should be true", func() {
				Expect(subject).To(Equal(true))
			})
		})

		Context("when invalid", func() {
			validator := orders.NewOrderListValidator(orders.VSetting(fixedViper, customViper, nil))
			subject := validator.IsValid()
			It("should be false", func() {
				Expect(subject).To(Equal(false))
			})
		})
	})

	Describe("Messages", func() {
		Context("when valid", func() {
			validator := orders.NewOrderListValidator(orders.Symbol("ethusdt"), orders.States("canceled"), orders.VSetting(fixedViper, customViper, nil))
			subject := validator.Messages()
			It("should be blank", func() {
				Expect(subject).To(Equal([]error{}))
			})
		})

		Context("when invalid", func() {
			Context("when symbol is blank", func() {
				validator := orders.NewOrderListValidator(orders.States("canceled"), orders.VSetting(fixedViper, customViper, nil))
				subject := validator.Messages()
				It("should return error message", func() {
					Expect(subject).To(Equal([]error{
						errors.New("{symbol: symbol is . symbol can't be blank.}"),
					}))
				})
			})

			Context("when states is invalid", func() {
				validator := orders.NewOrderListValidator(orders.Symbol("ethusdt"), orders.States("invalid_states"), orders.VSetting(fixedViper, customViper, nil))
				subject := validator.Messages()
				It("should return error message", func() {
					Expect(subject).To(Equal([]error{
						errors.New("{states: states is invalid_states. states is not included in the [submitted partial_filled canceled partial_canceled filled pending_cancel].}"),
					}))
				})
			})
		})
	})
})
