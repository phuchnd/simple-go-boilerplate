package http

import (
	"errors"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/phuchnd/simple-go-boilerplate/internal/config"
	"github.com/phuchnd/simple-go-boilerplate/internal/config/mocks"
	mocks2 "github.com/phuchnd/simple-go-boilerplate/internal/db/mysql/mocks"
	"github.com/phuchnd/simple-go-boilerplate/internal/logging"
	"github.com/phuchnd/simple-go-boilerplate/internal/service/http/entities"
	mocks3 "github.com/phuchnd/simple-go-boilerplate/internal/service/http/mocks"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Server", func() {
	var (
		logger       logging.Logger
		cfgProvider  *mocks.IConfig
		db           *mocks2.IMySqlDB
		handler      *mocks3.IHTTPService
		serverConfig = &config.ServerConfig{
			HTTPPort: 1234,
			Name:     "Simple Test",
			Env:      config.LocalEnv,
		}
	)
	BeforeEach(func() {
		cfgProvider = new(mocks.IConfig)
		cfgProvider.On("GetServerConfig").Return(serverConfig)
		cfgProvider.On("GetDBConfig").Return(&config.DBConfig{})
		cfgProvider.On("GetBookConfig").Return(&config.BookConfig{})
		cfgProvider.On("GetCronHealthCheckConfig").Return(&config.CronConfig{})
		db = new(mocks2.IMySqlDB)
		logger = logging.NewLogger(cfgProvider)
	})
	It("should not be nil when init", func() {
		s, err := NewServer(logger, cfgProvider, db, handler)

		Expect(s).ShouldNot(BeNil())
		Expect(err).Should(BeNil())
	})

	Describe("listBookV0", func() {
		Describe("with correct input", func() {
			It("should return 200 with not empty body when ListBooks() return ok", func() {
				w := httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(w)
				ctx.Request, _ = http.NewRequest("GET", "/v0/books?limit=5&cursor=0", nil)
				handler = new(mocks3.IHTTPService)
				handler.On("ListBooks", ctx.Request.Context(), &entities.ListBookRequest{
					Limit:  5,
					Cursor: 0,
				}).Return(&entities.ListBookResponse{}, nil)

				s, err := NewServer(logger, cfgProvider, db, handler)

				Expect(s).ShouldNot(BeNil())
				Expect(err).Should(BeNil())

				s.(*httpServerImpl).listBookV0(ctx)

				Expect(w.Body.Bytes()).ShouldNot(BeNil())
				Expect(w.Code).Should(Equal(http.StatusOK))
			})
			It("should return 500 when Order() return error", func() {
				w := httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(w)
				ctx.Request, _ = http.NewRequest("GET", "/v0/books?limit=5&cursor=0", nil)
				handler = new(mocks3.IHTTPService)
				handler.On("ListBooks", ctx.Request.Context(), &entities.ListBookRequest{
					Limit:  5,
					Cursor: 0,
				}).Return(nil, errors.New("err ListBooks"))

				s, err := NewServer(logger, cfgProvider, db, handler)

				Expect(s).ShouldNot(BeNil())
				Expect(err).Should(BeNil())

				s.(*httpServerImpl).listBookV0(ctx)

				Expect(w.Body.Bytes()).ShouldNot(BeNil())
				Expect(w.Code).Should(Equal(http.StatusInternalServerError))
			})
		})
		Describe("with incorrect input", func() {
			gin.SetMode(gin.TestMode)
			It("should return 400 when limit is empty", func() {
				w := httptest.NewRecorder()
				ctx, _ := gin.CreateTestContext(w)
				ctx.Request, _ = http.NewRequest("GET", "/v0/books", nil)
				handler = new(mocks3.IHTTPService)
				handler.On("ListBooks", ctx.Request.Context(), &entities.ListBookRequest{
					Limit:  5,
					Cursor: 0,
				}).Return(nil, errors.New("err ListBooks"))

				s, err := NewServer(logger, cfgProvider, db, handler)

				Expect(s).ShouldNot(BeNil())
				Expect(err).Should(BeNil())

				s.(*httpServerImpl).listBookV0(ctx)

				Expect(w.Body.Bytes()).ShouldNot(BeNil())
				Expect(w.Code).Should(Equal(http.StatusBadRequest))
			})
		})
	})
})
