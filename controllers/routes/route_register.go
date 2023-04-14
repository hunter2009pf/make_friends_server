package routes

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 路由注册函数
func SetupRouter() *gin.Engine {
	r := gin.Default()
	// 设置中间件，必须在路由注册之前
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// r.Use(middleware.AccessControl())
	UserRoute(r)
	return r
}
