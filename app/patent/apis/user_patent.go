package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/patent/service"
	"go-admin/app/patent/service/dto"
)

type UserPatent struct {
	api.Api
}

// DeleteUserPatentRelationship
// @Summary 根据专利id、TYPE删除用户专利关系
// @Description  根据专利id、TYPE删除用户专利关系
// @Tags 用户专利关系表
// @Param PatentId query string false "专利ID"
// @Param Type query string false "关系类型"
// @Router /api/v1/user-agent/user-patent/{patent_id}/{type} [delete]
// @Security Bearer
func (e UserPatent) DeleteUserPatentRelationship(c *gin.Context) {
	s := service.UserPatent{}
	req := dto.UserPatentObject{}
	req.UserId = user.GetUserId(c)

	req.SetUpdateBy(user.GetUserId(c))

	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req). //在这一步传入request数据
		MakeService(&s.Service).
		Errors

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// 数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.RemoveRelationship(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req, "删除成功")
}

// UpdateUserPatentRelationship
// @Summary 修改用户专利关系
// @Description 需要输入专利id更新用户专利关系
// @Tags 用户专利关系表
// @Accept  application/json
// @Product application/json
// @Param data body dto.UpDateUserPatentObject true "body"
// @Router /api/v1/user-agent/user-patent [put]
// @Security Bearer
func (e UserPatent) UpdateUserPatentRelationship(c *gin.Context) {
	s := service.UserPatent{}
	req := dto.UpDateUserPatentObject{}
	req.UserId = user.GetUserId(c)
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, binding.JSON).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.SetUpdateBy(user.GetUserId(c))
	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	if req.PatentId == 0 {
		e.Logger.Error(err)
		e.Error(404, err, "请输入专利id")
		return
	}

	err = s.UpdateUserPatent(&req)

	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req, "更新成功")
}
