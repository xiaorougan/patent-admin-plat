package service

import (
	"errors"
	"github.com/go-admin-team/go-admin-core/sdk/service"
	"github.com/prometheus/common/log"
	"go-admin/app/admin-agent/service/dtos"
	amodels "go-admin/app/admin/models"
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
func (e *Patent) GetGraphByPatents(listp []models.Patent, Inventorgraph *models.Graph) error {
	var err error
	links := make([]models.Link, 0)
	listInventorId := make(map[string]int)
	listip := make([]models.InventorPatent, 0)
	err = FindTheInventorFromPatents(&listInventorId, &listip, listp)
	if err != nil {
		e.Log.Errorf("Can`t get Inventor", err)
		return err
	}
	usertimes := make(map[int]int)
	for i := 0; i < len(listip); i++ {
		if usertimes[listip[i].InventorId] == 0 {
			usertimes[listip[i].InventorId] = 1
		} else {
			usertimes[listip[i].InventorId]++
		}

	}
	usertimesordered := rankByWordCount(usertimes)
	var members int
	if len(usertimesordered) < 500 {
		members = len(usertimesordered)
	} else {
		members = 500
	}
	UserIsNode := make([]bool, members)
	var StrongRelationNode int
	if members >= 100 {
		StrongRelationNode = 10
	} else {
		StrongRelationNode = members / 10
	}
	NodeList := make([]models.Node, StrongRelationNode)
	for i := 0; i < StrongRelationNode; i++ {
		NodeList[i].NodeCategory = i
		NodeList[i].NodeId = strconv.FormatInt(int64(usertimesordered[i].Key), 10)
		NodeList[i].NodeValue = usertimesordered[i].Value
		UserIsNode[i] = true
	}
	userspatents := make([]models.OneUserPatents, members)
	for i := 0; i < len(listip); i++ {
		for j := 0; j < members; j++ {
			if listip[i].InventorId == usertimesordered[j].Key {
				userspatents[j].Patentsid = append(userspatents[j].Patentsid, listip[i].PatentId)
				break
			}
		}
	}
	useruserrelationfirst10 := make(map[int]int)
	first10 := StrongRelationNode
	firstlinks := first10
	for i := 0; i < first10; i++ {
		for j := i + 1; j < first10; j++ {
			RelationExist := 0
			for z := 0; z < len(userspatents[i].Patentsid); z++ {
				for z1 := 0; z1 < len(userspatents[j].Patentsid); z1++ {
					if userspatents[i].Patentsid[z] == userspatents[j].Patentsid[z1] {
						RelationExist++
						break
					}
				}
			}
			if RelationExist != 0 {
				useruserrelationfirst10[i*10+j] = RelationExist
			}
		}
	}
	useruserrelationfirst10ordered := rankByWordCount(useruserrelationfirst10) //给边排序
	for i := 0; i < minresult(firstlinks, len(useruserrelationfirst10ordered)); i++ {
		var nowlink models.Link
		nowlink.Source = strconv.FormatInt(int64(usertimesordered[useruserrelationfirst10ordered[i].Key/10].Key), 10)
		nowlink.Target = strconv.FormatInt(int64(usertimesordered[useruserrelationfirst10ordered[i].Key%10].Key), 10)
		nowlink.Value = useruserrelationfirst10ordered[i].Value
		links = append(links, nowlink)
	}

	if len(NodeList) < 1 {
		err = errors.New("删除数据不存在")
		return err
	}

	useruserrelationfirst10ToOthers := make(map[int]int)
	secondLinks := 200
	ExtendNodeTime := make([]int, members)

	for i := 0; i < first10; i++ {
		for j := first10; j < members; j++ {
			RelationExist := 0
			for z := 0; z < len(userspatents[i].Patentsid); z++ {
				for z1 := 0; z1 < len(userspatents[j].Patentsid); z1++ {
					if userspatents[i].Patentsid[z] == userspatents[j].Patentsid[z1] {
						RelationExist++
						break
					}
				}
			}
			if RelationExist != 0 {
				useruserrelationfirst10ToOthers[i*500+j] = RelationExist
			}
		}
	}
	useruserrelationfirst10ToOthersordered := rankByWordCount(useruserrelationfirst10ToOthers)
	NodelistIdToTimeList := make(map[string]int)

	for i := 0; i < minresult(secondLinks, len(useruserrelationfirst10ToOthersordered)); i++ {
		source := useruserrelationfirst10ToOthersordered[i].Key / 500
		target := useruserrelationfirst10ToOthersordered[i].Key % 500
		if ExtendNodeTime[source] >= 5 {
			continue
		} else {
			if UserIsNode[target] == false {
				UserIsNode[target] = true
				ExtendNodeTime[source]++
				var nowlink models.Link
				var nowNode models.Node
				nowNode.NodeCategory = NodeList[source].NodeCategory
				nowNode.NodeId = strconv.FormatInt(int64(usertimesordered[target].Key), 10)
				NodelistIdToTimeList[nowNode.NodeId] = target
				nowlink.Source = strconv.FormatInt(int64(usertimesordered[source].Key), 10)
				nowlink.Target = strconv.FormatInt(int64(usertimesordered[target].Key), 10)
				nowlink.Value = useruserrelationfirst10ToOthersordered[i].Value
				links = append(links, nowlink)
				NodeList = append(NodeList, nowNode)
			}
		}
	}

	useruserrelationOthersToOthers := make(map[int]int)

	for i := first10; i < len(NodeList); i++ {
		iToUserspatentsPosition := NodelistIdToTimeList[NodeList[i].NodeId]
		for j := i + 1; j < len(NodeList); j++ {
			RelationExist := 0
			jToUserspatentsPosition := NodelistIdToTimeList[NodeList[j].NodeId]
			for z := 0; z < len(userspatents[iToUserspatentsPosition].Patentsid); z++ {
				for z1 := 0; z1 < len(userspatents[jToUserspatentsPosition].Patentsid); z1++ {

					if userspatents[iToUserspatentsPosition].Patentsid[z] == userspatents[jToUserspatentsPosition].Patentsid[z1] {
						RelationExist++
						break
					}
				}
			}
			if RelationExist != 0 {
				useruserrelationOthersToOthers[i*500+j] = RelationExist
			}

		}
	}
	useruserrelationOthersToOthersOrdered := rankByWordCount(useruserrelationOthersToOthers)
	thirdlinks := 50
	for i := 0; i < minresult(thirdlinks, len(useruserrelationOthersToOthersOrdered)); i++ {
		source := useruserrelationOthersToOthersOrdered[i].Key / 500
		target := useruserrelationOthersToOthersOrdered[i].Key % 500

		if ExtendNodeTime[source] >= 2 {
			continue
		} else {
			ExtendNodeTime[source]++
			var nowlink models.Link
			nowlink.Source = NodeList[source].NodeId
			nowlink.Target = NodeList[target].NodeId
			nowlink.Value = useruserrelationOthersToOthersOrdered[i].Value
			links = append(links, nowlink)
		}
	}

	listu := make([]amodels.SysUser, 0)

	for i := 0; i < len(NodeList); i++ {
		var user1 amodels.SysUser
		user1.UserId, err = strconv.Atoi(NodeList[i].NodeId)
		for k, v := range listInventorId {
			if v == user1.UserId {
				user1.Username = k
			}
		}
		listu = append(listu, user1)
	}
	//err = e.MakeContext(c).
	//	MakeOrm().
	//	MakeService(&gservice.Service).
	//	Errors
	max := 0
	min := 100000
	NodeList[0].NodeValue = usertimesordered[0].Value
	NodeList[0].NodeSymbolizeSize = 50
	NodeList[0].NodeName = listu[0].Username
	for i := 1; i < first10; i++ {
		NodeList[i].NodeValue = usertimesordered[i].Value
		if NodeList[i].NodeValue > max {
			max = NodeList[i].NodeValue
		}
		if NodeList[i].NodeValue < min {
			min = NodeList[i].NodeValue
		}
	}
	for i := first10; i < len(NodeList); i++ {
		NodeList[i].NodeValue = usertimesordered[NodelistIdToTimeList[NodeList[i].NodeId]].Value
		if NodeList[i].NodeValue > max {
			max = NodeList[i].NodeValue
		}
		if NodeList[i].NodeValue < min {
			min = NodeList[i].NodeValue
		}
	}

	for i := 1; i < len(NodeList); i++ {
		NodeList[i].NodeSymbolizeSize = float32(float32(NodeList[i].NodeValue*30) / float32(maxresult((max), 1)))
		NodeList[i].NodeName = listu[i].Username
	}
	Inventorgraph.Links = links
	Inventorgraph.Nodes = NodeList
	return nil
}

// FindTheInventorFromPatents --------------------------------------------------------------------------
// 通过patents数组查找patents数组中的发明人以及专利和发明人的关系
func FindTheInventorFromPatents(listInventorId *map[string]int, listup2 *[]models.InventorPatent, listp []models.Patent) error {
	var err error
	count := 0
	for z := 0; z < len(listp); z++ {
		words := make([]string, 0)
		for i := 0; i < len(listp[z].PatentProperties); i++ {
			if listp[z].PatentProperties[i] == '"' && listp[z].PatentProperties[i+1] == 'P' && listp[z].PatentProperties[i+2] == 'I' && listp[z].PatentProperties[i-1] == ',' && listp[z].PatentProperties[i+8] != '"' {
				now := i + 8
				for j := i + 8; j < len(listp[z].PatentProperties); j++ {
					if listp[z].PatentProperties[j] == '"' {
						words = append(words, listp[z].PatentProperties[now:j])
						break
					}
					if listp[z].PatentProperties[j] == ';' {
						words = append(words, listp[z].PatentProperties[now:j])
						now = j + 1
					}
				}
				break
			}
		}
		for i := 0; i < len(words); i++ {
			_, ok := (*listInventorId)[words[i]]
			if !ok {
				(*listInventorId)[words[i]] = count
				count++
			}
			var inventorpatent models.InventorPatent
			inventorpatent.PatentId = listp[z].PatentId
			inventorpatent.InventorId = (*listInventorId)[words[i]]
			*listup2 = append(*listup2, inventorpatent)
		}
	}
	return err
}

// --------------------------------------------------------------------------------------------------------------------
// map按value的值排序
func rankByWordCount(wordFrequencies map[int]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	//从小到大排序
	//sort.Sort(pl)
	//从大到小排序
	sort.Sort(sort.Reverse(pl))
	return pl
}

type Pair struct {
	Key   int
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func minresult(a1 int, a2 int) int {
	if a1 >= a2 {
		return a2
	} else {
		return a1
	}
}
func maxresult(a1 int, a2 int) int {
	if a1 >= a2 {
		return a1
	} else {
		return a2
	}
}
