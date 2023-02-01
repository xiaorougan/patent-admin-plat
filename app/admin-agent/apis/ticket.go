package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/admin-agent/model"
	"go-admin/app/admin-agent/service"
	"go-admin/app/admin-agent/service/dtos"
	"strconv"
)

type Ticket struct {
	api.Api
}

// GetAllTicketPages
// @Summary 获取工单列表
// @Description 获取工单列表
// @Tags 工单
// @Accept  application/json
// @Product application/json
// @Router /api/v1/admin-agent/ticket [get]
// @Param pageIndex query int true "pageIndex"
// @Param pageSize query int true "pageSize"
// @Param query query string true "query"
// @Security Bearer
func (e Ticket) GetAllTicketPages(c *gin.Context) {
	s := service.Ticket{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	t := c.Query("type")
	status := c.Query("status")
	pageIndex, _ := strconv.Atoi(c.Query("pageIndex"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	req := dtos.TicketPagesReq{}
	req.PageIndex = pageIndex
	req.PageSize = pageSize
	req.Type = t
	req.Status = status

	list := make([]model.Ticket, 0)
	var count int64
	if err = s.GetTicketPages(&req, &list, &count); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.PageOK(list, int(count), req.PageSize, req.PageIndex, "查询成功")
}

// CreateTicket
// @Summary 新建工单
// @Description 新建工单
// @Tags 工单
// @Accept  application/json
// @Product application/json
// @Router /api/v1/admin-agent/ticket [post]
// @Security Bearer
func (e Ticket) CreateTicket(c *gin.Context) {
	s := service.Ticket{}
	req := dtos.TicketReq{}
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

	userID := user.GetUserId(c)
	req.UserID = userID

	if err = s.Create(&req); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req, "创建成功")
}

// UpdateTicket
// @Summary 更新工单
// @Description 更新工单
// @Tags 工单
// @Accept  application/json
// @Product application/json
// @Router /api/v1/admin-agent/ticket/{id}/update [post]
// @Security Bearer
func (e Ticket) UpdateTicket(c *gin.Context) {
	s := service.Ticket{}
	req := dtos.TicketReq{}
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

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	userID := user.GetUserId(c)
	req.UserID = userID

	if err = s.Update(id, &req); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req, "更新成功")
}

// CloseTicket
// @Summary 关闭工单
// @Description 关闭工单
// @Tags 工单
// @Accept  application/json
// @Product application/json
// @Router /api/v1/admin-agent/ticket/{id}/close [post]
// @Security Bearer
func (e Ticket) CloseTicket(c *gin.Context) {
	s := service.Ticket{}
	req := dtos.TicketReq{}
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

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	userID := user.GetUserId(c)
	req.UserID = userID

	if err = s.Close(id, &req); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(req, "关闭成功")
}

// RemoveTicket
// @Summary 删除工单
// @Description 删除工单
// @Tags 工单
// @Accept  application/json
// @Product application/json
// @Router /api/v1/admin-agent/ticket/{id} [delete]
// @Security Bearer
func (e Ticket) RemoveTicket(c *gin.Context) {
	s := service.Ticket{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	if err = s.Remove(id); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	e.OK(nil, "删除成功")
}
