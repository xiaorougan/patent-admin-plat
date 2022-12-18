package dtos

import "go-admin/common/dto"

type SimpleSearchReq struct {
	dto.Pagination

	Query string `json:"query"`
	DB    string `json:"DB"`

	UserId int `json:"-"`
}

type TableSearchReq struct {
	dto.Pagination

	Query string `json:"query"`
	DB    string `json:"DB"`
}
