package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SysTagInsertReq struct {
	TagId   int    `json:"tagId" comment:"标签ID"`   //标签ID
	TagName string `json:"tagName" comment:"标签名;"` //标签名称
	Desc    string `json:"desc"    comment:"标签描述"` //标签描述
	common.ControlBy
}

type SysTagUpdateReq struct {
	TagId   int    `json:"tagId" comment:"标签ID"`   //标签ID
	TagName string `json:"tagName" comment:"标签名;"` //标签名称
	Desc    string `json:"desc"    comment:"标签描述"` //标签描述
	common.ControlBy
}

type SysTagById struct {
	dto.ObjectById
	common.ControlBy
}

func (s *SysTagInsertReq) Generate(model *models.SysTag) {
	if s.TagId != 0 {
		model.TagId = s.TagId
	}
	model.TagName = s.TagName
	model.Desc = s.TagName
}

func (s *SysTagUpdateReq) Generate(model *models.SysTag) {
	if s.TagId != 0 {
		model.TagId = s.TagId
	}
	model.TagName = s.TagName
	model.Desc = s.TagName
}

func (s *SysTagInsertReq) GetId() interface{} {
	return s.TagId
}

func (s *SysTagUpdateReq) GetId() interface{} {
	return s.TagId
}

type SysTagGetReq struct {
	Id int `uri:"id"`
}

func (s *SysTagGetReq) GetId() interface{} {
	return s.Id
}
