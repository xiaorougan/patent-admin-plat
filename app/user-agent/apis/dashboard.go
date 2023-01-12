package apis

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service"
	"go-admin/app/user-agent/service/dto"
)

type Dashboard struct {
	api.Api
}

// GetDashboard
// @Summary Dashboard主页
// @Description  Dashboard主页
// @Tags Dashboard主页
// @Success 200 {object} dto.Dashboard
// @Router /apis/v1/user-agent/dashboard [get]
// @Security Bearer
func (e Dashboard) GetDashboard(c *gin.Context) {
	ups := service.UserPatent{}
	ps := service.Patent{}
	pks := service.Package{}
	rps := service.Report{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&ups.Service).
		MakeService(&ps.Service).
		MakeService(&pks.Service).
		MakeService(&rps.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req := dto.NewEmptyClaim()
	req.UserId = user.GetUserId(c)

	var focusCount int64
	if err = ups.GetFocusCount(req, &focusCount); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	claimList := make([]models.UserPatent, 0)
	var claimCount int64
	if err = ups.GetClaimLists(req, &claimList, &claimCount); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	ids := make([]int, 0, len(claimList))
	for _, claim := range claimList {
		ids = append(ids, claim.PatentId)
	}
	var patentCount int64
	patents, err := ps.GetPageByIds(ids, &patentCount)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	patentStatus := make(map[string]int)
	publicationDates := dto.NewPublicationDates()
	var totalPrice int
	for _, p := range patents {
		pd := dto.PatentDetail{}
		if err = json.Unmarshal([]byte(p.PatentProperties), &pd); err != nil {
			e.Logger.Error(err)
			e.Error(500, err, err.Error())
			return
		}
		if len(pd.Cls) == 0 {
			continue
		}
		if _, ok := patentStatus[pd.Cls]; ok {
			patentStatus[pd.Cls]++
		} else {
			patentStatus[pd.Cls] = 1
		}

		if len(pd.Ad) != 0 {
			publicationDates.AddYear(pd.Ad)
		}

		totalPrice += pd.Idx * dto.PatentPriceBase
	}

	pkgList := make([]models.Package, 0)
	listPkgReq := dto.PackageListReq{UserId: user.GetUserId(c)}
	err = pks.ListByUserId(&listPkgReq, &pkgList)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	pkgCount := len(pkgList)

	var reportCount int64
	err = rps.GetCountByUserID(user.GetUserId(c), &reportCount)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	res := &dto.Dashboard{
		PatentClaimCount:     int(claimCount),
		PatentFocusCount:     int(focusCount),
		PatentStatus:         patentStatus,
		PublicationDates:     publicationDates.List(),
		PackageCount:         pkgCount,
		PatentRecommendation: nil,
		ReportCount:          int(reportCount),
		PatentTotalPrice:     totalPrice,
	}

	e.OK(res, "查询成功")
}
