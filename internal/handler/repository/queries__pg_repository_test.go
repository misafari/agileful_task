package repository

import (
	"agileful_task/internal/core/domain"
	"agileful_task/internal/core/port"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"regexp"
	"testing"
	"time"
)

const (
	countQuery                         = `SELECT count(*) FROM "queries_info"`
	simpleQuery                        = `SELECT * FROM "queries_info"`
	queryWithSort                      = `SELECT * FROM "queries_info" ORDER BY time_spent %s`
	queryWithWhere                     = `SELECT * FROM "queries_info" WHERE type = $1`
	queryWithPagination                = `SELECT * FROM "queries_info" LIMIT %d OFFSET %d`
	queryWithWhereAndSort              = `SELECT * FROM "queries_info" WHERE type = $1 ORDER BY time_spent %s`
	queryWithWhereAndSortAndPagination = `SELECT * FROM "queries_info" WHERE type = $1 ORDER BY time_spent %s LIMIT %d OFFSET %d`
)

func TestCount(t *testing.T) {
	repository, sqlMock, getRepoErr := getRepository()
	assert.Nil(t, getRepoErr)

	sqlMock.ExpectQuery(regexp.QuoteMeta(countQuery)).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(10))

	count, err := repository.Count()
	assert.Nil(t, err)

	assert.Equal(t, int64(10), count)
	assert.Nil(t, sqlMock.ExpectationsWereMet())
}

