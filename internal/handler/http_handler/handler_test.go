package http_handler

import (
	"agileful_task/internal/core/domain"
	"agileful_task/mock"
	"github.com/gofiber/fiber/v2"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCount_http_handler_test(t *testing.T) {
	tests := []struct {
		name             string
		route            string
		expectedCode     int
		expectedResult   int64
		expectedResponse string
	}{
		{
			name:             "get HTTP status 200",
			route:            "/queries/count",
			expectedCode:     200,
			expectedResult:   10,
			expectedResponse: `{"data":10}`,
		},
	}

	app := fiber.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := mock.NewQueriesServiceMock(gomock.NewController(t))

			mockService.EXPECT().GetCount().Return(tt.expectedResult, nil)

			h := NewQueriesHttpHandler(mockService)
			app.Get("/queries/count", h.GetQueriesCountHandler)

			req := httptest.NewRequest("GET", tt.route, nil)
			resp, _ := app.Test(req, 1)

			body, err := ioutil.ReadAll(resp.Body)
			assert.Nil(t, err)

			assert.Equal(t, tt.expectedResponse, string(body))
			assert.Equal(t, tt.expectedCode, resp.StatusCode, tt.name)
		})
	}
}

func TestGetAll_http_handler_test(t *testing.T) {
	tests := []struct {
		name             string
		route            string
		option           domain.QueryOption
		expectedCode     int
		expectedResult   []domain.QueriesInfo
		expectedResponse string
	}{
		{
			name:  "get HTTP status 200",
			route: "/queries?sort=desc&type=SELECT&pagination.page=2&pagination.page_size=10",
			option: domain.QueryOption{
				Type: domain.SELECT,
				Sort: domain.DESC,
				Pagination: &domain.Pagination{
					Page:     2,
					PageSize: 10,
				},
			},
			expectedCode: 200,
			expectedResult: []domain.QueriesInfo{
				{
					ID:        1,
					Type:      domain.SELECT,
					TimeSpent: int64(1 * time.Millisecond),
				},
			},
			expectedResponse: `{"data":[{"ID":1,"Type":"SELECT","TimeSpent":1000000}]}`,
		},
	}

	app := fiber.New()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := mock.NewQueriesServiceMock(gomock.NewController(t))

			mockService.EXPECT().GetAll(&tt.option).Return(tt.expectedResult, nil)

			h := NewQueriesHttpHandler(mockService)
			app.Get("/queries", h.GetQueriesHandler)

			req := httptest.NewRequest("GET", tt.route, nil)
			resp, _ := app.Test(req, 1)

			body, err := ioutil.ReadAll(resp.Body)
			assert.Nil(t, err)

			assert.Equal(t, tt.expectedResponse, string(body))
			assert.Equal(t, tt.expectedCode, resp.StatusCode, tt.name)
		})
	}
}
