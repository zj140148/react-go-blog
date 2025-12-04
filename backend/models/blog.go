package models

import "time"

// Blog 博客文章模型
type Blog struct {
	ID          int       `json:"id"`          // 博客ID
	Title       string    `json:"title"`       // 博客标题
	Summary     string    `json:"summary"`     // 博客摘要
	Content     string    `json:"content"`     // 博客内容
	Author      string    `json:"author"`      // 作者
	CoverImage  string    `json:"cover_image"` // 封面图片URL
	Tags        []string  `json:"tags"`        // 标签列表
	ViewCount   int       `json:"view_count"`  // 浏览次数
	IsPublished bool      `json:"is_published"` // 是否已发布
	CreatedAt   time.Time `json:"created_at"`  // 创建时间
	UpdatedAt   time.Time `json:"updated_at"`  // 更新时间
}

// BlogListItem 博客列表项模型（用于列表展示，不包含完整内容）
type BlogListItem struct {
	ID          int       `json:"id"`          // 博客ID
	Title       string    `json:"title"`       // 博客标题
	Summary     string    `json:"summary"`     // 博客摘要
	Author      string    `json:"author"`      // 作者
	CoverImage  string    `json:"cover_image"` // 封面图片URL
	Tags        []string  `json:"tags"`        // 标签列表
	ViewCount   int       `json:"view_count"`  // 浏览次数
	CreatedAt   time.Time `json:"created_at"`  // 创建时间
}

// PaginationResponse 分页响应模型
type PaginationResponse struct {
	Data       interface{} `json:"data"`        // 数据列表
	Total      int         `json:"total"`       // 总记录数
	Page       int         `json:"page"`        // 当前页码
	PageSize   int         `json:"page_size"`   // 每页大小
	TotalPages int         `json:"total_pages"` // 总页数
}

// BlogListResponse 博客列表分页响应
type BlogListResponse struct {
	Data       []BlogListItem `json:"data"`        // 博客列表数据
	Total      int            `json:"total"`       // 总记录数
	Page       int            `json:"page"`        // 当前页码
	PageSize   int            `json:"page_size"`   // 每页大小
	TotalPages int            `json:"total_pages"` // 总页数
}