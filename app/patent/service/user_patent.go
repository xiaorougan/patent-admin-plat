package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/patent/models"
	"go-admin/app/patent/service/dto"
)

type UserPatent struct {
	service.Service
}

// 通过userid获得patentid列表，通过patentid获取patent列表

// GetClaimLists 通过userid获得patentid列表，通过patentid获取patent列表
func (e *UserPatent) GetClaimLists(c *dto.UserPatentGetPageReq, list *[]models.UserPatent, count *int64) error {
	var err error
	var data models.UserPatent
	err = e.Orm.Model(&data).Select("patent_id").
		Where("Type = ? AND user_id = ?", "认领", c.GetUserId()).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetCollectionLists 通过userid获得patentid列表，通过patentid获取patent列表
func (e *UserPatent) GetCollectionLists(c *dto.UserPatentGetPageReq, list *[]models.UserPatent, count *int64) error {
	var err error
	var data models.UserPatent
	err = e.Orm.Model(&data).
		Where("Type = ? AND user_id = ?", "关注", c.GetUserId()).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetPatentPagesByIds 获取patent列表
func (e *UserPatent) GetPatentPagesByIds(d *dto.PatentsByIdsForRelationshipUsers, list *[]models.Patent, count *int64) error {
	var err error
	var ids []int = d.GetPatentId()
	for i := 0; i < len(ids); i++ {
		if ids[i] != 0 {
			var data1 models.Patent
			err = e.Orm.Model(&data1).
				Where("patent_id = ? ", ids[i]).
				First(&data1).Limit(-1).Offset(-1).
				Count(count).Error
			*list = append(*list, data1)
			if err != nil {
				e.Log.Errorf("db error:%s", err)
				return err
			}
		}
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// InsertClaimRelationship 创建认领关系。
func (e *UserPatent) InsertClaimRelationship(c *dto.UserPatentInsertReq) error {
	var err error
	var data models.UserPatent
	var i int64
	err = e.Orm.Model(&data).Where("patent_id = ? AND user_id = ? AND type = ?", c.PatentId, c.UserId, c.Type).
		Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("关系已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}

	c.GenerateUserPatent(&data)
	c.Type = "认领"
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// InsertCollectionRelationship 创建关注关系。
func (e *UserPatent) InsertCollectionRelationship(c *dto.UserPatentInsertReq) error {
	var err error
	var data models.UserPatent
	var i int64
	err = e.Orm.Model(&data).Where("patent_id = ? AND user_id = ? AND type = ?", c.PatentId, c.UserId, c.Type).
		Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("关系已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}

	c.GenerateUserPatent(&data)
	c.Type = "关注"

	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// RemoveRelationship 根据专利id删除Patent（可以自定义根据专利id删除数据的个数，因为post的内容是一个json里面是PatentID的数组）
func (e *UserPatent) RemoveRelationship(c *dto.UserPatentObject) error {
	var err error
	var data models.UserPatent

	db := e.Orm.Delete(&data).
		Where("patent_id = ? AND user_id = ? AND type = ?", c.PatentId, c.UserId, c.Type)

	if db.Error != nil {
		err = db.Error
		e.Log.Errorf("Delete error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("无权删除该数据")
		return err
	}
	return nil
}
