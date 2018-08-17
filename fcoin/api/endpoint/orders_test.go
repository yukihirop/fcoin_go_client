package endpoint_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"fcoin_go_client/fcoin/api/endpoint/mock"
)

var _ = Describe("Orders", func() {
	Describe("CreateOrderLimit", func() {
		mockAPI := mock.NewMockAPI("./endpoint/orders/create_order_limit")
		subject, _ := mockAPI.CreateOrderLimit(mock.Symbol("ethusdt"), mock.Side("buy"), mock.Price(1), mock.Amount(0.001))

		It("should return response data", func() {
			Expect(subject).To(Equal(`{"status":0,"data":"9a0rGXh0EyTgAIPHaHSskK_34XrXitsAuk15XeS5T6E="}`))
		})
	})

	Describe("OrderList", func() {
		mockAPI := mock.NewMockAPI("./endpoint/orders/order_list")
		subject, _ := mockAPI.OrderList(mock.Symbol("ethusdt"), mock.States("canceled"))

		It("should return response data", func() {
			Expect(subject).To(Equal(`{"status":0,"data":[{"id":"9a0rGXh0EyTgAIPHaHSskK_34XrXitsAuk15XeS5T6E=","symbol":"ethusdt","amount":"0.001000000000000000","price":"1.000000000000000000","created_at":1534172684408,"type":"limit","side":"buy","filled_amount":"0.000000000000000000","executed_value":"0.000000000000000000","fill_fees":"0.000000000000000000","source":"api","state":"canceled"},{"id":"yFOUtZTtiUBOUX5yBFkDwpU-BK6O2dQaIh5X4oUdStE=","symbol":"ethusdt","amount":"0.001000000000000000","price":"1.000000000000000000","created_at":1534075848013,"type":"limit","side":"buy","filled_amount":"0.000000000000000000","executed_value":"0.000000000000000000","fill_fees":"0.000000000000000000","source":"api","state":"canceled"},{"id":"kW3cRiXIGHG-cHNdterbFBaVRsYfNqfoMzbeHEQcqp4=","symbol":"ethusdt","amount":"0.001000000000000000","price":"1000.000000000000000000","created_at":1531734147607,"type":"limit","side":"sell","filled_amount":"0.000000000000000000","executed_value":"0.000000000000000000","fill_fees":"0.000000000000000000","source":"api","state":"canceled"},{"id":"nMEC_VrW0LYlP4iCcWzmdL50jFrvNWZoaQxvZSjeUSA=","symbol":"ethusdt","amount":"0.001000000000000000","price":"1000.000000000000000000","created_at":1531732778736,"type":"limit","side":"sell","filled_amount":"0.000000000000000000","executed_value":"0.000000000000000000","fill_fees":"0.000000000000000000","source":"api","state":"canceled"},{"id":"V6WE1JJCUwvqJFgCjfT5lA37ioU_xOYi6_NSrtOVOsw=","symbol":"ethusdt","amount":"0.001000000000000000","price":"1000.000000000000000000","created_at":1531732713230,"type":"limit","side":"sell","filled_amount":"0.000000000000000000","executed_value":"0.000000000000000000","fill_fees":"0.000000000000000000","source":"api","state":"canceled"},{"id":"R0moy92q4Qaf_GDEQ6Y1IKCgl5wwJM2bz_Zyacp-Ek8=","symbol":"ethusdt","amount":"0.001000000000000000","price":"1000.000000000000000000","created_at":1531714333569,"type":"limit","side":"sell","filled_amount":"0.000000000000000000","executed_value":"0.000000000000000000","fill_fees":"0.000000000000000000","source":"api","state":"canceled"},{"id":"L7rbALEIoI0ymo3uOXBF4gT4BlyTxgHhGoZjvptIv2U=","symbol":"ethusdt","amount":"0.001000000000000000","price":"1000.000000000000000000","created_at":1531714218130,"type":"limit","side":"sell","filled_amount":"0.000000000000000000","executed_value":"0.000000000000000000","fill_fees":"0.000000000000000000","source":"api","state":"canceled"},{"id":"32ZZCBEpPz2J9oFIJ4RMTIbypltjrVD9PAdYxQTHhUE=","symbol":"ethusdt","amount":"0.001000000000000000","price":"1000.000000000000000000","created_at":1531714092732,"type":"limit","side":"sell","filled_amount":"0.000000000000000000","executed_value":"0.000000000000000000","fill_fees":"0.000000000000000000","source":"api","state":"canceled"},{"id":"VotO2IKI2opgyKRd1lhR5bYj9zNZ398IW85gcBNPisU=","symbol":"ethusdt","amount":"0.001000000000000000","price":"1000.000000000000000000","created_at":1531712709955,"type":"limit","side":"sell","filled_amount":"0.000000000000000000","executed_value":"0.000000000000000000","fill_fees":"0.000000000000000000","source":"api","state":"canceled"},{"id":"tYH6LczJxaVe_WhsLOzOk4YM53hK2q169nYn9ReiwGM=","symbol":"ethusdt","amount":"0.001000000000000000","price":"1000.000000000000000000","created_at":1531675732267,"type":"limit","side":"sell","filled_amount":"0.000000000000000000","executed_value":"0.000000000000000000","fill_fees":"0.000000000000000000","source":"web","state":"canceled"},{"id":"U50WtZkmIh_bbuVKoipAMayCIy0A7qk4hBLxpDvKdPk=","symbol":"ethusdt","amount":"0.025800000000000000","price":"491.100000000000000000","created_at":1529665880201,"type":"limit","side":"buy","filled_amount":"0.000000000000000000","executed_value":"0.000000000000000000","fill_fees":"0.000000000000000000","source":"web","state":"canceled"}]}`))
		})
	})

	Describe("ReferenceOrder", func() {
		mockAPI := mock.NewMockAPI("./endpoint/orders/reference_order")
		subject, _ := mockAPI.ReferenceOrder(mock.OrderId("9a0rGXh0EyTgAIPHaHSskK_34XrXitsAuk15XeS5T6E="))

		It("should return response data", func() {
			Expect(subject).To(Equal(`{"status":0,"data":{"id":"9a0rGXh0EyTgAIPHaHSskK_34XrXitsAuk15XeS5T6E=","symbol":"ethusdt","amount":"0.001000000000000000","price":"1.000000000000000000","created_at":1534172684408,"type":"limit","side":"buy","filled_amount":"0.000000000000000000","executed_value":"0.000000000000000000","fill_fees":"0.000000000000000000","source":"api","state":"submitted"}}`))
		})
	})

	Describe("CancelOrder", func() {
		mockAPI := mock.NewMockAPI("./endpoint/orders/cancel_order")
		subject, _ := mockAPI.CancelOrder(mock.OrderId("9a0rGXh0EyTgAIPHaHSskK_34XrXitsAuk15XeS5T6E="))

		It("should return response data", func() {
			Expect(subject).To(Equal(`{"status":0}`))
		})
	})

	Describe("OrderMatchResults", func() {
		mockAPI := mock.NewMockAPI("./endpoint/orders/order_match_results")
		subject, _ := mockAPI.OrderMatchResults(mock.OrderId("9a0rGXh0EyTgAIPHaHSskK_34XrXitsAuk15XeS5T6E="))

		It("should return response data", func() {
			Expect(subject).To(Equal(`{"status":0,"data":[]}`))
		})
	})
})
