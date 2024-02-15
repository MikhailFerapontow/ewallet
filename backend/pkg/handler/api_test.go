package handler

import (
	"errors"
	"net/http/httptest"
	"server/internal/logger/slogdiscard"
	"server/models"
	"server/pkg/service"
	mock_service "server/pkg/service/mocks"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	gomock "go.uber.org/mock/gomock"
)

func TestNewWallet(t *testing.T) {
	type mockBehavior func(s *mock_service.MockApi)

	tests := []struct {
		name                string
		mockBehavior        mockBehavior
		expectedRequestBody models.Wallet
		expectedStatusCode  int
	}{
		{
			name: "ok",
			mockBehavior: func(s *mock_service.MockApi) {
				s.EXPECT().NewWallet().Return(models.Wallet{
					Id:      "8a95472c2b026a995fa016aa119ea719",
					Balance: 100,
				}, nil)
			},
			expectedRequestBody: models.Wallet{
				Id:      "8a95472c2b026a995fa016aa119ea719",
				Balance: 100,
			},
			expectedStatusCode: 200,
		},
		{
			name: "error",
			mockBehavior: func(s *mock_service.MockApi) {
				s.EXPECT().NewWallet().Return(models.Wallet{}, errors.New("error"))
			},
			expectedRequestBody: models.Wallet{},
			expectedStatusCode:  400,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			api := mock_service.NewMockApi(c)
			testCase.mockBehavior(api)

			services := &service.Service{
				Api: api,
			}
			handler := NewHandler(services, slogdiscard.NewDiscardLogger())

			// Test Server
			r := gin.New()
			r.POST("/api/v1/wallet", handler.newWallet)

			// Create Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/wallet", nil)

			// Make Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body)
		})
	}
}
