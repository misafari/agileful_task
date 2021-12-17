package query_svc

import (
	"agileful_task/internal/core/domain"
	"agileful_task/internal/core/port"
)

type queriesService struct {
	queriesDatabaseRepository port.QueriesDatabaseRepository
}

func (q *queriesService) GetCount() (int64, error) {
	count, err := q.queriesDatabaseRepository.Count()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (q *queriesService) GetAll(option *domain.QueryOption) ([]domain.QueriesInfo, error) {
	result, err := q.queriesDatabaseRepository.FindAll(option)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func NewQueriesService(queriesDatabaseRepository port.QueriesDatabaseRepository) port.QueriesService {
	return &queriesService{
		queriesDatabaseRepository: queriesDatabaseRepository,
	}
}
