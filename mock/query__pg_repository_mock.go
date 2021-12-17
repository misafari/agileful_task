package mock

import (
	"agileful_task/internal/core/domain"
	"github.com/golang/mock/gomock"
	"reflect"
)

type QueriesDatabaseRepositoryMock struct {
	ctrl     *gomock.Controller
	recorder *QueriesDatabaseRepositoryRecorder
}

type QueriesDatabaseRepositoryRecorder struct {
	mock *QueriesDatabaseRepositoryMock
}

func (m *QueriesDatabaseRepositoryMock) EXPECT() *QueriesDatabaseRepositoryRecorder {
	return m.recorder
}

func (m *QueriesDatabaseRepositoryMock) Count() (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *QueriesDatabaseRepositoryRecorder) Count() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*QueriesDatabaseRepositoryMock)(nil).Count))
}

func (m *QueriesDatabaseRepositoryMock) FindAll(option *domain.QueryOption) ([]domain.QueriesInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll", option)
	ret0, _ := ret[0].([]domain.QueriesInfo)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *QueriesDatabaseRepositoryRecorder) FindAll(option *domain.QueryOption) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*QueriesDatabaseRepositoryMock)(nil).FindAll), option)
}

func NewMockQueriesDatabaseRepository(ctrl *gomock.Controller) *QueriesDatabaseRepositoryMock {
	mock := &QueriesDatabaseRepositoryMock{ctrl: ctrl}
	mock.recorder = &QueriesDatabaseRepositoryRecorder{mock}
	return mock;
}