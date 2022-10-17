package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/patent/apis"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSearchRouter)
}

// 需认证的路由代码
func registerSearchRouter(v1 *gin.RouterGroup) {
	api := apis.Search{}
	r := v1.Group("/search")
	{
		r.POST("/simple", api.SimpleSearch)
		//r.GET("table", api.Get)
		//r.GET("advance", api.Insert)
	}
}
