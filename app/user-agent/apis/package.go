package apis

import (
	"github.com/gin-gonic/gin/binding"
	amodels "go-admin/app/admin/models"
	aservice "go-admin/app/admin/service"
	adto "go-admin/app/admin/service/dto"
	"go-admin/app/user-agent/models"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"
	"go-admin/app/user-agent/service"
	"go-admin/app/user-agent/service/dto"
)

type Package struct {
	api.Api
}

//// GetPage
//// @Summary 列表专利包信息数据
//// @Description 获取JSON
//// @Tags 专利包
//// @Param packageName query string false "packageName"
//// @Router /api/v1/package [get]
//// @Security Bearer
//func (e Package) GetPage(c *gin.Context) {
//	s := service.Package{}
//	req := dtos.PackageGetPageReq{}
//	err := e.MakeContext(c).
//		MakeOrm().
//		Bind(&req).
//		MakeService(&s.Service).
//		Errors
//	if err != nil {
//		e.Logger.Error(err)
//		e.Error(500, err, err.Error())
//		return
//	}
//
//	//数据权限检查
//	//p := actions.GetPermissionFromContext(c)
//
//	list := make([]models.Package, 0)
//	var count int64
//
//	err = s.GetPage(&req, &list, &count)
//	if err != nil {
//		e.Error(500, err, "查询失败")
//		return
//	}
//
//	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
//}

// ListByCurrentUser
// @Summary 获取当前用户专利包列表
// @Description 获取JSON
// @Tags 专利包
// @Router /api/v1/user-agent/package [get]
// @Security Bearer
func (e Package) ListByCurrentUser(c *gin.Context) {
	s := service.Package{}
	req := dto.PackageListReq{}
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

	req.UserId = user.GetUserId(c)

	list := make([]models.Package, 0)

	err = s.ListByUserId(&req, &list)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(list, "查询成功")
}

// Get
// @Summary 获取专利包
// @Description 获取JSON
// @Tags 专利包
// @Param packageId path int true "package_id"
// @Router /api/v1/user-agent/package/{package_id} [get]
// @Security Bearer
func (e Package) Get(c *gin.Context) {
	s := service.Package{}
	req := dto.PackageById{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req, nil).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.Package
	//数据权限检查
	//p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, &object)
	if err != nil {
		e.Error(http.StatusUnprocessableEntity, err, "查询失败")
		return
	}
	e.OK(object, "查询成功")
}

// Insert
// @Summary 创建专利包
// @Description 获取JSON
// @Tags 专利包
// @Accept  application/json
// @Product application/json
// @Param data body dto.PackageInsertReq true "专利包数据"
// @Router /api/v1/user-agent/package [post]
// @Security Bearer
func (e Package) Insert(c *gin.Context) {
	s := service.Package{}
	req := dto.PackageInsertReq{}
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

// Update
// @Summary 修改专利包数据
// @Description 获取JSON
// @Tags 专利包
// @Accept  application/json
// @Product application/json
// @Param data body dto.PackageUpdateReq true "body"
// @Router /api/v1/user-agent/package/{package_id} [put]
// @Security Bearer
func (e Package) Update(c *gin.Context) {
	s := service.Package{}
	req := dto.PackageUpdateReq{}
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

	if pid, err := strconv.Atoi(c.Param("id")); err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	} else {
		req.PackageId = pid
	}

	req.SetUpdateBy(user.GetUserId(c))

	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	err = s.Update(&req)
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(nil, "更新成功")
}

