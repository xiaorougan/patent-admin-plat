package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/admin/router"
	"go-admin/app/patent/apis"
	"go-admin/common/middleware"
)

func init() {
	router.RouterCheckRole = append(router.RouterCheckRole, registerSysListRouter)
}

// 需认证的路由代码
func registerSysListRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.SysList{}

	r := v1.Group("/patent-list").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("/get_patent_lists", api.GetLists)
		r.GET("/get_by_patent_id/:patent_id", api.GetPatentById)
		//r.GET("/patent-name/:ti", api.GetPatentByName)
		r.POST("/post_a_patent/", api.InsertLists)
		r.PUT("/change_a_patent/", api.UpdateLists)
		r.DELETE("/delete_a_patent_by_id/:patent_id", api.DeletePatentByPatentId)
	}

	r1 := v1.Group("/patent-relationship").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r1.GET("", api.GetLists)
		r1.GET("/get-patent-lists-by-userid/:user_id", api.GetPatentById)
	}

}

//一般来说,Controller是Handler,但Handler不一定是Controller。
