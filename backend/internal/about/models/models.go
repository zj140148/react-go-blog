package models

// AboutInfo 关于页面信息模型
type AboutInfo struct {
	Content string `json:"content"` // Markdown内容
	Author  string `json:"author"`  // 作者
	Version string `json:"version"` // 版本号
}

// AboutResponse 关于页面响应模型
type AboutResponse struct {
	Success bool       `json:"success"`
	Message string     `json:"message"`
	Data    AboutInfo  `json:"data"`
}