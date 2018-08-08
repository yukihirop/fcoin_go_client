package endpoint_test

import (
	"fcoin_go_client/fcoin/api/endpoint/mock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Market", func() {
	Describe("MarketTicker", func() {
		Context("When symbol given", func() {
			mockAPI := mock.NewMockAPI("./endpoint/market/ticker")
			subject, _ := mockAPI.MarketTicker(mock.Symbol("ethusdt"))
			It("should return response data", func() {
				Expect(subject).To(Equal(`{"status":0,"data":{"ticker":[408.090000000,0.030000000,407.970000000,0.030000000,408.100000000,0.005000000,405.300000000,410.980000000,400.600000000,9468.580550920,3860139.362681636880000000],"type":"ticker.ethusdt","seq":107843190}}`))
			})
		})

		Context("when symbol do not support", func() {
			mockAPI := mock.NewMockAPI("./endpoint/market/ticker")
			subject, _ := mockAPI.MarketTicker(mock.Symbol("do_not_support"))
			It("should return error message", func() {
				Expect(subject).To(Equal(`{"status":40003,"msg":"invalid symbol"}`))
			})
		})
	})

	Describe("MarketDepth", func() {
		Context("when symbol and level given", func() {
			mockAPI := mock.NewMockAPI("./endpoint/market/depth")
			subject, _ := mockAPI.MarketDepth(mock.Symbol("ethusdt"), mock.Level("L20"))
			It("housld return response data", func() {
				Expect(subject).To(Equal(`{"status":0,"data":{"bids":[408.750000000,0.001400000,408.710000000,0.380000000,408.700000000,0.005000000,408.620000000,0.012600000,408.500000000,0.005000000,408.300000000,5.005000000,408.160000000,2.689000000,408.100000000,0.005000000,408.060000000,17.900000000,408.000000000,0.100000000,407.940000000,0.100000000,407.900000000,0.105000000,407.870000000,0.100000000,407.860000000,0.100000000,407.840000000,20.805000000,407.820000000,3.100000000,407.810000000,0.200000000,407.800000000,0.170400000,407.790000000,0.100000000,407.780000000,0.100000000],"asks":[408.800000000,0.167000000,408.880000000,0.030000000,408.980000000,0.104900000,409.000000000,1.859700000,409.010000000,10.671200000,409.030000000,0.095000000,409.040000000,0.108400000,409.070000000,0.067500000,409.090000000,0.160000000,409.100000000,0.105000000,409.150000000,0.027800000,409.200000000,0.100000000,409.250000000,2.320000000,409.270000000,1.741400000,409.280000000,0.655600000,409.290000000,0.080000000,409.300000000,0.208700000,409.310000000,0.065900000,409.320000000,0.001500000,409.330000000,0.001900000],"ts":1533654896017,"seq":107857355,"type":"depth.L20.ethusdt"}}`))
			})
		})

		Context("when level is invalid", func() {
			mockAPI := mock.NewMockAPI("./endpoint/market/depth")
			subject, _ := mockAPI.MarketDepth(mock.Symbol("ethusdt"), mock.Level("invalid_level"))
			It("should return error message", func() {
				Expect(subject).To(Equal(`{"status":40003,"msg":"invalid depth-level"}`))
			})
		})
	})

	Describe("MarketTrades", func() {
		Context("when symbol given", func() {
			mockAPI := mock.NewMockAPI("./endpoint/market/trades")
			subject, _ := mockAPI.MarketTrades(mock.Symbol("ethusdt"))
			It("should return response data", func() {
				Expect(subject).To(Equal(`{"status":0,"data":[{"amount":0.500000000,"ts":1533655458586,"id":107863107000,"side":"buy","price":408.460000000}]}`))
			})
		})

		Context("when symbol do not support", func() {
			mockAPI := mock.NewMockAPI("./endpoint/market/trades")
			subject, _ := mockAPI.MarketTrades(mock.Symbol("do_not_support"))
			It("should return error message", func() {
				Expect(subject).To(Equal(`{"status":40003,"msg":"invalid symbol"}`))
			})
		})
	})

	Describe("MarketCandles", func() {
		Context("when symbol and resolution given", func() {
			mockAPI := mock.NewMockAPI("./endpoint/market/candles")
			subject, _ := mockAPI.MarketCandles(mock.Symbol("ethusdt"), mock.Resolution("MN"))
			It("should return error message", func() {
				Expect(subject).To(Equal(`{"status":0,"data":[{"open":417.930000000,"close":408.470000000,"high":418.740000000,"quote_vol":14185815.897998743670000000,"id":1533052800,"count":211190,"low":400.020000000,"seq":10787977000006,"base_vol":34785.127187351},{"id":1530374400,"seq":10277314000000,"high":516.090000000000000000,"low":418.290000000000000000,"open":449.100000000000000000,"close":430.100000000000000000,"count":6460281,"base_vol":3214156.779928503000000000,"quote_vol":1481280928.428273525940000000},{"id":1527782400,"seq":7385778300000,"high":555.130000000000000000,"low":406.110000000000000000,"open":523.970000000000000000,"close":449.040000000000000000,"count":20382481,"base_vol":28258854.226292702000000000,"quote_vol":13579650310.735054759560000000},{"id":1525104000,"seq":2847062700000,"high":582.990000000000000000,"low":559.660000000000000000,"open":563.230000000000000000,"close":575.770000000000000000,"count":4617,"base_vol":4159.317883344000000000,"quote_vol":2380070.431548222800000000}]}`))
			})
		})

		Context("when resolution is invalid", func() {
			mockAPI := mock.NewMockAPI("./endpoint/market/candles")
			subject, _ := mockAPI.MarketCandles(mock.Symbol("ethusdt"), mock.Resolution("invalid_resolution"))
			It("shoul return error message", func() {
				Expect(subject).To(Equal(`{"status":40003,"msg":"invalid period"}`))
			})
		})
	})
})
