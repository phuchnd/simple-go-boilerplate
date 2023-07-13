package http

import (
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("Server", func() {
	//appConfig := config.ServerConfig{
	//	HTTPPort: 1234,
	//	Name:     "Simple Test",
	//	Env:      config.LocalEnv,
	//}
	//BeforeEach(func() {
	//
	//})
	//It("should not be nil when init", func() {
	//	h := handler.NewHandler(c)
	//	s := NewServer(h, appConfig, metricReporter)
	//
	//	Expect(s).ShouldNot(BeNil())
	//})
	//
	//Describe("handleV0Order", func() {
	//	Describe("with correct input", func() {
	//		It("should return 200 with not empty body when Order() return ok", func() {
	//			w := httptest.NewRecorder()
	//			ctx, _ := gin.CreateTestContext(w)
	//			ctx.Request, _ = http.NewRequest("POST", "/v0/order", bytes.NewBuffer([]byte(`{"status":1,"restaurant_id":886562,"update_type":2,"order_code":"06121-531111111","pick_time":"2021-12-06 14:59:00","serial":"0x0000000000000006"}`)))
	//			h := new(handler.MockIHandler)
	//			h.On("HandleOrderWebhookV0", ctx.Request.Context(), dto.OrderWebhookRequest{
	//				OrderCode:           "06121-531111111",
	//				UpdateType:          dto.UPDATE_ORDER_STATUS,
	//				RestaurantID:        886562,
	//				PickTime:            "2021-12-06 14:59:00",
	//				Status:              dto.PICKED,
	//				MerchantNote:        "",
	//				NoteForShipper:      "",
	//				PartnerRestaurantID: "",
	//				Serial:              "0x0000000000000006",
	//			}).Return(&dto.OrderWebhookResponse{}, nil)
	//			s := NewServer(h, appConfig, metricReporter)
	//			s.(*serverImpl).handleV0Order(ctx)
	//
	//			Expect(w.Body.Bytes()).ShouldNot(BeNil())
	//			Expect(w.Code).Should(Equal(http.StatusOK))
	//		})
	//		It("should return 500 when Order() return error", func() {
	//			w := httptest.NewRecorder()
	//			ctx, _ := gin.CreateTestContext(w)
	//			ctx.Request, _ = http.NewRequest("POST", "/v0/order", bytes.NewBuffer([]byte(`{"status":1,"restaurant_id":886562,"update_type":2,"order_code":"06121-531111111","pick_time":"2021-12-06 14:59:00","serial":"0x0000000000000006"}`)))
	//			h := new(handler.MockIHandler)
	//			h.On("HandleOrderWebhookV0", ctx.Request.Context(), dto.OrderWebhookRequest{
	//				OrderCode:           "06121-531111111",
	//				UpdateType:          dto.UPDATE_ORDER_STATUS,
	//				RestaurantID:        886562,
	//				PickTime:            "2021-12-06 14:59:00",
	//				Status:              dto.PICKED,
	//				MerchantNote:        "",
	//				NoteForShipper:      "",
	//				PartnerRestaurantID: "",
	//				Serial:              "0x0000000000000006",
	//			}).Return(nil, errors.New("error random"))
	//			s := NewServer(h, appConfig, metricReporter)
	//			s.(*serverImpl).handleV0Order(ctx)
	//
	//			Expect(w.Body.Bytes()).ShouldNot(BeNil())
	//			Expect(w.Code).Should(Equal(http.StatusInternalServerError))
	//		})
	//	})
	//	Describe("with incorrect input", func() {
	//		gin.SetMode(gin.TestMode)
	//		It("should return 422 when input is empty", func() {
	//			w := httptest.NewRecorder()
	//			ctx, _ := gin.CreateTestContext(w)
	//			ctx.Request, _ = http.NewRequest("POST", "/v0/order", bytes.NewBuffer([]byte(``)))
	//			h := handler.NewHandler(c)
	//			s := NewServer(h, appConfig, metricReporter)
	//			s.(*serverImpl).handleV0Order(ctx)
	//
	//			Expect(w.Body.Bytes()).Should(BeNil())
	//			Expect(w.Code).Should(Equal(http.StatusUnprocessableEntity))
	//		})
	//	})
	//})
})
