package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/zngue/go_helper/pkg/api"
	"io/ioutil"
)

func NotifyMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		all, err := ioutil.ReadAll(ctx.Request.Body)
		if err != nil {
			api.DataWithErr(ctx,err,nil)
			ctx.Abort()
			return
		}
		if len(all)==0 {
			api.DataWithErr(ctx,errors.New("body is null"),nil)
			ctx.Abort()
			return
		}
		//ctx.Next()
	}
}
