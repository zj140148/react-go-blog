package router

import (
	"backend/internal/blog/controller"
	"github.com/gin-gonic/gin"
)

// SetupBlogRoutes 设置博客相关的路由
func SetupBlogRoutes(r *gin.Engine) {
	blogController := controller.NewBlogController()
	
	blogGroup := r.Group("/api/blogs")
	{
		blogGroup.GET("", blogController.GetBlogs)                    // 获取博客列表
		blogGroup.GET("/paginated", blogController.GetBlogsPaginated) // 分页获取博客列表
		blogGroup.GET("/:id", blogController.GetBlogByID)             // 获取博客详情
	}
}