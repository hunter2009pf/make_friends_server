package routes

import (
	"angel_clothes.make_friends/m/v2/controllers/handlers"
	"github.com/gin-gonic/gin"
)

func UserRoute(r *gin.Engine) {
	// 注册用户相关路由
	r.GET("/onLogin", handlers.OnLoginHandler)
	r.POST("/uploadPhoto", handlers.UploadPhoto)
	r.POST("/uploadPhotoToList", handlers.UploadPhotoToList)
	r.DELETE("/deletePhotoFromList", handlers.DeletePhotoFromList)
	r.POST("/uploadInfo", handlers.UploadInfo)
	r.GET("/randomGetInfo", handlers.RandomGetInfo)
	r.GET("/randomGetPhotos", handlers.RandomGetPhotos)
}