// Delete
// @Summary 删除专利包
// @Description 删除专利包
// @Tags 专利包
// @Param packageId path int true "packageId"
// @Router /api/v1/user-agent/package/{package_id} [delete]
// @Security Bearer
func (e Package) Delete(c *gin.Context) {
	s := service.Package{}
	req := dto.PackageById{}
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.Id, err = strconv.Atoi(c.Param("id"))
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

//----------------------------------------patent-package---------------------------------------

// todo: please modify the swagger comment

// GetPackagePatents
// @Summary 获取指定专利包中的专利列表
// @Description 获取指定专利包中的专利列表
// @Tags 专利包
// @Param packageId path int true "packageId"
// @Router /api/v1/user-agent/package/{package_id}/patent [get]
// @Security Bearer
func (e Package) GetPackagePatents(c *gin.Context) {

	s := service.PatentPackage{}
	s1 := service.Patent{}
	req := dto.PackagePageGetReq{}
	req1 := dto.PatentsIds{}

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

	req.PackageId, err = strconv.Atoi(c.Param("id"))

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	//数据权限检查
	//p := actions.GetPermissionFromContext(c)

	list := make([]models.PatentPackage, 0)
	list1 := make([]models.Patent, 0)
	var count int64

	err = s.GetPatentIdByPackageId(&req, &list, &count)

	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	var count2 int64

	err = e.MakeContext(c).
		MakeOrm().
		Bind(&req1).
		MakeService(&s1.Service).
		Errors

	req1.PatentIds = make([]int, len(list))

	for i := 0; i < len(list); i++ {
		req1.PatentIds[i] = list[i].PatentId
	}

	err = s1.GetPageByIds(&req1, &list1, &count2)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}
	e.PageOK(list1, int(count2), req.GetPageIndex(), req.GetPageSize(), "查询成功")

}

// IsPatentInPackage
// @Summary 查询专利是否已在专利包中
// @Description 查询专利是否已在专利包中
// @Tags 专利包
// @Param packageId path int true "packageId"
// @Router /api/v1/user-agent/package/{package_id}/patent/{patent_id}/isExist [get]
// @Security Bearer
func (e Package) IsPatentInPackage(c *gin.Context) {
	var err error

	pps := service.PatentPackage{}
	req := dto.PatentPackageReq{}

	req.PNM = c.Param("PNM")

	req.PackageId, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.CreateBy = user.GetUserId(c)

	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&pps.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	existed, err := pps.IsPatentInPackage(&req)
	if err != nil {
		e.Error(500, err, "查询失败")
		return
	}

	e.OK(&dto.IsPatentInPackageResp{Existed: existed}, "查询成功")
}

// InsertPackagePatent
// @Summary 将专利加入专利包
// @Description  将专利加入专利包
// @Tags 专利包
// @Accept  application/json
// @Product application/json
// @Param data body dto.PatentReq true "专利表数据"
// @Router /api/v1/user-agent/package/{package_id}/patent/{patent_id} [post]
// @Security Bearer
func (e Package) InsertPackagePatent(c *gin.Context) {
	var err error
	pps := service.PatentPackage{}
	req := dto.PatentPackageReq{}

	ps := service.Patent{}
	patentReq := dto.PatentReq{}
	err = e.MakeContext(c).
		MakeOrm().
		Bind(&patentReq).
		MakeService(&ps.Service).
		Errors
	patentReq.CreateBy = user.GetUserId(c)
	p, err := ps.InsertIfAbsent(&patentReq)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.PatentId = p.PatentId
	req.PNM = p.PNM
	req.PackageId, err = strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	req.CreateBy = user.GetUserId(c)

	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&pps.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	err = pps.InsertPatentPackage(&req)
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	e.OK(nil, "创建成功")
}

// DeletePackagePatent
// @Summary 删除专利包专利
// @Description  删除专利包专利
// @Tags 专利包
// @Param PatentId query string false "专利ID"
// @Param PackageId query string false "专利包ID"
// @Router /api/v1/user-agent/package/{package_id}/patent/{patent_id} [delete]
// @Security Bearer
func (e Package) DeletePackagePatent(c *gin.Context) {
	s := service.PatentPackage{}
	req := dto.PackagePageGetReq{}
	req.SetUpdateBy(user.GetUserId(c))
	err := e.MakeContext(c).
		MakeOrm().
		MakeService(&s.Service).
		Errors

	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	packageId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.PackageId = packageId

	patentId, err := strconv.Atoi(c.Param("pid"))
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.PatentId = patentId

	err = s.RemovePackagePatent(&req)

	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(req.PackageBack, "删除成功")
}

