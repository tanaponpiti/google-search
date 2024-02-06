package model

type PaginationRequest[T any] struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Filter   *T  `json:"filter"`
}

func NewPaginationRequest[T any]() PaginationRequest[T] {
	return PaginationRequest[T]{
		Page:     1,
		PageSize: 10,
	}
}

type PageResponse[T any, D any] struct {
	Pagination PaginationRequest[T] `json:"pagination"`
	Data       []D                  `json:"data"`
}
