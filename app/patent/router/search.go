package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerSearchRouter)
}

// 需认证的路由代码
func registerSearchRouter(v1 *gin.RouterGroup) {
	api := apis.SysUser{}
	r := v1.Group("/search")
	{
		r.GET("simple", api.GetPage)
		//r.GET("table", api.Get)
		//r.GET("advance", api.Insert)
	}
}
