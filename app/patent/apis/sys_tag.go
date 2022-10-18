package apis

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"go-admin/app/admin/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/patent/service"
	"go-admin/app/patent/service/dto"
)

type SysTag struct {
	api.Api
}

// Get
// @Summary 获取Tag数据
// @Description 获取JSON
// @Tags 标签/Tag
// @Param tagId path string false "tagId"
// @Router /api/v1/tag/{id} [get]
// @Security Bearer
func (e SysTag) Get(c *gin.Context) {
	s := service.SysTag{}
	req := dto.SysTagGetReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, fmt.Sprintf(" %s ", err.Error()))
		return
	}
	var object models.SysTag

	err = s.Get(&req, &object)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(object, "查询成功")
}

// Insert
// @Summary 增加标签
// @Description 获取JSON
// @Tags 标签/Tag
// @Param data body dto.SysTagInsertReq true "标签数据"
// @Router /api/v1/tag [post]
// @Security Bearer
func (e SysTag) Insert(c *gin.Context) {
	s := service.SysTag{}
	req := dto.SysTagInsertReq{}
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
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))
	err = s.Insert(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Delete
// @Summary 删除标签
// @Description 获取JSON
// @Tags 标签/Tag
// @Param data body dto.ObjectById true "标签数据"
// @Router /api/v1/tag [delete]
// @Security Bearer
func (e SysTag) Delete(c *gin.Context) {
	s := service.SysTag{}
	req := dto.SysTagById{}
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

	// 设置编辑人
	req.SetUpdateBy(user.GetUserId(c))

	// 数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetId(), "删除成功")
}

// Update
// @Summary 更新标签
// @Description 获取JSON
// @Tags 标签/Tag
// @Param data body dto.SysTagUpdateReq true "标签数据"
// @Router /api/v1/tag [put]
// @Security Bearer
func (e SysTag) Update(c *gin.Context) {
	s := service.SysTag{}
	req := dto.SysTagUpdateReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
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

	err = s.Update(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.GetId(), "更新成功")
}
