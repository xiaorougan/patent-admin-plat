package service

import (
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"go-admin/app/patent/models"
	"go-admin/app/patent/service/dto"
	cDto "go-admin/common/dto"
)

type PatentTag struct {
	service.Service
}

//通过Tagid获取PatentId列表，通过Patentid获取Patent列表
//通过PatentId获取TagId列表，通过TagId获取Tag列表

// GetPage 通过userid获得patentid列表，通过patentid获取patent列表
func (e *PatentTag) GetPage(c *dto.PatentTagGetPageReq, list *[]models.PatentTag, count *int64) error {
	var err error
	var data models.PatentTag
	//
	//var patentids []models.PatentTag = make([]models.PatentTag, 0)
	//db := e.Orm.First(model, d.GetPatentId())
	//err = db.Error
	//patentids = e.Orm.Model(&data).Where()
	//err = patentids.Error
	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}
