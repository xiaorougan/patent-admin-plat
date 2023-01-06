package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"github.com/prometheus/common/log"
	"go-admin/app/admin-agent/service/dtos"
	"go-admin/app/user-agent/models"
	"go-admin/app/user-agent/service/dto"
	"gorm.io/gorm"
	"sort"
	"strconv"
)

type Patent struct {
	service.Service
}

// GetPage 获取Patent列表
func (e *Patent) GetPage(c *dto.PatentReq, list *[]models.Patent, count *int64) error {
	var err error
	var data models.Patent
	err = e.Orm.Model(&data).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GetPageByIds 通过Id数组获取Patent对象列表
func (e *Patent) GetPageByIds(d *dto.PatentsIds, list *[]models.Patent, count *int64) error {
	var err error
	var ids []int = d.GetPatentId()
	for i := 0; i < len(ids); i++ {
		if ids[i] != 0 {
			var data1 models.Patent
			err = e.Orm.Model(&data1).
				Where("Patent_Id = ? ", ids[i]).
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

// GetInventorPageByIds 通过Id数组获取Patent对象列表
func (e *Patent) GetInventorPageByIds(d *dto.PatentsIds, list *[]models.Patent, count *int64) error {
	var err error
	var ids []int = d.GetPatentId()
	for i := 0; i < len(ids); i++ {
		if ids[i] != 0 {
			var data1 models.Patent
			err = e.Orm.Model(&data1).
				Where("Patent_Id = ? ", ids[i]).
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

// Get 获取Patent对象
func (e *Patent) Get(d *dto.PatentById, model *models.Patent) error {
	var err error
	db := e.Orm.First(model, d.PatentId)
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看专利不存在或无权查看")
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if db.Error != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GeByPNM 获取Patent对象
func (e *Patent) GeByPNM(d *dto.PatentBriefInfo, model *models.Patent) error {
	//引用传递、函数名、形参、返回值
	var err error
	db := e.Orm.First(model, d.PNM)
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看专利不存在或无权查看")
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if db.Error != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// GeById 获取Patent对象
func (e *Patent) GeById(d *dtos.ReportRelaReq, model *models.Patent) error {
	var err error
	db := e.Orm.Where("patent_id = ?", d.PatentId).First(model)
	err = db.Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看专利不存在或无权查看")
		e.Log.Errorf("db error:%s", err)
		return err
	}
	if db.Error != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 添加Patent对象
func (e *Patent) Insert(c *dto.PatentReq) error {
	var err error
	var data models.Patent
	var i int64
	err = e.Orm.Model(&data).Where("PNM = ?", c.PNM).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if i > 0 {
		err := errors.New("专利ID已存在！")
		e.Log.Errorf("db error: %s", err)
		return err
	}
	c.GenerateList(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	return nil
}

// InsertIfAbsent 返回值与上面不同
func (e *Patent) InsertIfAbsent(c *dto.PatentReq) (*models.Patent, error) {
	var err error
	var data models.Patent
	var i int64
	err = e.Orm.Model(&data).Where("PNM = ? OR patent_id = ?", c.PNM, c.PatentId).Count(&i).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}
	if i > 0 {
		err = e.Orm.Model(&data).Where("PNM = ?", c.PNM).First(&data).Error
		if err != nil {
			e.Log.Errorf("db error: %s", err)
			return nil, err
		}
		return &data, nil
	}
	c.GenerateList(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("db error: %s", err)
		return nil, err
	}
	return &data, nil
}

// UpdateLists 根据PatentId修改Patent对象
func (e *Patent) UpdateLists(c *dto.PatentReq) error {
	var err error
	var model models.Patent
	db := e.Orm.First(&model, c.PatentId)
	if err = db.Error; err != nil {
		e.Log.Errorf("Service Update Patent error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("专利不存在")
	}
	c.GenerateList(&model)
	update := e.Orm.Model(&model).Where("patent_id = ?", &model.PatentId).Updates(&model)
	if err = update.Error; err != nil {
		e.Log.Errorf("db error: %s", err)
		return err
	}
	if update.RowsAffected == 0 {
		err = errors.New("update patent-info error")
		log.Warnf("db update error")
		return err
	}
	return nil
}

// Remove 根据id删除Patent
func (e *Patent) Remove(c *dto.PatentById) error {
	var err error
	var data models.Patent

	db := e.Orm.Delete(&data, c.PatentId)
	if db.Error != nil {
		err = db.Error
		e.Log.Errorf("Delete error: %s", err)
		return err
	}
	if db.RowsAffected == 0 {
		err = errors.New("删除数据不存在")
		return err
	}
	return nil
}

// GetGraphByPatents 通过Patent数组获得专利发明人的关系图
func (e *Patent) GetGraphByPatents(ListPatents []models.Patent, Inventorgraph *models.Graph) error {
	var err error
	NodeList := make([]models.Node, 0)                                                       //return node
	LinkList := make([]models.Link, 0)                                                       //return link
	MaxInventors := 200                                                                      //check max 200 Inventors in the ListPatents
	PreNodeList := make([]models.PreNode, 0)                                                 //Similar to node struct and need change some attributes to become NodeList
	PreLinkList := make([]models.PreLink, 0)                                                 //Similar to link struct and need change some attributes to become LinkList
	Inventors, Relations := FindInventorsAndRelationsFromPatents(ListPatents, &MaxInventors) //relations is an Upper Triangle
	if len(Inventors) == 0 {
		return err
	}
	StrongRelationInventors := MinResult(len(Inventors), 10) //chose the top10%(maximum is 10) inventors as StrongRelationInventors(must show)
	for i, inventor := range Inventors[0:StrongRelationInventors] {
		PreNodeList = append(PreNodeList, models.PreNode{NodeId: inventor.Id, NodeCategory: i})
	}
	//build top10% to top10% relationship
	MaxNumberOfLinks := 15 //the maximum find link
	MaxExpansion := 4      //every Node can extend the number of nodes
	NowLink := FindRelationFrequencyAndSort(Relations, Inventors[0:StrongRelationInventors], Inventors[0:StrongRelationInventors], MaxNumberOfLinks, MaxExpansion, MaxInventors)
	PreLinkList = append(PreLinkList, NowLink...)
	//build top10% to others relationship
	MaxNumberOfLinks = 200
	MaxExpansion = 5
	NowLink = FindRelationFrequencyAndSort(Relations, Inventors[0:StrongRelationInventors], Inventors[StrongRelationInventors:], MaxNumberOfLinks, MaxExpansion, MaxInventors)
	PreLinkList = append(PreLinkList, NowLink...)
	//add the extended nodes
	for _, i := range NowLink {
		if !Inventors[i.Target].InTheGraph {
			PreNodeList = append(PreNodeList, models.PreNode{NodeId: i.Target, NodeCategory: PreNodeList[i.Source].NodeCategory})
			Inventors[i.Target].InTheGraph = true
		}
	}
	//build others to others relationship
	SourceNode := make([]models.Inventor, 0)
	TargetNode := make([]models.Inventor, 0)
	for _, node := range PreNodeList[StrongRelationInventors:] {
		SourceNode = append(SourceNode, models.Inventor{Id: node.NodeId})
		TargetNode = append(SourceNode, models.Inventor{Id: node.NodeId})
	}
	MaxNumberOfLinks = 100
	MaxExpansion = 3
	NowLink = FindRelationFrequencyAndSort(Relations, SourceNode, TargetNode, MaxNumberOfLinks, MaxExpansion, MaxInventors)
	PreLinkList = append(PreLinkList, NowLink...)
	//deal the struct PreLink and PreNode
	MaxSizeofNode := 50
	for _, node := range PreNodeList {
		NodeList = append(NodeList, models.Node{
			NodeId:            strconv.FormatInt(int64(node.NodeId), 10),
			NodeName:          Inventors[node.NodeId].Name,
			NodeValue:         Inventors[node.NodeId].TheNumberOfPatents,
			NodeSymbolizeSize: float32((Inventors[node.NodeId].TheNumberOfPatents) * MaxSizeofNode / Inventors[0].TheNumberOfPatents),
			NodeCategory:      node.NodeCategory,
		})
	}
	for _, link := range PreLinkList {
		LinkList = append(LinkList, models.Link{
			Source: strconv.FormatInt(int64(link.Source), 10),
			Target: strconv.FormatInt(int64(link.Target), 10),
			Value:  link.Value,
		})
	}
	Inventorgraph.Links = LinkList
	Inventorgraph.Nodes = NodeList
	return nil
}

// FindRelationFrequencyAndSort 建立边关系
func FindRelationFrequencyAndSort(relations []int, sources []models.Inventor, targets []models.Inventor, MaxNumberOfLinks int, MaxExtend int, MaxInventors int) []models.PreLink {
	LinkList := make([]models.PreLink, 0)
	LinkSum := 0
	LinkExtend := make(map[int]int)
	LinkReturnList := make([]models.PreLink, 0)
	LinkSearch := make(map[int]bool) //avoid duplicate link
	//init LinkExtend
	for _, i2 := range sources {
		LinkExtend[i2.Id] = 0
	}
	//find all Relation between sources and targets
	for _, source := range sources {
		for _, target := range targets {
			if source.Id == target.Id || relations[target.Id*MaxInventors+source.Id] == 0 && relations[source.Id*MaxInventors+target.Id] == 0 {
				continue
			}
			if _, ok := LinkSearch[source.Id*MaxInventors+target.Id]; ok {
				continue
			}
			LinkSearch[source.Id*MaxInventors+target.Id] = true
			if source.Id < target.Id {
				LinkList = append(LinkList, models.PreLink{Source: source.Id, Target: target.Id, Value: relations[source.Id*MaxInventors+target.Id]})
			} else {
				LinkList = append(LinkList, models.PreLink{Source: source.Id, Target: target.Id, Value: relations[target.Id*MaxInventors+source.Id]})
			}
		}
	}
	//sort LinkList
	sort.Slice(LinkList, func(i, j int) bool {
		if LinkList[i].Value > LinkList[j].Value {
			return true
		}
		return false
	})
	//pick return links condition(MaxExtend,MaxNumberOfLinks)
	MaxNumberOfLinks = MinResult(MaxNumberOfLinks, len(LinkList))
	for _, link := range LinkList {
		if LinkSum >= MaxNumberOfLinks {
			break
		}
		if LinkExtend[link.Source] >= MaxExtend {
			continue
		}
		LinkSum++
		LinkExtend[link.Source]++
		LinkReturnList = append(LinkReturnList, link)
	}
	return LinkReturnList
}

// FindInventorsAndRelationsFromPatents --------------------------------------------------------------------------
// 通过patents数组查找patents数组中的发明人以及专利和发明人的关系
func FindInventorsAndRelationsFromPatents(listpatents []models.Patent, n *int) ([]models.Inventor, []int) {

	ListPreInventors := make([]models.PreInventor, 0)
	//find every patents' Inventor and count
	for z := 0; z < len(listpatents); z++ {
		words := make([]string, 0)
		for i := 0; i < len(listpatents[z].PatentProperties); i++ {
			if listpatents[z].PatentProperties[i] == '"' && listpatents[z].PatentProperties[i+1] == 'P' && listpatents[z].PatentProperties[i+2] == 'I' && listpatents[z].PatentProperties[i-1] == ',' && listpatents[z].PatentProperties[i+8] != '"' {
				now := i + 8
				for j := i + 8; j < len(listpatents[z].PatentProperties); j++ {
					if listpatents[z].PatentProperties[j] == '"' {
						words = append(words, listpatents[z].PatentProperties[now:j])
						break
					}
					if listpatents[z].PatentProperties[j] == ';' {
						words = append(words, listpatents[z].PatentProperties[now:j])
						now = j + 1
					}
				}
				break
			}
		}
		for i := 0; i < len(words); i++ {
			InventorExist := false
			for j := 0; j < len(ListPreInventors); j++ {
				if ListPreInventors[j].Name == words[i] {
					ListPreInventors[j].TheNumberOfPatents++
					InventorExist = true
					ListPreInventors[j].PatentsId = append(ListPreInventors[j].PatentsId, listpatents[z].PatentId)
					break
				}
			}
			if !InventorExist {
				NewPatents := make([]int, 0)
				ListPreInventors = append(ListPreInventors, models.PreInventor{
					Name:               words[i],
					TheNumberOfPatents: 1,
					PatentsId:          append(NewPatents, listpatents[z].PatentId),
				})
			}
		}
	}
	//sort Inventors
	sort.Slice(ListPreInventors, func(i, j int) bool {
		if ListPreInventors[i].TheNumberOfPatents > ListPreInventors[j].TheNumberOfPatents {
			return true
		}
		return false
	})

	//write the id to Inventors
	for i := 0; i < len(ListPreInventors); i++ {
		ListPreInventors[i].Id = i
	}
	//create Relations
	*n = MinResult(*n, len(ListPreInventors))
	ListRelations := make([]int, (*n)*(*n))
	for i := 0; i < *n; i++ {
		for j := i; j < *n; j++ {
			var count int
			source := make(map[int]bool)
			for _, OneOfPatentId := range ListPreInventors[i].PatentsId {
				source[OneOfPatentId] = true

			}
			for _, OneOfPatentId := range ListPreInventors[j].PatentsId {
				if _, ok := source[OneOfPatentId]; ok {
					count++
				}

			}
			ListRelations[i*(*n)+j] = count
		}
	}

	//change preInventors to Inventors(delete preInventor->patents)
	ListInventors := make([]models.Inventor, 0)
	for _, i := range ListPreInventors {
		NowInventor := models.Inventor{Id: i.Id, Name: i.Name, TheNumberOfPatents: i.TheNumberOfPatents}
		ListInventors = append(ListInventors, NowInventor)
	}
	return ListInventors, ListRelations
}

// MinResult --------------------------------------------------------------------------------------------------------------------
func MinResult(a1 int, a2 int) int {
	if a1 >= a2 {
		return a2
	} else {
		return a1
	}
}
