package port

import "agileful_task/internal/core/domain"

type QueriesDatabaseRepository interface {
	Count() (int64, error)
	FindAll(*domain.QueryOption) ([]domain.QueriesInfo, error)
}
