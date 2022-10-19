package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
	"go-admin/app/admin/router"
	"go-admin/app/patent/apis"
	"go-admin/common/middleware"
)

func init() {
	router.RouterCheckRole = append(router.RouterCheckRole, registerUserPatentRouter)
}

// 需认证的路由代码
func registerUserPatentRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.UserPatent{}

	r1 := v1.Group("/user-patent").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r1.GET("/claim", api.GetClaimPatentByUserId)
		r1.GET("/collection", api.GetCollectionPatentByUserId)
		r1.POST("", api.InsertUserPatentRelationship)

		r1.DELETE("/:patent_id/:type", api.DeleteUserPatentRelationship)

		r1.PUT("", api.GetClaimPatentByUserId)
	}

}
