package models

import "go-admin/common/models"

type SysTag struct {
	TagId   int    `json:"tagId" gorm:"primaryKey;autoIncrement"` //标签ID
	TagName string `json:"tagName" gorm:"size:128;"`              //标签名称
	Desc    string `json:"desc"  gorm:"size:255"  `               //标签描述
	models.ControlBy
	models.ModelTime
}

func (SysTag) TableName() string {
	return "Tag_table"

}
func (e *SysTag) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysTag) GetId() interface{} {
	return e.TagId
}
