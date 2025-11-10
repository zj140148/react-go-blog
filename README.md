# 📝 Personal Blog System (React + Go)

一个基于 **React 前端** 与 **Go 后端** 的全栈博客系统，支持用户登录、文章管理与 Markdown 渲染。  
项目旨在练习现代前后端分离架构与工程化开发流程。

---

## 🚀 技术栈

### 前端（React）
- ⚛️ **React 19**：前端框架
- 🧭 **React Router DOM**：路由管理
- 🔗 **Axios**：前后端数据交互
- ✍️ **React Markdown + remark-gfm + rehype-highlight + rehype-raw**：Markdown 渲染与代码高亮
- 🕒 **date-fns**：日期格式化
- 🎨 **CSS Modules / 组件化样式**：样式管理与隔离

### 后端（Go）
- 🧰 **Gin**：Web 框架
- 🗄️ **GORM**：数据库 ORM
- 🔐 **JWT**：用户认证
- 🛢️ **MySQL / SQLite**：数据存储

---

## 📁 项目结构

```bash
.
├── frontend/                # React 前端
│   ├── src/
│   │   ├── pages/           # 页面（Home、Login、Post等）
│   │   ├── components/      # 通用组件（Header、Footer等）
│   │   ├── styles/          # 全局样式
│   │   └── App.jsx
│   └── package.json
│
└── backend/                 # Go 后端
    ├── main.go
    ├── models/              # 数据模型
    ├── handlers/            # 路由处理
    ├── middleware/          # 中间件（JWT、CORS等）
    ├── database/            # 数据库初始化
    └── config/              # 配置文件
---
```

## 学习日记
- 11.9 学习如何使用vite创建react项目，前端初步写了一点框架，写了导航栏的样式
- 11.10 完善了主页home的样式，blog页面写了部分样式
- 11.10晚 使用docker配置了开发环境，现在可以一键配置开发环境。优化了页面的结构，开始使用组件化开发，将页面拆分成多个组件，方便一键管理，减少无效复用和代码量