func TestGetQueries_Complex_GoodGuy(t *testing.T) {
	repository, sqlMock, getRepoErr := getRepository()
	assert.Nil(t, getRepoErr)

	tests := []struct {
		name          string
		expectedQuery string
		queryOption   domain.QueryOption
	}{
		{
			name:          "Should be ok (Where Clause only)",
			expectedQuery: queryWithWhere,
			queryOption: domain.QueryOption{
				Sort: domain.NONE,
				Type: domain.SELECT,
			},
		},
		{
			name:          "Should be ok (Where Clause and Sort)",
			expectedQuery: fmt.Sprintf(queryWithWhereAndSort, domain.ASC),
			queryOption: domain.QueryOption{
				Sort: domain.ASC,
				Type: domain.SELECT,
			},
		},
		{
			name:          "Should be ok (Where Clause and Sort)",
			expectedQuery: fmt.Sprintf(queryWithWhereAndSortAndPagination, domain.ASC, 10, 20),
			queryOption: domain.QueryOption{
				Sort: domain.ASC,
				Type: domain.SELECT,
				Pagination: &domain.Pagination{
					Page:     3,
					PageSize: 10,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlMock.ExpectQuery(regexp.QuoteMeta(tt.expectedQuery)).WillReturnRows(sqlmock.NewRows([]string{}))

			_, err := repository.FindAll(&tt.queryOption)
			assert.Nil(t, err)

			assert.Nil(t, sqlMock.ExpectationsWereMet())
		})
	}
}

func TestGetQueries_Sort_Where_GoodGuy(t *testing.T) {
	repository, sqlMock, getRepoErr := getRepository()
	assert.Nil(t, getRepoErr)

	tests := []struct {
		name          string
		expectedQuery string
		queryOption   domain.QueryOption
	}{
		{
			name:          "Should be ok (With Where Clause And Sort)",
			expectedQuery: fmt.Sprintf(queryWithWhereAndSort, domain.ASC),
			queryOption: domain.QueryOption{
				Sort: domain.ASC,
				Type: domain.SELECT,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlMock.ExpectQuery(regexp.QuoteMeta(tt.expectedQuery)).WillReturnRows(sqlmock.NewRows([]string{}))

			_, err := repository.FindAll(&tt.queryOption)
			assert.Nil(t, err)

			assert.Nil(t, sqlMock.ExpectationsWereMet())
		})
	}
}

func TestGetQueries_Where_GoodGuy(t *testing.T) {
	repository, sqlMock, getRepoErr := getRepository()
	assert.Nil(t, getRepoErr)

	tests := []struct {
		name          string
		expectedQuery string
		queryOption   domain.QueryOption
	}{
		{
			name:          "Should be ok (With Where Clause)",
			expectedQuery: queryWithWhere,
			queryOption: domain.QueryOption{
				Type: domain.SELECT,
			},
		},
		{
			name:          "Should be ok (Without Where Clause)",
			expectedQuery: simpleQuery,
			queryOption: domain.QueryOption{
				Type: domain.All,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlMock.ExpectQuery(regexp.QuoteMeta(tt.expectedQuery)).WillReturnRows(sqlmock.NewRows([]string{}))

			_, err := repository.FindAll(&tt.queryOption)
			assert.Nil(t, err)

			assert.Nil(t, sqlMock.ExpectationsWereMet())
		})
	}
}

func TestGetQueries_Pagination_GoodGuy(t *testing.T) {
	repository, sqlMock, getRepoErr := getRepository()
	assert.Nil(t, getRepoErr)

	tests := []struct {
		name          string
		expectedQuery string
		queryOption   domain.QueryOption
	}{
		{
			name:          "Should be ok",
			expectedQuery: fmt.Sprintf(queryWithPagination, 10, 20),
			queryOption: domain.QueryOption{
				Type: domain.All,
				Sort: domain.NONE,
				Pagination: &domain.Pagination{
					Page:     3,
					PageSize: 10,
				},
			},
		},
		{
			name:          "Should be ok",
			expectedQuery: fmt.Sprintf(queryWithPagination, 10, 40),
			queryOption: domain.QueryOption{
				Type: domain.All,
				Sort: domain.NONE,
				Pagination: &domain.Pagination{
					Page:     5,
					PageSize: 10,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlMock.ExpectQuery(regexp.QuoteMeta(tt.expectedQuery)).WillReturnRows(sqlmock.NewRows([]string{}))

			_, err := repository.FindAll(&tt.queryOption)
			assert.Nil(t, err)

			assert.Nil(t, sqlMock.ExpectationsWereMet())
		})
	}
}

func TestGetQueries_Sort_GoodGuy(t *testing.T) {
	repository, sqlMock, getRepoErr := getRepository()
	assert.Nil(t, getRepoErr)

	tests := []struct {
		name          string
		expectedQuery string
		queryOption   domain.QueryOption
	}{
		{
			name:          "Should be ok (ASC)",
			expectedQuery: fmt.Sprintf(queryWithSort, domain.ASC),
			queryOption: domain.QueryOption{
				Sort: domain.ASC,
			},
		},
		{
			name:          "Should be ok (DESC)",
			expectedQuery: fmt.Sprintf(queryWithSort, domain.DESC),
			queryOption: domain.QueryOption{
				Sort: domain.DESC,
			},
		},
		{
			name:          "Should be ok (Without Sort)",
			expectedQuery: simpleQuery,
			queryOption: domain.QueryOption{
				Sort: domain.NONE,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlMock.ExpectQuery(regexp.QuoteMeta(tt.expectedQuery)).WillReturnRows(sqlmock.NewRows([]string{}))

			_, err := repository.FindAll(&tt.queryOption)
			assert.Nil(t, err)

			assert.Nil(t, sqlMock.ExpectationsWereMet())
		})
	}
}

func TestSimpleGetQueries_GoodGuy(t *testing.T) {
	repository, sqlMock, getRepoErr := getRepository()
	assert.Nil(t, getRepoErr)

	tests := []struct {
		name              string
		expectedResultRow *sqlmock.Rows
		expectedResultLen int
	}{
		{
			name: "Should Be ok (With Result)",
			expectedResultRow: sqlmock.NewRows([]string{"id", "type", "time_spent"}).
				AddRow(1, domain.INSERT, 1*time.Second).
				AddRow(2, domain.SELECT, 2*time.Second).
				AddRow(2, domain.DELETE, 1*time.Second).
				AddRow(2, domain.UPDATE, 2*time.Second),
			expectedResultLen: 4,
		},
		{
			name:              "Should Be ok (Without Result)",
			expectedResultRow: sqlmock.NewRows([]string{}),
			expectedResultLen: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlMock.ExpectQuery(regexp.QuoteMeta(simpleQuery)).WillReturnRows(tt.expectedResultRow)

			result, err := repository.FindAll(nil)
			assert.Nil(t, err)

			assert.Equal(t, tt.expectedResultLen, len(result))
			assert.Nil(t, sqlMock.ExpectationsWereMet())
		})
	}
}

func TestSimpleGetQueries_BadGuy(t *testing.T) {
	repository, sqlMock, getRepoErr := getRepository()
	assert.Nil(t, getRepoErr)
	sqlMock.ExpectQuery(regexp.QuoteMeta(simpleQuery)).WillReturnError(errors.New("fake error"))

	_, err := repository.FindAll(nil)

	assert.NotNil(t, err)
	assert.Nil(t, sqlMock.ExpectationsWereMet())
}

func getRepository() (port.QueriesDatabaseRepository, sqlmock.Sqlmock, error) {
	gdb, mock, err := createGormAndSqlMock()
	if err != nil {
		return nil, nil, err
	}
	return NewQueriesDatabaseRepository(gdb, false), mock, nil
}

func createGormAndSqlMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		return nil, nil, err
	}

	director := postgres.New(postgres.Config{
		Conn:                 sqlDB,
		PreferSimpleProtocol: true,
	})

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	gdb, gormErr := gorm.Open(director, &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 newLogger,
	})
	if gormErr != nil {
		return nil, nil, err
	}

	return gdb, mock, nil
}
