package system

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	{
		baseRouter.GET("login", func(c *gin.Context) {
			c.JSON(http.StatusOK, "login")
			return
		})
	}
	return baseRouter
}
