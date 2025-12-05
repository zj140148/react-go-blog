package controller

import (
	"backend/internal/about/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AboutController 关于页面控制器
type AboutController struct {
	aboutService service.AboutService
}

// NewAboutController 创建关于页面控制器实例
func NewAboutController() *AboutController {
	return &AboutController{
		aboutService: service.NewAboutService(),
	}
}

// GetAboutMD 获取关于页面的Markdown内容
// @Summary 获取关于页面内容
// @Description 获取关于页面的Markdown内容，包括作者信息和版本
// @Tags 关于
// @Accept json
// @Produce json
// @Success 200 {object} models.AboutResponse
// @Failure 500 {object} models.AboutResponse
// @Router /api/about-md [get]
func (ctrl *AboutController) GetAboutMD(c *gin.Context) {
	// 调用服务层获取关于页面内容
	aboutInfo, err := ctrl.aboutService.GetAboutContent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "获取关于页面失败",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "获取关于页面成功",
		"data":    aboutInfo,
	})
}