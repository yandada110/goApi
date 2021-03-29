package routes

import (
	"github.com/wuyan94zl/api/app/controllers/article"
	"github.com/gin-gonic/gin"
	"github.com/wuyan94zl/api/app/controllers/admin"
	"github.com/wuyan94zl/api/pkg/utils"
)

// 注册路由列表
func PermissionRouter(router *gin.RouterGroup) {
	utils.AddRoute(router, "POST", "/admin/role", admin.SetRole)

	// start admin
	utils.AddRoute(router,"POST","/admin/create",admin.Create)
	utils.AddRoute(router,"POST","/admin/update",admin.Update)
	utils.AddRoute(router,"GET","/admin/delete",admin.Delete)
	utils.AddRoute(router,"GET","/admin/info",admin.Info)
	utils.AddRoute(router,"POST","/admin/paginate",admin.Paginate)
	// end admin

	// start article
	utils.AddRoute(router,"POST","/article/create",article.Create)
	utils.AddRoute(router,"POST","/article/update",article.Update)
	utils.AddRoute(router,"GET","/article/delete",article.Delete)
	utils.AddRoute(router,"GET","/article/info",article.Info)
	utils.AddRoute(router,"POST","/article/paginate",article.Paginate)
	// end article
}


