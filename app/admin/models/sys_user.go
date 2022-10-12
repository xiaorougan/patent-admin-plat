package models

import (
	"go-admin/common/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type SysUser struct {
	UserId    int    `gorm:"primaryKey;autoIncrement;comment:编码"  json:"userId"`
	Username  string `json:"username" gorm:"size:64;comment:用户名"`
	Password  string `json:"-" gorm:"size:128;comment:密码"`
	NickName  string `json:"nickName" gorm:"size:128;comment:昵称"`
	Phone     string `json:"phone" gorm:"size:11;comment:手机号"`
	RoleId    int    `json:"roleId" gorm:"size:20;comment:角色ID"`
	Salt      string `json:"-" gorm:"size:255;comment:加盐"`
	Avatar    string `json:"avatar" gorm:"size:255;comment:头像"`
	Sex       string `json:"sex" gorm:"size:255;comment:性别"`
	Email     string `json:"email" gorm:"size:128;comment:邮箱"`
	Remark    string `json:"remark" gorm:"size:255;comment:备注"`
	Bio       string `json:"interest" gorm:"size:1024;comment:简介"`
	Interest  string `json:"bio" gorm:"size:1024;comment:研究兴趣"`
	Departure string `json:"departure" gorm:"size:128;comment:单位"`
	Status    string `json:"status" gorm:"size:4;comment:状态"`
	RoleIds   []int  `json:"roleIds" gorm:"-"`
	models.ControlBy
	models.ModelTime
}

func (e *SysUser) TableName() string {
	return "sys_user"
}

func (e *SysUser) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SysUser) GetId() interface{} {
	return e.UserId
}

//加密
func (e *SysUser) Encrypt() (err error) {
	if e.Password == "" {
		return
	}

	var hash []byte
	if hash, err = bcrypt.GenerateFromPassword([]byte(e.Password), bcrypt.DefaultCost); err != nil {
		return
	} else {
		e.Password = string(hash)
		return
	}
}

func (e *SysUser) BeforeCreate(_ *gorm.DB) error {
	return e.Encrypt()
}

func (e *SysUser) BeforeUpdate(_ *gorm.DB) error {
	var err error
	if e.Password != "" {
		err = e.Encrypt()
	}
	return err
}

func (e *SysUser) AfterFind(_ *gorm.DB) error {
	e.RoleIds = []int{e.RoleId}
	return nil
}
