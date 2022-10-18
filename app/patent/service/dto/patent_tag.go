package dto

import (
	"go-admin/app/patent/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type PatentTagInsertReq struct {
	TagId   int    `json:"tagId" comment:"标签ID"`   //标签ID
	TagName string `json:"tagName" comment:"标签名;"` //标签名称
	Desc    string `json:"desc"    comment:"标签描述"` //标签描述
	common.ControlBy
}

type PatentTagUpdateReq struct {
	TagId   int    `json:"tagId" comment:"标签ID"`   //标签ID
	TagName string `json:"tagName" comment:"标签名;"` //标签名称
	Desc    string `json:"desc"    comment:"标签描述"` //标签描述
	common.ControlBy
}

type PatentTagById struct {
	dto.ObjectById
	common.ControlBy
}

func (s *PatentTagInsertReq) Generate(model *models.PatentTag) {
	if s.TagId != 0 {
		model.TagId = s.TagId
	}
	model.TagName = s.TagName
	model.Desc = s.TagName
}

func (s *PatentTagUpdateReq) Generate(model *models.PatentTag) {
	if s.TagId != 0 {
		model.TagId = s.TagId
	}
	model.TagName = s.TagName
	model.Desc = s.TagName
}

func (s *PatentTagInsertReq) GetId() interface{} {
	return s.TagId
}

func (s *PatentTagUpdateReq) GetId() interface{} {
	return s.TagId
}

type PatentTagGetReq struct {
	Id int `uri:"id"`
}

func (s *PatentTagGetReq) GetId() interface{} {
	return s.Id
}
