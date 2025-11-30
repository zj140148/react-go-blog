# TechSpace 前端项目文档

## 项目概述

TechSpace 是一个现代化的技术博客平台前端项目，采用 React + Vite 技术栈构建，提供响应式的用户界面和良好的用户体验。

## 技术栈

### 核心框架
- **React 18** - 前端UI框架
- **Vite 5** - 构建工具和开发服务器
- **JavaScript (ES6+)** - 编程语言

### 路由管理
- **React Router DOM 6** - 单页应用路由管理

### 样式方案
- **CSS Modules** - 模块化CSS方案
- **现代CSS特性** - Flexbox, Grid, CSS变量等

### 数据交互
- **Axios** - HTTP客户端
- **React Hooks** - 状态管理和副作用处理

### 内容渲染
- **React Markdown** - Markdown内容渲染

### 开发工具
- **ESLint** - 代码质量检查
- **Vite** - 热更新和构建优化

## 项目结构

```
techspace-frontend/
├── public/                    # 静态资源
│   ├── images/               # 图片资源
│   └── vite.svg             # Vite图标
├── src/                      # 源代码目录
│   ├── assets/              # 静态资源
│   │   └── react.svg        # React图标
│   ├── components/          # 可复用组件
│   │   ├── Body/            # 主体内容组件
│   │   │   ├── Body.jsx
│   │   │   └── Body.module.css
│   │   ├── Card/            # 卡片组件
│   │   │   ├── Card.jsx
│   │   │   └── Card.module.css
│   │   ├── Footer/          # 页脚组件
│   │   │   ├── Footer.jsx
│   │   │   └── Footer.nodule.jsx
│   │   ├── Header/          # 头部组件
│   │   │   ├── Header.jsx
│   │   │   └── Header.module.css
│   │   ├── Layout/          # 布局组件
│   │   │   ├── Layout.jsx
│   │   │   └── Layout.module.css
│   │   ├── PageContainer/   # 页面容器组件
│   │   │   ├── PageContainer.jsx
│   │   │   └── PageContainer.module.css
│   │   ├── PageTitle/       # 页面标题组件
│   │   │   ├── PageTitle.jsx
│   │   │   └── PageTitle.module.css
│   │   └── mdCard/          # Markdown卡片组件
│   │       ├── mdCard.jsx
│   │       └── mdCard.module.css
│   ├── pages/               # 页面组件
│   │   ├── About/           # 关于页面
│   │   │   ├── About.jsx
│   │   │   └── About.module.css
│   │   ├── Blog/            # 博客列表页面
│   │   │   ├── Blog.jsx
│   │   │   └── Blog.module.css
│   │   ├── BlogDetail/      # 博客详情页面
│   │   │   ├── BlogDetail.jsx
│   │   │   └── BlogDetail.module.css
│   │   ├── Contact/         # 联系页面
│   │   │   ├── Contact.jsx
│   │   │   └── Contact.module.css
│   │   └── Home/            # 首页
│   │       ├── Home.jsx
│   │       └── Home.module.css
│   ├── services/            # API服务
│   │   └── blogService.js   # 博客相关API
│   ├── App.jsx              # 根组件
│   ├── App.css              # 全局样式
│   ├── index.css            # 入口样式
│   └── main.jsx             # 应用入口
├── document/                # 项目文档
├── package.json             # 项目配置
├── vite.config.js           # Vite配置
├── eslint.config.js         # ESLint配置
├── Dockerfile               # Docker配置
└── README.md                # 项目说明
```

## 组件架构

### 布局组件
- **Layout** - 整体页面布局框架
- **Header** - 网站头部导航
- **Footer** - 网站页脚信息
- **Body** - 主体内容容器
- **PageContainer** - 页面内容容器

### 内容组件
- **Card** - 通用卡片组件，用于展示内容块
- **mdCard** - Markdown内容卡片
- **PageTitle** - 页面标题组件

### 页面组件
- **Home** - 首页，展示网站概览
- **Blog** - 博客列表页面
- **BlogDetail** - 博客详情页面
- **About** - 关于页面
- **Contact** - 联系页面，包含表单

## 样式规范

### CSS Modules使用
每个组件都有对应的 `.module.css` 文件，采用模块化CSS方案：
- 类名使用驼峰命名法
- 避免全局样式污染
- 样式作用域限定在组件内部

### 颜色方案
- 深色文字：`#1e293b`, `#334155`
- 浅色背景：`rgba(255, 255, 255, 0.7)`
- 文本阴影：`rgba(255, 255, 255, 0.8)`

### 响应式设计
- 移动端优先原则
- 使用现代CSS布局技术
- 支持各种屏幕尺寸

## 数据流

### API服务
- **blogService.js** - 提供博客相关的API调用
- 使用Axios进行HTTP请求
- 统一的错误处理机制

### 状态管理
- 使用React Hooks进行状态管理
- 组件内部状态使用useState
- 副作用处理使用useEffect

## 开发规范

### 组件开发
1. 每个组件独立文件夹
2. 包含组件文件和样式文件
3. 使用函数组件和Hooks
4. 添加必要的注释说明

### 代码风格
1. 使用ESLint进行代码检查
2. 遵循React最佳实践
3. 保持代码简洁可读
4. 适当的错误处理

## 构建和部署

### 开发环境
```bash
npm run dev      # 启动开发服务器
npm run build    # 构建生产版本
npm run preview  # 预览构建结果
```

### 生产部署
- 使用Vite进行构建优化
- 支持Docker容器化部署
- 静态资源优化处理

## 后续优化方向

1. **TypeScript迁移** - 提升代码类型安全性
2. **状态管理优化** - 考虑使用Context API或Redux
3. **性能优化** - 代码分割、懒加载等
4. **测试覆盖** - 添加单元测试和集成测试
5. **文档完善** - 使用Storybook进行组件文档管理

---

*本文档基于项目当前状态编写，后续会根据项目发展持续更新。*