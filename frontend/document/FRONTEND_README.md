# 前端项目文档

欢迎来到前端项目文档！本文档旨在为开发者提供全面的项目信息和使用指南。

## 📚 文档结构

### [项目概述](frontend_documentation.md)
- 项目简介和技术栈
- 项目结构详解
- 组件架构说明
- 样式规范
- 数据流设计
- 构建和部署指南

### [组件API文档](components_api.md)
- 布局组件API（Layout、Header、Footer、Body、PageContainer）
- 内容组件API（Card、mdCard、PageTitle）
- 页面组件API（Home、Blog、BlogDetail、About、Contact）
- 服务组件API（blogService）
- 样式规范和最佳实践

### [开发指南](development_guide.md)
- 环境搭建和项目初始化
- 开发规范和最佳实践
- 调试技巧和性能优化
- 常见问题和解决方案
- 部署指南和扩展开发

## 🚀 快速开始

1. **环境准备**
   ```bash
   # 确保已安装Node.js 18.x或更高版本
   node --version
   ```

2. **项目初始化**
   ```bash
   cd techspace-frontend
   npm install
   npm run dev
   ```

3. **访问项目**
   - 开发服务器地址：`http://localhost:5173`
   - 项目预览地址：`http://localhost:5173`

## 📁 项目结构速览

```
techspace-frontend/
├── document/           # 📖 项目文档（当前目录）
├── src/
│   ├── components/     # 🧩 可复用组件
│   ├── pages/          # 📄 页面组件
│   ├── services/       # 🔌 API服务
│   └── ...
├── public/             # 🌐 静态资源
└── package.json        # 📦 项目配置
```

## 🎯 核心特性

- **现代化技术栈**: React 18 + Vite 5 + CSS Modules
- **响应式设计**: 支持移动端、平板、桌面端
- **模块化架构**: 清晰的组件分层和代码组织
- **性能优化**: 代码分割、懒加载、缓存策略
- **开发友好**: 热更新、ESLint、详细文档

## 🎨 设计系统

### 颜色方案
- **深色文字**: `#1e293b`, `#334155`
- **浅色背景**: `rgba(255, 255, 255, 0.7)`
- **文本阴影**: `rgba(255, 255, 255, 0.8)`
- **渐变效果**: 现代化渐变背景

### 组件风格
- **圆角设计**: `border-radius: 12px`
- **阴影效果**: 多层次阴影系统
- **模糊背景**: `backdrop-filter: blur(10px)`
- **过渡动画**: 平滑的状态切换

## 🔧 开发工具

### 推荐编辑器
- **VS Code** (强烈推荐)
  - ESLint插件
  - Prettier插件
  - Reactjs code snippets

### 浏览器扩展
- **React Developer Tools** - React组件调试
- **CSS Peeper** - CSS样式检查

### 命令行工具
```bash
# 开发模式
npm run dev

# 生产构建
npm run build

# 代码检查
npm run lint

# 预览构建
npm run preview
```

## 📋 开发检查清单

### 创建新组件时
- [ ] 在 `src/components/` 创建组件文件夹
- [ ] 创建 `ComponentName.jsx` 和 `ComponentName.module.css`
- [ ] 遵循组件开发规范
- [ ] 添加必要的注释说明
- [ ] 测试响应式布局

### 创建新页面时
- [ ] 在 `src/pages/` 创建页面文件夹
- [ ] 创建页面组件和样式文件
- [ ] 在路由配置中添加新路由
- [ ] 更新导航菜单（如需要）
- [ ] 测试页面加载性能

### 提交代码前
- [ ] 运行 ESLint 检查
- [ ] 测试所有功能正常
- [ ] 检查响应式布局
- [ ] 验证浏览器兼容性
- [ ] 清理调试代码

## 🐛 常见问题

**Q: 样式不生效？**
A: 检查CSS Modules类名拼写，确认样式文件正确导入

**Q: 组件重新渲染过多？**
A: 使用React.memo、useCallback、useMemo进行优化

**Q: API请求失败？**
A: 检查网络连接，确认API端点，添加错误处理

**Q: 响应式布局问题？**
A: 检查媒体查询断点，使用相对单位，测试不同设备

详细解决方案请参考 [开发指南](development_guide.md)

## 📞 获取帮助

如果在开发过程中遇到问题：

1. **查看文档** - 详细阅读相关文档章节
2. **检查代码** - 确认遵循开发规范
3. **调试工具** - 使用浏览器开发者工具
4. **搜索解决方案** - 查阅相关技术文档和社区

## 🔄 文档更新

本文档会随项目发展持续更新，建议定期查看最新版本。

---

**Happy Coding!** 🎉

*开始你的前端开发之旅吧！*