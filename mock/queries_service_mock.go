package mock

import (
	"agileful_task/internal/core/domain"
	"github.com/golang/mock/gomock"
	"reflect"
)

type QueriesServiceMock struct {
	ctrl     *gomock.Controller
	recorder *QueriesServiceRecorder
}

type QueriesServiceRecorder struct {
	mock *QueriesServiceMock
}

func (m *QueriesServiceMock) EXPECT() *QueriesServiceRecorder {
	return m.recorder
}

func (q *QueriesServiceMock) GetCount() (int64, error) {
	q.ctrl.T.Helper()
	ret := q.ctrl.Call(q, "GetCount")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *QueriesServiceRecorder) GetCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCount", reflect.TypeOf((*QueriesServiceMock)(nil).GetCount))
}

func (q *QueriesServiceMock) GetAll(option *domain.QueryOption) ([]domain.QueriesInfo, error) {
	q.ctrl.T.Helper()
	ret := q.ctrl.Call(q, "GetAll", option)
	ret0, _ := ret[0].([]domain.QueriesInfo)
	ret1, _ := ret[1].(error)

	return ret0, ret1
}

func (mr *QueriesServiceRecorder) GetAll(option *domain.QueryOption) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*QueriesServiceMock)(nil).GetAll), option)
}


func NewQueriesServiceMock(ctrl *gomock.Controller) *QueriesServiceMock {
	mock := &QueriesServiceMock{ctrl: ctrl}
	mock.recorder = &QueriesServiceRecorder{mock}
	return mock;
}
