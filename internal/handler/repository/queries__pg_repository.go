package repository

import (
	"agileful_task/internal/core/domain"
	"agileful_task/internal/core/port"
	"fmt"
	"gorm.io/gorm"
	"log"
)

type queriesDatabaseRepository struct {
	db *gorm.DB
}

func (q *queriesDatabaseRepository) Count() (int64, error) {
	var count int64
	err := q.db.Model(&domain.QueriesInfo{}).Count(&count).Error
	return count, err
}

func (q *queriesDatabaseRepository) FindAll(options *domain.QueryOption) ([]domain.QueriesInfo, error) {
	var infos []domain.QueriesInfo

	d := q.db.Model(domain.QueriesInfo{})

	if options == nil {
		err := d.Find(&infos).Error
		return infos, err
	}

	if options.Type != domain.All {
		d.Where("type = ?", options.Type)
	}

	if options.Sort != domain.NONE {
		d.Order(fmt.Sprintf("%s %s", "time_spent", options.Sort))
	}

	if options.Pagination != nil {
		p := options.Pagination.Page
		if p == 0 {
			p = 1
		}

		ps := options.Pagination.PageSize
		switch {
		case ps > 100:
			ps = 100
		case ps <= 0:
			ps = 10
		}

		offset := (p - 1) * ps

		d.Offset(offset).Limit(ps)
	}

	err := d.Find(&infos).Error
	return infos, err
}

func NewQueriesDatabaseRepository(db *gorm.DB, withAutoMigrate bool) port.QueriesDatabaseRepository {
	if withAutoMigrate {
		if err := db.AutoMigrate(&domain.QueriesInfo{}); err != nil {
			log.Panic(err)
		}
	}

	return &queriesDatabaseRepository{
		db: db,
	}
}
