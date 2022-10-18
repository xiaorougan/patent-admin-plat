package models

import "go-admin/common/models"

type PatentTag struct {
	TagId   int    `json:"tagId" gorm:"primaryKey;autoIncrement"` //标签ID
	TagName string `json:"tagName" gorm:"size:128;"`              //标签名称
	Desc    string `json:"desc"  gorm:"size:255"  `               //标签描述
	models.ControlBy
	models.ModelTime
}

func (PatentTag) TableName() string {
	return "Tag_table"

}
func (e *PatentTag) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *PatentTag) GetId() interface{} {
	return e.TagId
}
