package api

import "go-admin/app/patent/router"

func init() {
	AppRouters = append(AppRouters, router.InitRouter)
}
