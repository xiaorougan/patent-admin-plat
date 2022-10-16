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

	r := v1.Group("/sys-list").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetLists) //测试成功
		//r.GET("/:patent_id", api.GetLists) 错误无法使用id
		r.POST("", api.InsertListsByPatentId) //测试成功
		//r.POST("/:id", api.InsertListsByPatentId)
		//r.POST("/:ti", api.InsertListsByPatentName)
		r.PUT("", api.UpdateLists)        //测试成功
		r.DELETE("/:id", api.DeleteLists) //测试成功

	}
}
