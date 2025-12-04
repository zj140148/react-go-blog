package services

import (
	"backend/models"
	"errors"
	"fmt"
	"sort"
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
	blogs []models.Blog
}

// NewBlogService 创建博客服务实例
func NewBlogService() BlogService {
	// 初始化一些示例博客数据
	blogs := []models.Blog{
		{
			ID:          1,
			Title:       "React入门指南",
			Summary:     "这是一篇关于React入门的指南，介绍了React的基本概念、组件化开发、JSX语法等。",
			Content:     "# React入门指南\n\nReact是Facebook开发的用于构建用户界面的JavaScript库。本指南将带你从零开始学习React。\n\n## 1. React简介\n\nReact是一个用于构建用户界面的JavaScript库，特别适合构建单页应用程序。它采用组件化的开发模式，让UI开发变得更加模块化和可维护。\n\n## 2. 基本概念\n\n### 组件\n\nReact应用由组件构成，组件是可重用的代码块，每个组件管理自己的状态。\n\n```jsx\nfunction Welcome(props) {\n  return <h1>Hello, {props.name}</h1>;\n}\n```\n\n### JSX\n\nJSX是一种JavaScript的语法扩展，它允许我们在JavaScript代码中编写类似HTML的代码。\n\n```jsx\nconst element = <h1>Hello, world!</h1>;\n```\n\n### Props和State\n\nProps是组件的输入，类似于函数参数。State是组件的内部状态，当状态改变时，组件会重新渲染。\n\n## 3. 环境搭建\n\n要开始React开发，你需要安装Node.js和npm。然后使用Create React App创建新项目：\n\n```bash\nnpx create-react-app my-app\ncd my-app\nnpm start\n```\n\n## 4. 总结\n\nReact是一个强大而灵活的前端框架，通过组件化的方式让UI开发更加简单和高效。希望这篇指南能帮助你入门React开发。",
			Author:      "张三",
			CoverImage:  "/images/react-bg.jpg",
			Tags:        []string{"React", "前端", "JavaScript"},
			ViewCount:   128,
			IsPublished: true,
			CreatedAt:   time.Now().AddDate(0, 0, -7),
			UpdatedAt:   time.Now().AddDate(0, 0, -7),
		},
		{
			ID:          2,
			Title:       "Go语言入门指南",
			Summary:     "这是一篇关于Go语言入门的指南，介绍了Go的基本语法、并发编程、标准库等。",
			Content:     "# Go语言入门指南\n\nGo是Google开发的开源编程语言，专为构建简单、可靠和高效的软件而设计。\n\n## 1. Go语言简介\n\nGo（又称Golang）是Google开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。\n\n## 2. 基本语法\n\n### 变量声明\n\nGo语言有多种声明变量的方式：\n\n```go\n// 标准声明\nvar name string = \"张三\"\n\n// 类型推断\nvar age = 25\n\n// 短变量声明（只能在函数内使用）\nisStudent := true\n```\n\n### 函数\n\n```go\nfunc add(a, b int) int {\n    return a + b\n}\n\nfunc main() {\n    result := add(3, 5)\n    fmt.Println(result) // 输出: 8\n}\n```\n\n## 3. 并发编程\n\nGo语言的并发是其核心特性之一，通过goroutine和channel实现：\n\n```go\nfunc say(s string) {\n    for i := 0; i < 5; i++ {\n        time.Sleep(100 * time.Millisecond)\n        fmt.Println(s)\n    }\n}\n\nfunc main() {\n    go say(\"world\")\n    say(\"hello\")\n}\n```\n\n## 4. 标准库\n\nGo语言拥有丰富的标准库，包括：\n- net/http：HTTP客户端和服务器实现\n- encoding/json：JSON编码和解码\n- fmt：格式化I/O\n- os：操作系统接口\n\n## 5. 总结\n\nGo语言简洁、高效，特别适合构建网络服务和分布式系统。希望这篇指南能帮助你入门Go语言开发。",
			Author:      "李四",
			CoverImage:  "/images/go-bg.jpg",
			Tags:        []string{"Go", "后端", "编程语言"},
			ViewCount:   95,
			IsPublished: true,
			CreatedAt:   time.Now().AddDate(0, 0, -5),
			UpdatedAt:   time.Now().AddDate(0, 0, -5),
		},
		{
			ID:          3,
			Title:       "全栈开发最佳实践",
			Summary:     "分享全栈开发中的最佳实践，包括前后端分离、API设计、数据库优化等内容。",
			Content:     "# 全栈开发最佳实践\n\n全栈开发需要掌握前端和后端技术，本文将分享一些全栈开发的最佳实践。\n\n## 1. 前后端分离\n\n前后端分离是现代Web应用的主流架构模式，它有以下优势：\n- 前后端可以独立开发、部署\n- 提高开发效率\n- 便于团队协作\n\n## 2. RESTful API设计\n\n设计良好的API应该遵循RESTful原则：\n- 使用合适的HTTP方法（GET, POST, PUT, DELETE）\n- 使用名词而非动词作为资源路径\n- 返回适当的HTTP状态码\n\n## 3. 数据库设计\n\n良好的数据库设计是应用性能的基础：\n- 合理设计表结构\n- 创建适当的索引\n- 避免N+1查询问题\n\n## 4. 安全性考虑\n\n- 使用HTTPS加密传输\n- 实施身份验证和授权\n- 防止常见的安全漏洞（XSS, CSRF, SQL注入等）\n\n## 5. 性能优化\n\n- 前端：代码分割、懒加载、缓存策略\n- 后端：数据库优化、缓存、异步处理\n\n## 6. 总结\n\n全栈开发需要综合考虑多个方面，通过遵循最佳实践，可以构建出高质量、可维护的应用程序。",
			Author:      "王五",
			CoverImage:  "/images/fullstack-bg.jpg",
			Tags:        []string{"全栈开发", "最佳实践", "架构"},
			ViewCount:   156,
			IsPublished: true,
			CreatedAt:   time.Now().AddDate(0, 0, -3),
			UpdatedAt:   time.Now().AddDate(0, 0, -2),
		},
		// 添加更多测试数据以支持分页
		{
			ID:          4,
			Title:       "Vue.js 3.0 新特性详解",
			Summary:     "详细介绍Vue.js 3.0的新特性，包括Composition API、性能优化、TypeScript支持等。",
			Content:     "# Vue.js 3.0 新特性详解\n\nVue.js 3.0带来了许多激动人心的新特性，本文将详细介绍这些变化。\n\n## 1. Composition API\n\nComposition API是Vue 3.0最重要的新特性之一，它提供了更灵活的代码组织方式。\n\n```javascript\nimport { ref, computed } from 'vue'\n\nexport default {\n  setup() {\n    const count = ref(0)\n    const doubled = computed(() => count.value * 2)\n    \n    return { count, doubled }\n  }\n}\n```\n\n## 2. 性能优化\n\nVue 3.0在性能方面有了显著提升：\n- 更小的包体积\n- 更快的渲染速度\n- 更好的内存管理\n\n## 3. TypeScript支持\n\nVue 3.0对TypeScript提供了更好的支持，让类型安全更加完善。\n\n## 4. 总结\n\nVue.js 3.0的新特性让开发体验更加优秀，值得学习和使用。",
			Author:      "赵六",
			CoverImage:  "/images/vue-bg.jpg",
			Tags:        []string{"Vue.js", "前端", "JavaScript"},
			ViewCount:   89,
			IsPublished: true,
			CreatedAt:   time.Now().AddDate(0, 0, -2),
			UpdatedAt:   time.Now().AddDate(0, 0, -2),
		},
		{
			ID:          5,
			Title:       "Docker容器化部署实践",
			Summary:     "介绍如何使用Docker进行应用容器化部署，包括Dockerfile编写、镜像优化等内容。",
			Content:     "# Docker容器化部署实践\n\nDocker已经成为现代应用部署的标准工具，本文将分享Docker部署的最佳实践。\n\n## 1. Docker基础\n\nDocker是一个开源的容器化平台，它允许开发者将应用及其依赖打包到轻量级容器中。\n\n## 2. Dockerfile最佳实践\n\n```dockerfile\nFROM node:16-alpine\nWORKDIR /app\nCOPY package*.json ./\nRUN npm ci --only=production\nCOPY . .\nEXPOSE 3000\nCMD [\"npm\", \"start\"]\n```\n\n## 3. 镜像优化\n\n- 使用多阶段构建\n- 选择合适的基础镜像\n- 减少镜像层数\n\n## 4. 总结\n\nDocker容器化部署能够提高应用的可移植性和一致性，是现代DevOps的重要工具。",
			Author:      "钱七",
			CoverImage:  "/images/docker-bg.jpg",
			Tags:        []string{"Docker", "DevOps", "部署"},
			ViewCount:   112,
			IsPublished: true,
			CreatedAt:   time.Now().AddDate(0, 0, -1),
			UpdatedAt:   time.Now().AddDate(0, 0, -1),
		},
		{
			ID:          6,
			Title:       "TypeScript进阶技巧",
			Summary:     "分享TypeScript的高级用法和最佳实践，包括泛型、装饰器、高级类型等。",
			Content:     "# TypeScript进阶技巧\n\nTypeScript为JavaScript添加了类型系统，本文将介绍一些进阶技巧。\n\n## 1. 泛型编程\n\n泛型让我们能够编写可重用的类型安全代码：\n\n```typescript\nfunction identity<T>(arg: T): T {\n  return arg\n}\n\ninterface GenericIdentityFn<T> {\n  (arg: T): T\n}\n```\n\n## 2. 高级类型\n\nTypeScript提供了许多高级类型特性：\n- 联合类型和交叉类型\n- 条件类型\n- 映射类型\n\n## 3. 装饰器\n\n装饰器是一种特殊类型的声明，可以用来修改类和行为。\n\n## 4. 总结\n\n掌握TypeScript的进阶技巧能够让我们编写更加健壮和可维护的代码。",
			Author:      "孙八",
			CoverImage:  "/images/typescript-bg.jpg",
			Tags:        []string{"TypeScript", "前端", "JavaScript"},
			ViewCount:   76,
			IsPublished: true,
			CreatedAt:   time.Now().AddDate(0, 0, -4),
			UpdatedAt:   time.Now().AddDate(0, 0, -4),
		},
		{
			ID:          7,
			Title:       "微服务架构设计原则",
			Summary:     "探讨微服务架构的设计原则和最佳实践，包括服务拆分、通信机制、数据一致性等。",
			Content:     "# 微服务架构设计原则\n\n微服务架构已经成为现代应用开发的主流模式，本文将探讨其设计原则。\n\n## 1. 服务拆分原则\n\n- 单一职责原则\n- 业务边界清晰\n- 数据自治\n\n## 2. 服务通信\n\n- 同步通信：REST API、gRPC\n- 异步通信：消息队列、事件驱动\n\n## 3. 数据一致性\n\n- 最终一致性\n- 分布式事务\n- 事件溯源\n\n## 4. 总结\n\n微服务架构能够提高系统的可扩展性和可维护性，但需要合理的设计和实施。",
			Author:      "周九",
			CoverImage:  "/images/microservices-bg.jpg",
			Tags:        []string{"微服务", "架构", "后端"},
			ViewCount:   134,
			IsPublished: true,
			CreatedAt:   time.Now().AddDate(0, 0, -6),
			UpdatedAt:   time.Now().AddDate(0, 0, -6),
		},
		{
			ID:          8,
			Title:       "前端性能优化实战",
			Summary:     "分享前端性能优化的实战经验，包括代码分割、懒加载、缓存策略等技巧。",
			Content:     "# 前端性能优化实战\n\n前端性能优化是提升用户体验的关键，本文将分享一些实用的优化技巧。\n\n## 1. 代码分割\n\n使用动态导入实现代码分割：\n\n```javascript\nconst LazyComponent = React.lazy(() => import('./LazyComponent'))\n\nfunction App() {\n  return (\n    <Suspense fallback={<div>Loading...</div>}>\n      <LazyComponent />\n    </Suspense>\n  )\n}\n```\n\n## 2. 图片优化\n\n- 使用现代图片格式\n- 实现懒加载\n- 响应式图片\n\n## 3. 缓存策略\n\n- 浏览器缓存\n- CDN缓存\n- Service Worker\n\n## 4. 总结\n\n前端性能优化需要综合考虑多个方面，通过合理的优化策略可以显著提升用户体验。",
			Author:      "吴十",
			CoverImage:  "/images/performance-bg.jpg",
			Tags:        []string{"性能优化", "前端", "最佳实践"},
			ViewCount:   98,
			IsPublished: true,
			CreatedAt:   time.Now().AddDate(0, 0, -8),
			UpdatedAt:   time.Now().AddDate(0, 0, -8),
		},
	}

	return &blogService{
		blogs: blogs,
	}
}

