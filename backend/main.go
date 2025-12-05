package main

import (
	"backend/internal/app"
)

func main() {
	// 创建应用实例
	application := app.NewApp()

	// 初始化应用（中间件、路由等）
	application.Initialize()

	// 获取引擎并启动服务
	engine := application.GetEngine()
	engine.Run(":5000")
}
