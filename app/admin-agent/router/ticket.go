package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/admin-agent/apis"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerTicketRouter)
}

// 需认证的路由代码
func registerTicketRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {

	api := apis.Ticket{}

	r := v1.Group("/ticket").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", api.GetAllTicketPages)
		r.POST("", api.CreateTicket)
		r.POST("/:id/update", api.UpdateTicket)
		r.POST("/:id/close", api.CloseTicket)

		r.DELETE("/:id", api.RemoveTicket)
	}

}
