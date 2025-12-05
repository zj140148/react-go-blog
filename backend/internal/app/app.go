package app

import (
	aboutrouter "backend/internal/about/router"
	blogrouter "backend/internal/blog/router"

	"github.com/gin-gonic/gin"
)

// App 应用结构体
type App struct {
	engine *gin.Engine
}

// NewApp 创建新的应用实例
func NewApp() *App {
	return &App{
		engine: gin.New(),
	}
}

// SetupMiddleware 设置中间件
func (app *App) SetupMiddleware() {
	// 添加Logger中间件
	app.engine.Use(gin.Logger())

	// 添加Recovery中间件
	app.engine.Use(gin.Recovery())

	// 添加CORS中间件（与原来保持一致）
	app.engine.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}
		c.Next()
	})
}

// SetupRoutes 设置路由
func (app *App) SetupRoutes() {
	// 设置about路由
	aboutrouter.SetupAboutRoutes(app.engine)

	// 设置blog路由
	blogrouter.SetupBlogRoutes(app.engine)
}

// GetEngine 获取gin引擎
func (app *App) GetEngine() *gin.Engine {
	return app.engine
}

// Initialize 初始化应用
func (app *App) Initialize() {
	app.SetupMiddleware()
	app.SetupRoutes()
}
