package service

import (
	"errors"
	log "github.com/go-admin-team/go-admin-core/logger"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"github.com/google/uuid"
	"go-admin/app/admin-agent/model"
	"go-admin/app/admin-agent/service/dtos"
	"go-admin/app/user-agent/my_config"
	"go-admin/app/user-agent/service/dto"
	"go-admin/app/user-agent/service/report"
	"go-admin/app/user-agent/utils"
	cDto "go-admin/common/dto"
	"gorm.io/gorm"
	"strconv"
	"strings"
	//"gorm.io/gorm"
)

type Report struct {
	service.Service
}

// UserGetRepoById 获取Report对象
func (e *Report) UserGetRepoById(id int, model *model.Report) error {
	//引用传递、函数名、形参、返回值
	var err error
	db := e.Orm.Where("report_Id = ?  ", id).First(model)
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看估值报告不存在或无权查看")
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if db.Error != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// InsertReport 添加Report对象
func (e *Report) InsertReport(c *dtos.ReportGetPageReq, typer string) (error, *model.Report) {

	var err error
	var data model.Report
	var i int64
	err = e.Orm.Model(&data).Where("report_id = ?", c.ReportId).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err, nil
	}
	if i > 0 {
		err := errors.New("报告ID已存在！")
		e.Log.Errorf("db error: %s", err)
		return err, nil
	}

	c.Generate(&data)
	data.CreatedAt = dtos.UpdateTime()
	data.ReportName = typer + ".defaultName." + uuid.New().String()
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err, nil
	}

	return nil, &data
}

func (e *Report) InsertRelation(c *dtos.ReportRelaReq) error {
	var err error
	var data model.ReportRelation
	var i int64

	err = e.Orm.Model(&data).
		Where("report_id = ? and user_id = ? and patent_id = ?", c.ReportId, c.UserId, c.PatentId).
		Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("该关系已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	c.GenerateRela(&data)
	data.CreatedAt = dtos.UpdateTime()

	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}

	return nil
}

// UpdateReport 撤销申请报告
func (e *Report) UpdateReport(c *dtos.ReportGetPageReq) error {
	var err error
	var model model.Report
	db := e.Orm.Where("Report_Id = ? ", c.ReportId).First(&model)
	if err = db.Error; err != nil {
		e.Log.Errorf("Service Update report error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("报告不存在")
	}
	c.Generate(&model)
	model.UpdatedAt = dtos.UpdateTime()
	update := e.Orm.Model(&model).Where("report_id = ?", &model.ReportId).Updates(&model)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update report-info error")
		log.Warnf("db update error")
		return err
	}
	return nil
}

//
//// DeleteReportsRela 删除申请关系
//func (e *Report) DeleteReportsRela(c *dtos.ReportRelaReq) error {
//	var err error
//	var model model.ReportRelation
//	db := e.Orm.Where("Report_Id = ?  ", c.ReportId).First(&model)
//	if err = db.Error; err != nil {
//		e.Log.Errorf("Service Update report error: %s", err)
//		return err
//	}
//	if db.RowsAffected == 0 {
//		return errors.New("报告不存在")
//
//	}
//	c.GenerateRela(&model)
//	model.UpdatedAt = dtos.UpdateTime()
//
//	update := e.Orm.Model(&model).Where("report_id = ?", &model.ReportId).Delete(&model)
//	if err = update.Error; err != nil {
//		e.Log.Errorf("db error: %s", err)
//		return err
//	}
//	if update.RowsAffected == 0 {
//		err = errors.New("update report-info error")
//		log.Warnf("db update error")
//		return err
//	}
//
//	return nil
//}

