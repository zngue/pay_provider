package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg/api"
	"github.com/zngue/go_helper/pkg/common_run"
	"github.com/zngue/pay_provider/app/model"
	"github.com/zngue/pay_provider/app/router"
	"gorm.io/gorm"
)
func main() {

	common_run.CommonGinRun(
		common_run.FnRouter(func(engine *gin.Engine) {
			apis := engine.Group("mchprovider")
			router.Router(apis)
			engine.NoRoute(func(c *gin.Context) {
				api.Error(c,api.Code(404),api.Msg("路由不存在"))
			})
		}),
		common_run.MysqlConn(func(db *gorm.DB) {
			db.AutoMigrate(model.NewMerchant(),model.NewOrder(),model.NewPayLog())
		}),
		common_run.IsRegisterCenter(false),
	)

}



