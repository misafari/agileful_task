package domain

type Sort string

const (
	NONE Sort = ""
	ASC  Sort = "asc"
	DESC Sort = "desc"
)

type QueryType string

const (
	All    QueryType = ""
	SELECT QueryType = "SELECT"
	INSERT QueryType = "INSERT"
	UPDATE QueryType = "UPDATE"
	DELETE QueryType = "DELETE"
)

type QueryOption struct {
	Type       QueryType   `query:"type"`
	Sort       Sort        `query:"sort"`
	Pagination *Pagination `query:"pagination"`
}

type Pagination struct {
	Page     int `query:"page"`
	PageSize int `query:"page_size"`
}
