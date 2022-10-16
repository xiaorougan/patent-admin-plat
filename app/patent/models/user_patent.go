package models

import "go-admin/common/models"

type SysUserPatent struct {
	models.Model
	PatentId int    `gorm:"foreignKey:PatentId" json:"PatentId" `
	UserId   int    `gorm:"primaryKey;autoIncrement;comment:编码"  json:"userId"`
	ID       int    `json:"ID" gorm:"size:128;comment:关系ID"`
	Type     string `json:"Type" gorm:"size:128;comment:关系类型（关注/认领）"`
	models.ControlBy
}

func (e *SysUserPatent) TableName() string {
	return "user_patent"
}

func (e *SysUserPatent) GetUserId() interface{} {
	return e.UserId
}

func (e *SysUserPatent) GetPatentId() interface{} {
	return e.PatentId
}
