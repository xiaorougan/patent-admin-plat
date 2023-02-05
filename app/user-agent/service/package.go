package service

import (
	"errors"
	"fmt"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service/dto"
	"gorm.io/gorm"
)

type Package struct {
	service.Service
}

func (e *Package) ListByUserId(c *dto.PackageListReq, list *[]models.Package) error {
	var err error
	//var data models.Package
	// todo: check
	err = e.Orm.Debug().
		Where("create_by = ?", c.UserId).
		Find(list).Limit(-1).Offset(-1).
		Error

	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	return nil
}

func (e *Package) FindForCurrentUser(c *dto.PackageFindReq, list *[]models.Package) error {
	var err error
	//var data models.Package
	// todo: check
	err = e.Orm.Debug().
		Where("create_by = ? AND package_name LIKE ?", c.UserId, fmt.Sprintf("%%%s%%", c.Query)).
		Find(list).Limit(-1).Offset(-1).
		Error

	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	return nil
}

// Get 获取Package对象
func (e *Package) Get(d *dto.PackageById, model *models.Package) error {
	var data models.Package

	err := e.Orm.Model(&data).Debug().
		//Scopes(
		//	actions.Permission(data.TableName(), p),
		//).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Insert 创建Package对象
func (e *Package) Insert(c *dto.PackageInsertReq) error {
	var err error
	var data models.Package
	var i int64
	err = e.Orm.Model(&data).Where("package_name = ?", c.PackageName).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("专利包名已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// Update 修改Package对象
func (e *Package) Update(c *dto.PackageUpdateReq) error {
	var err error
	var model models.Package
	db := e.Orm.First(&model, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Service UpdateSysUser error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("package not found")

	}

	switch c.FilesOpt {
	case dto.FilesAdd:
		c.GenerateAndAddFiles(&model)
	case dto.FilesDelete:
		c.GenerateAndDeleteFiles(&model)
	default:
		c.Generate(&model)
	}

	update := e.Orm.Model(&model).Where("package_id = ?", &model.PackageId).Updates(&model)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update userinfo error")
		log.Warnf("db update error")
		return err
	}
	return nil
}

// Remove 删除SysUser
func (e *Package) Remove(c *dto.PackageById) error {
	var err error
	var data models.Package

	db := e.Orm.Model(&data).
		Delete(&data, c.GetId())
	if err = db.Error; err != nil {
		e.Log.Errorf("Error found in  RemoveSysUser : %s", err)
		return err
	}
	//if db.RowsAffected == 0 {
	//	return errors.New("无权删除该数据")
	//}
	return nil
}
