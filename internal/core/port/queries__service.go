package port

import "agileful_task/internal/core/domain"

type QueriesService interface {
	GetCount() (int64, error)
	GetAll(*domain.QueryOption) ([]domain.QueriesInfo, error)
}
