package dto

import (
	"github.com/ihangsen/common/src/utils/strs"
)

type PageDto struct {
	PageNum   int    // 当前页码
	PageSize  int    `binding:"required,gte=1,lte=100"`  // 每页展示数量
	OrderBy   string `binding:"required"`                // 排序字段
	OrderType string `binding:"required,oneof=DESC ASC"` // 排序方式 DESC/ASC
}

func (dto *PageDto) GetOffset() int {
	return dto.PageNum * dto.PageSize
}

func (dto *PageDto) GetOrder() string {
	return strs.Join(dto.OrderBy, " ", dto.OrderType)
}

type IdPageDto struct {
	LastId   uint64 `json:"lastId"`
	PageSize int    `json:"pageSize" binding:"required,gte=1,lte=100"` // 每页展示数量
}

type SearchPageDto struct {
	PageNum  int64 `json:"pageNum" binding:"lte=100"`                 // 当前页码
	PageSize int64 `json:"pageSize" binding:"required,gte=1,lte=100"` // 每页展示数量
}

func (d SearchPageDto) Offset() int64 {
	return d.PageSize * d.PageNum
}