// GetAllBlogs 获取所有已发布的博客列表
func (s *blogService) GetAllBlogs() ([]models.BlogListItem, error) {
	var blogList []models.BlogListItem

	for _, blog := range s.blogs {
		if blog.IsPublished {
			blogList = append(blogList, models.BlogListItem{
				ID:         blog.ID,
				Title:      blog.Title,
				Summary:    blog.Summary,
				Author:     blog.Author,
				CoverImage: blog.CoverImage,
				Tags:       blog.Tags,
				ViewCount:  blog.ViewCount,
				CreatedAt:  blog.CreatedAt,
			})
		}
	}

	// 按ID降序排列（ID大的在前面，ID小的在后面）
	sort.Slice(blogList, func(i, j int) bool {
		return blogList[i].ID > blogList[j].ID
	})

	return blogList, nil
}

// GetBlogsPaginated 分页获取博客列表
func (s *blogService) GetBlogsPaginated(page, pageSize int) (*models.BlogListResponse, error) {
	// 参数验证
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 6 // 默认每页6条
	}

	// 获取所有已发布的博客
	allBlogs, err := s.GetAllBlogs()
	if err != nil {
		return nil, err
	}

	total := len(allBlogs)
	totalPages := (total + pageSize - 1) / pageSize // 向上取整

	// 计算分页范围
	start := (page - 1) * pageSize
	end := start + pageSize

	// 防止越界
	if start >= total {
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
	for _, blog := range s.blogs {
		if blog.ID == id && blog.IsPublished {
			// 增加浏览次数
			blog.ViewCount++
			return &blog, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("ID为%d的博客不存在或未发布", id))
}
