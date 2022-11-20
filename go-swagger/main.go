package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "go-swagger-dly/docs"
	"net/http"
)

// @title 标题
// @version 1.0 (版本)
// @description 声明（可不写）
// @termsOfService https://www.test.com

// @contact.name www.test.com
// @contact.url https://www.test.com
// @contact.email me@test.me

// @license.name Apache 2.0 (必填)
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath
func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := r.Group("/api/v1")
	{
		v1.GET("/hello", HandleHello)
	}
	r.Run(":8080")
}

// HandleHello @Summary 测试SayHello
// @Description 向你说Hello
// @Tags 测试
// @Accept json
// @Param who query string true "人名"
// @Success 200 {string} string "{"msg": "hello Razeen"}"
// @Failure 400 {string} string "{"msg": "who are you"}"
// @Router /hello [get]
func HandleHello(c *gin.Context) {
	who := c.Query("who")

	if who == "" {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "who are u?"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "hello " + who})
}
