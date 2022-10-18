package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"go-admin/app/patent/service"
	"go-admin/app/patent/service/dto"
)

type Search struct {
	api.Api
}

// SimpleSearch
// @Summary 简单查询
// @Description 根据查询字符串进行模糊搜索
// @Tags 专利检索
// @Param data body dto.SimpleSearchReq true "用户数据"
// @Success 200 {object} dto.SwagSimpleSearchResp
// @Router /api/v1/search/simple [post]
// @Security Bearer
func (e Search) SimpleSearch(c *gin.Context) {
	ic := service.GetCurrentInnojoy()
	s := service.Search{}
	req := dto.SimpleSearchReq{}

	err := e.MakeContext(c).
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	if len(req.DB) == 0 {
		// todo: 设置默认数据库配置文件
		req.DB = "wgzl,syxx,fmzl"
	}

	ps, err := ic.SimpleSearch(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.PageOK(ps, req.PageSize, req.PageIndex, len(ps), "查询成功")
}
