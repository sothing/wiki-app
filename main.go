package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main()  {
	fmt.Println("gin.version:",gin.Version)

	//构建路由
	app := gin.New()

	//全局的中间件
	app.Use(gin.Logger())

	app.Run(":" + fmt.Sprintf("%d",config.ServerConfig.Port))

}
