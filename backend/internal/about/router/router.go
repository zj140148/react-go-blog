package router

import (
	"backend/internal/about/controller"
	"github.com/gin-gonic/gin"
)

// SetupAboutRoutes 设置关于页面的路由
func SetupAboutRoutes(r *gin.Engine) {
	aboutController := controller.NewAboutController()
	
	r.GET("/api/about-md", aboutController.GetAboutMD)
}