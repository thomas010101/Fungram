package types

import "math"

type BasePage struct {
	PageNo     int64 `json:"pageNo"`
	PageSize   int64 `json:"pageSize"`
	TotalPage  int64 `json:"totalPage"`
	TotalCount int64 `json:"totalCount"`
}

func NewBasePage(pageNo, pageSize, totalCount int64) BasePage {
	return BasePage{
		PageNo:     pageNo,
		PageSize:   pageSize,
		TotalPage:  int64(math.Ceil(float64(totalCount) / float64(pageSize))),
		TotalCount: totalCount,
	}
}
