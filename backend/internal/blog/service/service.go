package service

import (
	"backend/internal/blog/models"
	"fmt"
	"math"
	"time"
)

// BlogService 博客服务接口
type BlogService interface {
	GetAllBlogs() ([]models.BlogListItem, error)
	GetBlogsPaginated(page, pageSize int) (*models.BlogListResponse, error)
	GetBlogByID(id int) (*models.Blog, error)
}

// blogService 博客服务实现
type blogService struct {
	// 这里可以添加数据库连接等依赖
}

// NewBlogService 创建博客服务实例
func NewBlogService() BlogService {
	return &blogService{}
}

// GetAllBlogs 获取所有博客列表（不包含完整内容）
func (s *blogService) GetAllBlogs() ([]models.BlogListItem, error) {
	// 模拟数据 - 在实际项目中这里应该从数据库获取
	blogs := []models.BlogListItem{
		{
			ID:         1,
			Title:      "Go语言入门指南",
			Summary:    "本文将带你了解Go语言的基础知识和核心概念，适合初学者阅读。",
			Author:     "张三",
			CoverImage: "https://via.placeholder.com/300x200",
			Tags:       []string{"Go", "后端", "教程"},
			ViewCount:  1250,
			CreatedAt:  parseTime("2024-01-15"),
		},
		{
			ID:         2,
			Title:      "React Hooks详解",
			Summary:    "深入理解React Hooks的工作原理，以及如何在项目中正确使用它们。",
			Author:     "李四",
			CoverImage: "https://via.placeholder.com/300x200",
			Tags:       []string{"React", "前端", "Hooks"},
			ViewCount:  980,
			CreatedAt:  parseTime("2024-01-20"),
		},
		{
			ID:         3,
			Title:      "微服务架构设计",
			Summary:    "探讨微服务架构的设计原则、最佳实践以及常见问题的解决方案。",
			Author:     "王五",
			CoverImage: "https://via.placeholder.com/300x200",
			Tags:       []string{"架构", "微服务", "设计"},
			ViewCount:  1500,
			CreatedAt:  parseTime("2024-02-01"),
		},
		{
			ID:         4,
			Title:      "Docker容器化实践",
			Summary:    "从零开始学习Docker，包括镜像构建、容器管理和编排等核心概念。",
			Author:     "赵六",
			CoverImage: "https://via.placeholder.com/300x200",
			Tags:       []string{"Docker", "DevOps", "容器"},
			ViewCount:  750,
			CreatedAt:  parseTime("2024-02-10"),
		},
		{
			ID:         5,
			Title:      "数据库性能优化",
			Summary:    "分享数据库性能优化的实用技巧，包括索引优化、查询优化等。",
			Author:     "钱七",
			CoverImage: "https://via.placeholder.com/300x200",
			Tags:       []string{"数据库", "优化", "性能"},
			ViewCount:  2100,
			CreatedAt:  parseTime("2024-02-15"),
		},
		{
			ID:         6,
			Title:      "前端工程化实践",
			Summary:    "介绍现代前端工程化的概念和工具，包括构建、测试、部署等流程。",
			Author:     "孙八",
			CoverImage: "https://via.placeholder.com/300x200",
			Tags:       []string{"前端", "工程化", "工具"},
			ViewCount:  890,
			CreatedAt:  parseTime("2024-02-20"),
		},
		{
			ID:         7,
			Title:      "API设计最佳实践",
			Summary:    "探讨RESTful API设计的最佳实践，包括版本控制、错误处理、安全等。",
			Author:     "周九",
			CoverImage: "https://via.placeholder.com/300x200",
			Tags:       []string{"API", "设计", "REST"},
			ViewCount:  1320,
			CreatedAt:  parseTime("2024-03-01"),
		},
		{
			ID:         8,
			Title:      "机器学习入门",
			Summary:    "机器学习的基础概念和算法介绍，适合想要入门AI的开发者。",
			Author:     "吴十",
			CoverImage: "https://via.placeholder.com/300x200",
			Tags:       []string{"机器学习", "AI", "算法"},
			ViewCount:  1680,
			CreatedAt:  parseTime("2024-03-05"),
		},
	}

	return blogs, nil
}

