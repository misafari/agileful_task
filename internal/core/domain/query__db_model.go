package domain

type QueriesInfo struct {
	ID        uint `gorm:"primaryKey"`
	Type      QueryType
	TimeSpent int64
}

func (QueriesInfo) TableName() string {
	return "queries_info"
}