//---------------------------------------------------patent--graph-------------------------------------------------------

// GetTheGraphByPackageId3
// @Summary 获取专利包中专利的发明人的关系
// @Description  获取专利包中专利的发明人的关系
// @Tags 专利表
// @Router /api/v1/user-agent/package/{packageId}/relationship3 [get]
// @Security Bearer
func (e Package) GetTheGraphByPackageId3(c *gin.Context) {
	sup := service.UserPatent{}
	su := aservice.SysUser{}
	sp := service.Patent{}
	//gservice := service.Node{}
	reqp := dto.PatentsIds{}
	requ := adto.SysUserById{}
	spp := service.PatentPackage{}
	reqpp := dto.PackagePageGetReq{}
	Inventorgraph := models.Graph{}
	var err error
	reqpp.PackageId, err = strconv.Atoi(c.Param("id"))
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&spp.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	reqpp.SetUpdateBy(user.GetUserId(c))
	listpp := make([]models.PatentPackage, 0)
	var count int64
	err = spp.GetPatentIdByPackageId(&reqpp, &listpp, &count)
	reqp.PatentIds = make([]int, len(listpp))
	for i := 0; i < len(listpp); i++ {
		reqp.PatentIds[i] = listpp[i].PatentId
	}
	listp := make([]models.Patent, 0)
	var count2 int64
	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&sp.Service).
		Errors
	err = sp.GetPageByIds(&reqp, &listp, &count2)
	links := make([]models.Link, 0)

	err = e.MakeContext(c).
		MakeOrm().
		MakeService(&sup.Service).
		Errors
	listInventorId := make(map[string]int)
	listup2 := make([]models.InventorPatent, 0)
	FindTheInventorFromPatents(&listInventorId, &listup2, listp)
	usertimes := make(map[int]int)
	for i := 0; i < len(listup2); i++ {
		if usertimes[listup2[i].UserId] == 0 {
			usertimes[listup2[i].UserId] = 1
		} else {
			usertimes[listup2[i].UserId]++
		}

	}
	usertimes1 := rankByWordCount(usertimes)
	var members int
	if len(usertimes1) < 500 {
		members = len(usertimes1)
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
		NodeList[i].NodeId = strconv.FormatInt(int64(usertimes1[i].Key), 10)
		NodeList[i].NodeValue = usertimes1[i].Value
		UserIsNode[i] = true
	}
	userspatents := make([]models.OneUserPatents, members)
	for i := 0; i < len(listup2); i++ {
		for j := 0; j < members; j++ {
			if listup2[i].UserId == usertimes1[j].Key {
				userspatents[j].Patentsid = append(userspatents[j].Patentsid, listup2[i].PatentId)
				break
			}
		}
	}
	useruserrelation1 := make(map[int]int)
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
				useruserrelation1[i*10+j] = RelationExist
			}
		}
	}
	useruserrelation2 := rankByWordCount(useruserrelation1) //给边排序
	for i := 0; i < minresult(firstlinks, len(useruserrelation2)); i++ {
		var nowlink models.Link
		nowlink.Source = strconv.FormatInt(int64(usertimes1[useruserrelation2[i].Key/10].Key), 10)
		nowlink.Target = strconv.FormatInt(int64(usertimes1[useruserrelation2[i].Key%10].Key), 10)
		nowlink.Value = useruserrelation2[i].Value
		links = append(links, nowlink)
	}

	if len(NodeList) < 1 {
		Inventorgraph.Links = links
		Inventorgraph.Nodes = NodeList
		e.GraphOK(NodeList, links, "查询成功")
	}
	useruserrelation3 := make(map[int]int)
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
				useruserrelation3[i*500+j] = RelationExist
			}
		}
	}
	useruserrelation4 := rankByWordCount(useruserrelation3)
	NodelistIdToTimeList := make(map[string]int)

	for i := 0; i < minresult(secondLinks, len(useruserrelation4)); i++ {
		source := useruserrelation4[i].Key / 500
		target := useruserrelation4[i].Key % 500
		if ExtendNodeTime[source] >= 5 {
			continue
		} else {
			if UserIsNode[target] == false {
				UserIsNode[target] = true
				ExtendNodeTime[source]++
				var nowlink models.Link
				var nowNode models.Node
				nowNode.NodeCategory = NodeList[source].NodeCategory
				nowNode.NodeId = strconv.FormatInt(int64(usertimes1[target].Key), 10)
				NodelistIdToTimeList[nowNode.NodeId] = target
				nowlink.Source = strconv.FormatInt(int64(usertimes1[source].Key), 10)
				nowlink.Target = strconv.FormatInt(int64(usertimes1[target].Key), 10)
				nowlink.Value = useruserrelation4[i].Value
				links = append(links, nowlink)
				NodeList = append(NodeList, nowNode)
			}
		}
	}

	useruserrelation5 := make(map[int]int)

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
				useruserrelation5[i*500+j] = RelationExist
			}

		}
	}
	useruserrelation6 := rankByWordCount(useruserrelation5)
	thirdlinks := 50
	for i := 0; i < minresult(thirdlinks, len(useruserrelation6)); i++ {
		source := useruserrelation6[i].Key / 500
		target := useruserrelation6[i].Key % 500

		if ExtendNodeTime[source] >= 2 {
			continue
		} else {
			ExtendNodeTime[source]++
			var nowlink models.Link
			nowlink.Source = NodeList[source].NodeId
			nowlink.Target = NodeList[target].NodeId
			nowlink.Value = useruserrelation6[i].Value
			links = append(links, nowlink)
		}
	}

	listu := make([]amodels.SysUser, 0)
	err = e.MakeContext(c).
		MakeOrm().
		Bind(&requ, nil).
		MakeService(&su.Service).
		Errors
	for i := 0; i < len(NodeList); i++ {
		requ.Id, err = strconv.Atoi(NodeList[i].NodeId)
		var user1 amodels.SysUser
		user1.UserId = requ.Id
		for k, v := range listInventorId {
			if v == requ.Id {
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
	NodeList[0].NodeValue = usertimes1[0].Value
	NodeList[0].NodeSymbolizeSize = 50
	NodeList[0].NodeName = listu[0].Username
	for i := 1; i < first10; i++ {
		NodeList[i].NodeValue = usertimes1[i].Value
		if NodeList[i].NodeValue > max {
			max = NodeList[i].NodeValue
		}
		if NodeList[i].NodeValue < min {
			min = NodeList[i].NodeValue
		}
	}
	for i := first10; i < len(NodeList); i++ {
		NodeList[i].NodeValue = usertimes1[NodelistIdToTimeList[NodeList[i].NodeId]].Value
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
	e.OK(Inventorgraph, "查询成功")
}

// FindTheInventorFromPatents --------------------------------------------------------------------------
// 查找patents中的发明人
func FindTheInventorFromPatents(listInventorId *map[string]int, listup2 *[]models.InventorPatent, listp []models.Patent) error {
	var err error
	count := 0
	for z := 0; z < len(listp); z++ {
		words := make([]string, 0)
		for i := 0; i < len(listp[z].PatentProperties); i++ {
			if listp[z].PatentProperties[i] == '"' && listp[z].PatentProperties[i+1] == 'P' && listp[z].PatentProperties[i+2] == 'I' && listp[z].PatentProperties[i-1] == ',' {
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
			inventorpatent.UserId = (*listInventorId)[words[i]]
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
