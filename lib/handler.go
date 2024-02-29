package lib

import (
	"fmt"
	"math"
	"strconv"

	"github.com/adamelfsborg-code/food/culinary/data"
)

type Pagination struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
	PageCount int `json:"pageCount"`
}

type cacheable interface {
	data.FoodTableDto | data.BrandDto | data.CategoryDto | data.FoodTypeTableDto
}

func NewPagination(pageIndex, pageSize string) (*Pagination, error) {
	index, err := strconv.Atoi(pageIndex)
	if err != nil {
		fmt.Println("Failed to parse page index: ", err)
		return nil, err
	}

	size, err := strconv.Atoi(pageSize)
	if err != nil {
		fmt.Println("Failed to parse page size: ", err)
		return nil, err
	}

	pagination := &Pagination{
		PageIndex: index,
		PageSize:  size,
	}

	return pagination, nil
}

type PaginatedResponse[T cacheable] struct {
	Rows       []T        `json:"rows"`
	Pagination Pagination `json:"pagination"`
}

func NewPaginatedResponse[T cacheable](rows []T, count int, pagination Pagination) PaginatedResponse[T] {
	setPageCount(&pagination, count)

	response := PaginatedResponse[T]{
		Rows:       rows,
		Pagination: pagination,
	}
	fmt.Println(response.Pagination)
	return response
}

func setPageCount(pagination *Pagination, count int) {
	pagination.PageCount = int(math.Ceil(float64(count) / float64(pagination.PageSize)))
}
