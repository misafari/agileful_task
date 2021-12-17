package query_svc

import (
	"agileful_task/internal/core/domain"
	"agileful_task/mock"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCount(t *testing.T) {
	tests := []struct {
		name   string
		result int64
		err    error
	}{
		{
			name:   "Should be ok",
			result: 10,
			err:    nil,
		},
		{
			name:   "Should not be ok",
			result: 0,
			err:    errors.New("fake err"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mock.NewMockQueriesDatabaseRepository(gomock.NewController(t))

			mockRepository.EXPECT().Count().Return(tt.result, tt.err)

			svc := NewQueriesService(mockRepository)

			result, err := svc.GetCount()

			if tt.err != nil {
				assert.NotNil(t, err)
			}

			assert.Equal(t, tt.result, result)
		})
	}
}

func TestGetAll(t *testing.T) {
	tests := []struct {
		name         string
		options      domain.QueryOption
		expectResult []domain.QueriesInfo
		expectErr    error
	}{
		{
			name: "Should Be ok",
			options: domain.QueryOption{
				Sort: domain.ASC,
				Type: domain.All,
			},
			expectResult: []domain.QueriesInfo{{
				ID:        1,
				Type:      domain.INSERT,
				TimeSpent: int64(1 * time.Millisecond),
			}},
			expectErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepository := mock.NewMockQueriesDatabaseRepository(gomock.NewController(t))

			mockRepository.EXPECT().FindAll(&tt.options).Return(tt.expectResult, tt.expectErr)

			svc := NewQueriesService(mockRepository)

			result, err := svc.GetAll(&tt.options)

			if tt.expectErr != nil {
				assert.NotNil(t, err)
			}

			assert.Equal(t, tt.expectResult, result)
		})
	}
}