// GetBlogsPaginated 分页获取博客列表
func (s *blogService) GetBlogsPaginated(page, pageSize int) (*models.BlogListResponse, error) {
	// 获取所有博客
	allBlogs, err := s.GetAllBlogs()
	if err != nil {
		return nil, err
	}

	total := len(allBlogs)
	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	// 计算分页范围
	start := (page - 1) * pageSize
	end := start + pageSize

	// 防止越界
	if start > total {
		return &models.BlogListResponse{
			Data:       []models.BlogListItem{},
			Total:      total,
			Page:       page,
			PageSize:   pageSize,
			TotalPages: totalPages,
		}, nil
	}

	if end > total {
		end = total
	}

	// 获取当前页数据
	pagedBlogs := allBlogs[start:end]

	return &models.BlogListResponse{
		Data:       pagedBlogs,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

// GetBlogByID 根据ID获取博客详情
func (s *blogService) GetBlogByID(id int) (*models.Blog, error) {
	// 模拟数据 - 在实际项目中这里应该从数据库获取
	blogs := map[int]*models.Blog{
		1: {
			ID:          1,
			Title:       "Go语言入门指南",
			Summary:     "本文将带你了解Go语言的基础知识和核心概念，适合初学者阅读。",
			Content:     "Go语言是由Google开发的开源编程语言，它具有简洁的语法、高效的并发性能和快速的编译速度。本文将详细介绍Go语言的特性，包括变量声明、函数定义、结构体、接口、goroutine和channel等核心概念。通过实际示例，你将学会如何使用Go语言构建高性能的后端服务。",
			Author:      "张三",
			CoverImage:  "https://via.placeholder.com/300x200",
			Tags:        []models.Tag{{ID: 1, Name: "Go"}, {ID: 2, Name: "后端"}, {ID: 3, Name: "教程"}},
			ViewCount:   1250,
			IsPublished: true,
			CreatedAt:   parseTime("2024-01-15"),
			UpdatedAt:   parseTime("2024-01-15"),
		},
		2: {
			ID:          2,
			Title:       "React Hooks详解",
			Summary:     "深入理解React Hooks的工作原理，以及如何在项目中正确使用它们。",
			Content:     "React Hooks是React 16.8引入的新特性，它让你在不编写class的情况下使用state以及其他的React特性。本文将详细介绍useState、useEffect、useContext、useReducer等常用Hooks的使用方法和最佳实践。你将学会如何自定义Hooks，以及如何在函数组件中管理状态和副作用。",
			Author:      "李四",
			CoverImage:  "https://via.placeholder.com/300x200",
			Tags:        []models.Tag{{ID: 4, Name: "React"}, {ID: 5, Name: "前端"}, {ID: 6, Name: "Hooks"}},
			ViewCount:   980,
			IsPublished: true,
			CreatedAt:   parseTime("2024-01-20"),
			UpdatedAt:   parseTime("2024-01-20"),
		},
		3: {
			ID:          3,
			Title:       "微服务架构设计",
			Summary:     "探讨微服务架构的设计原则、最佳实践以及常见问题的解决方案。",
			Content:     "微服务架构是一种将应用程序构建为一组小型服务的方法，每个服务运行在自己的进程中，服务之间通过轻量级的通信机制进行通信。本文将介绍微服务架构的优势和挑战，包括服务拆分原则、数据一致性、服务发现、负载均衡、容错处理等关键技术。通过实际案例，你将学会如何设计和实现可扩展的微服务系统。",
			Author:      "王五",
			CoverImage:  "https://via.placeholder.com/300x200",
			Tags:        []models.Tag{{ID: 7, Name: "架构"}, {ID: 8, Name: "微服务"}, {ID: 9, Name: "设计"}},
			ViewCount:   1500,
			IsPublished: true,
			CreatedAt:   parseTime("2024-02-01"),
			UpdatedAt:   parseTime("2024-02-01"),
		},
	}

	blog, exists := blogs[id]
	if !exists {
		return nil, fmt.Errorf("博客不存在")
	}

	return blog, nil
}

// parseTime 辅助函数，解析时间字符串
func parseTime(timeStr string) time.Time {
	t, _ := time.Parse("2006-01-02", timeStr)
	return t
}
