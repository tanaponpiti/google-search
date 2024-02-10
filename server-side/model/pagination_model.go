package model

type PaginationRequest[T any] struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
	Filter   *T  `json:"filter"`
}

type PaginationResponse[T any] struct {
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
	TotalPage int `json:"totalPage"`
	Filter    *T  `json:"filter"`
}

func getTotalPages(totalItems int64, pageSize int) int {
	if pageSize <= 0 {
		return 0
	}
	pageSize64 := int64(pageSize)
	totalPages := totalItems / pageSize64
	if totalItems%pageSize64 != 0 {
		totalPages++
	}
	return int(totalPages)
}

func CreatePaginationResponse[T any](request PaginationRequest[T], totalItems int64) PaginationResponse[T] {
	totalPages := getTotalPages(totalItems, request.PageSize)
	response := PaginationResponse[T]{
		Page:      request.Page,
		PageSize:  request.PageSize,
		TotalPage: totalPages,
		Filter:    request.Filter,
	}
	return response
}

func NewPaginationRequest[T any]() PaginationRequest[T] {
	return PaginationRequest[T]{
		Page:     1,
		PageSize: 10,
	}
}

type PageResponse[T any, D any] struct {
	Pagination PaginationResponse[T] `json:"pagination"`
	Data       []D                   `json:"data"`
}
