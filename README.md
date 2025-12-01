# React-Go 博客项目

这是一个使用React前端和Go后端构建的简单博客系统。

## 项目结构

```
react-go-blog/
├── backend/                 # Go后端代码
│   ├── handlers/           # 请求处理器
│   │   ├── about/         # 关于页面处理器
│   │   └── blog/          # 博客相关处理器
│   ├── models/            # 数据模型
│   ├── router/            # 路由配置
│   ├── services/          # 业务逻辑层
│   └── main.go            # 后端入口文件
├── frontend/     # React前端代码
│   ├── src/
│   │   ├── components/    # React组件
│   │   ├── pages/         # 页面组件
│   │   ├── services/      # API服务
│   │   └── App.jsx        # 前端入口文件
└── docker-compose.yml      # Docker配置
```

## 功能实现

### 后端功能

1. **博客列表API** (`GET /api/blogs`)
   - 获取所有已发布的博客列表
   - 不包含完整内容，只包含摘要信息

2. **博客详情API** (`GET /api/blogs/{id}`)
   - 根据ID获取博客的完整内容
   - 每次访问会增加浏览次数

3. **关于页面API** (`GET /api/about-md`)
   - 获取关于页面的Markdown内容

### 前端功能

1. **博客列表页面** (`/blog`)
   - 展示所有博客文章的卡片列表
   - 点击卡片可以跳转到博客详情页面

2. **博客详情页面** (`/blog/{id}`)
   - 展示博客的完整内容
   - 显示作者、发布时间、浏览量等信息
   - 显示文章标签

## 如何运行
### 推荐使用Docker Compose启动开发环境
进入项目根目录，执行以下命令启动开发环境：
```bash
docker compose up -d
```
### 分别启动后端服务器和前端开发服务器
#### 启动后端服务器

```bash
cd backend
go mod tidy
go run main.go
```

后端服务器将在 `http://localhost:5000` 上运行。

#### 启动前端开发服务器

```bash
cd techspace-frontend
npm install
npm run dev
```

前端开发服务器将在 `http://localhost:3000` 上运行。

## 文档
详细的前端文档请参考 [techspace-frontend/document/FRONTEND_README.md](techspace-frontend/document/FRONTEND_README.md)
详细的API文档请参考 [backend/docs/blog_api.md](backend/docs/blog_api.md)

## 技术栈

### 后端
- Go
- Gin Web框架
- 内存存储（模拟数据库）

### 前端
- React
- React Router
- Vite
- CSS Modules

## 注意事项

1. 当前版本使用内存存储，重启服务器后数据会重置
2. 前端API请求地址硬编码为 `http://localhost:5000/api`，实际项目中应考虑环境变量
3. 博客内容的Markdown渲染使用了简单的实现，生产环境建议使用专门的Markdown解析库

## 开发日记
- 11.9 学习如何使用vite创建react项目，前端初步写了一点框架，写了导航栏的样式
- 11.10 完善了主页home的样式，blog页面写了部分样式
- 11.10晚 使用docker配置了开发环境，现在可以一键配置开发环境。优化了页面的结构，开始使用组件化开发，将页面拆分成多个组件，方便一键管理，减少无效复用和代码量
- 11.30 完成了contact页面，重新设计了页面文字的样式，添加了博客详细内容的页面PageDetail并写好了页面的样式。
   写好了PageWrapper组件，用于包裹页面内容，添加了页面过渡效果
   