// GetNovelty 获取查新报告
func (e *Report) GetNovelty(c *dto.NoveltyReportReq) (string, error) {
	query := report.GenQuery(c.KeyWords)

	checkMap := make(map[string]*dto.PatentDetail)
	for i := 0; i < len(query) && len(checkMap) < 30; i++ {
		searchReq := &dto.SimpleSearchReq{
			Pagination: cDto.Pagination{
				PageIndex: 1,
				PageSize:  100,
			},
			Query: query[i],
			DB:    "wgzl,syxx,fmzl",
		}
		checkListTemp, err := GetCurrentInnojoy().Search(searchReq, nil)
		if err != nil {
			return "", err
		}
		for _, res := range checkListTemp {
			if _, ok := checkMap[res.Pnm]; !ok {
				checkMap[res.Pnm] = res
			}
		}
	}

	var checkList []*dto.PatentDetail
	for _, pd := range checkMap {
		checkList = append(checkList, pd)
	}

	var totalinfo []string
	for j := 0; j < len(checkList); j++ {
		totalinfo = append(totalinfo, checkList[j].Ti+"，"+checkList[j].Abst)
	}

	var sentence = c.Title + "。" + c.CL
	segments := report.Seg.Segment([]byte(sentence))
	result := report.GenKey(segments)
	GenVerb := report.GetResult(segments, 2)
	ts2 := report.NewTextSimilarity(strings.Split(sentence, "。"))

	var sims []report.Similarity
	for j := 0; j < len(checkList); j++ {
		var temp = report.Similarity{}
		for i := 0; i < len(result); i++ {
			if strings.Contains(checkList[j].Ti, result[i]) || strings.Contains(checkList[j].Abst, result[i]) {
				temp.Count++
				temp.Words = append(temp.Words, result[i])
			}
		}
		segments1 := report.Seg.Segment([]byte(checkList[j].Ti + checkList[j].Abst))
		resWords1 := report.GetResult(segments1, 0)
		result1 := report.RemoveStop(report.Unique(resWords1))
		temp.Score, _ = ts2.Similarity(result1, result)
		sims = append(sims, temp)
	}
	keywords := ts2.Keywords(-1, 5, 0)
	keywords = report.Unique(keywords)
	var searchList string
	var searchWord = make([][]string, len(keywords))
	for i := 0; i < len(keywords); i++ {
		contain := false
		for j := 0; j < len(GenVerb); j++ {
			if keywords[i] == GenVerb[j] {
				contain = true
				break
			}
		}
		if contain == true {
			continue
		}
		similar := report.GetSimilar(keywords[i])
		searchList += keywords[i] + " " + strings.Join(similar, " ") + "\n"
		searchWord[i] = make([]string, 0)
		searchWord[i] = append(searchWord[i], keywords[i])
		searchWord[i] = append(searchWord[i], similar...)
	}
	queryExpression := report.GenQueryExpression(searchWord)
	n := len(sims)
	var conclusion []string
	var retconc string
	var closeCount = 1 //密切相关的专利数量
	for i := 0; i < n-1 && closeCount < 15; i++ {
		maxNumIndex := i // 无序区第一个
		for j := i + 1; j < n; j++ {
			if sims[j].Score > sims[maxNumIndex].Score {
				maxNumIndex = j
			}
		}
		sims[i], sims[maxNumIndex] = sims[maxNumIndex], sims[i]
		checkList[i], checkList[maxNumIndex] = checkList[maxNumIndex], checkList[i]
		if sims[i].Score > 0.4 {
			if c.Title == checkList[i].Ti {
				continue
			}
			header := report.GenConclusionHeader(closeCount, checkList[i].Pinn, checkList[i].Pa, checkList[i].Ti,
				checkList[i].Pnm, checkList[i].Ad, report.Score2Str(sims[i].Score), checkList[i].Abst)
			conclusion = append(conclusion, header)
			retconc = retconc + "专利" + strconv.Itoa(closeCount) + "是" + checkList[i].Cl + "\n"
			closeCount++
		}
	}
	scale := float64(closeCount-1) / float64(len(checkList))
	if scale > 0.5 {
		retconc = retconc + "而本专利是" + c.CL + report.DisclaimerTemplate + report.NegativeResult
	} else {
		retconc = retconc + "本专利是" + c.CL + report.DisclaimerTemplate + report.PositiveResult
	}
	realteCount := 0
	if len(checkList) > 0 {
		realteCount = len(checkList) - 1
	} else {
		realteCount = len(checkList)
	}
	reportBase := report.NewNoveltyTemplate()
	reportBase.Replace("$NUMBER", uuid.New().String()).
		Replace("$DEPART_NAME", my_config.CurrentPatentConfig.NoveltyReportConfig.DepartName).
		Replace("$CONTACT_ADDR", my_config.CurrentPatentConfig.NoveltyReportConfig.ContactAddr).
		Replace("$ZIP_CODE", my_config.CurrentPatentConfig.NoveltyReportConfig.ZipCode).
		Replace("$MANAGER_NAME", my_config.CurrentPatentConfig.NoveltyReportConfig.ManagerName).
		Replace("$MANAGER_TEL", my_config.CurrentPatentConfig.NoveltyReportConfig.ManagerTel).
		Replace("$CONTACT_NAME", my_config.CurrentPatentConfig.NoveltyReportConfig.ContactName).
		Replace("$CONTACT_TEL", my_config.CurrentPatentConfig.NoveltyReportConfig.ContactTel).
		Replace("$EMAIL", my_config.CurrentPatentConfig.NoveltyReportConfig.Email).
		Replace("$DATABASE", my_config.CurrentPatentConfig.NoveltyReportConfig.DataBase).
		Replace("$PATENT_NAME", c.Title).
		Replace("$USER_NAME", c.Applicant).
		Replace("$Institution", c.Org).
		Replace("$FINISH_DATE", utils.FormatCurrentTime()).
		Replace("$TECH_POINT", c.CL).
		Replace("$QUERY_WORD", report.ToHtml(searchList)).
		Replace("$QUERY_EXPRESSION", report.ToHtml(queryExpression)).
		Replace("$RELATIVE_NUM", strconv.Itoa(realteCount)).
		Replace("$VERY_RELATIVE_NUM", strconv.Itoa(closeCount-1)).
		Replace("$SEARCH_RESULT", report.ToHtml(strings.Join(conclusion, "\n"))).
		Replace("$CONCLUSION", report.ToHtml(retconc))

	return reportBase.String(), nil
}
