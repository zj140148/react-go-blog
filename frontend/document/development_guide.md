# 开发指南

## 环境搭建

### 前置要求
- Node.js (推荐版本: 22.19.0 或更高)
- npm 或 yarn 包管理器
- Git 版本控制

### 项目初始化
```bash
# 克隆项目
git clone [项目地址]

# 进入项目目录
cd techspace-frontend

# 安装依赖
npm install

# 启动开发服务器
npm run dev
```

### 开发服务器
- 默认地址: `http://localhost:5173`
- 热更新: 文件修改后自动刷新
- 代理配置: 可在 `vite.config.js` 中配置API代理
默认配置
```js
proxy: {
      '/api': {
        target: 'http://localhost:5000',  // 后端服务器地址
        changeOrigin: true,
        secure: false,
        rewrite: (path) => path  // 保持路径不变
      }
```

---

## 项目结构详解

### 目录结构说明
```
src/
├── components/     # 可复用组件
│   ├── Layout/     # 布局组件
│   ├── Card/       # 卡片组件
│   └── ...         # 其他组件
├── pages/          # 页面组件
│   ├── Home/       # 首页
│   ├── Blog/       # 博客页面
│   └── ...         # 其他页面
├── services/       # API服务
├── assets/         # 静态资源
├── App.jsx         # 根组件
└── main.jsx        # 应用入口
```

### 文件命名规范
- 组件文件: 使用 PascalCase (如: `Card.jsx`)
- 样式文件: 使用同名 + `.module.css` (如: `Card.module.css`)
- 工具函数: 使用 camelCase (如: `blogService.js`)
- 常量文件: 使用 UPPER_SNAKE_CASE (如: `API_CONSTANTS.js`)

---

## 开发规范

### 组件开发规范

#### 1. 组件结构
```jsx
// 导入语句
import React, { useState, useEffect } from 'react';
import styles from './ComponentName.module.css';

// 组件定义
function ComponentName({ prop1, prop2 }) {
  // 状态定义
  const [state, setState] = useState(initialValue);
  
  // 副作用处理
  useEffect(() => {
    // 副作用逻辑
  }, [dependency]);
  
  // 事件处理函数
  const handleEvent = () => {
    // 事件处理逻辑
  };
  
  // 渲染逻辑
  return (
    <div className={styles.container}>
      {/* JSX内容 */}
    </div>
  );
}

export default ComponentName;
```

#### 2. Props定义
```jsx
// 推荐：使用解构赋值
function ComponentName({ title, items, onClick }) {
  // 组件逻辑
}

// 避免：使用props.xxx
function ComponentName(props) {
  // 不推荐的方式
}
```

#### 3. 状态管理
```jsx
// 推荐：使用函数式更新
setCount(prevCount => prevCount + 1);

// 避免：直接依赖当前状态
setCount(count + 1); // 可能导致更新问题
```

### CSS Modules使用规范

#### 1. 类名命名
```css
/* 推荐：使用驼峰命名法 */
.container {
  padding: 20px;
}

.cardTitle {
  font-size: 18px;
  color: #1e293b;
}

/* 避免：使用连字符 */
.card-title {
  /* 不推荐在JSX中使用 */
}
```

#### 2. 样式组织
```css
/* ComponentName.module.css */

/* 容器样式 */
.container {
  /* 容器相关样式 */
}

/* 标题样式 */
.title {
  /* 标题相关样式 */
}

/* 内容样式 */
.content {
  /* 内容相关样式 */
}

/* 响应式设计 */
@media (max-width: 768px) {
  .container {
    /* 移动端样式 */
  }
}
```

### 响应式设计规范

#### 1. 断点设置
```css
/* 标准断点 */
@media (max-width: 768px) { /* 平板和手机 */ }
@media (max-width: 480px) { /* 手机 */ }
@media (min-width: 769px) { /* 桌面 */ }
```

#### 2. 移动端优先
```css
/* 先写移动端样式 */
.container {
  padding: 10px; /* 移动端 */
}

/* 再写桌面端样式 */
@media (min-width: 769px) {
  .container {
    padding: 20px; /* 桌面端 */
  }
}
```

---

## 调试技巧

### 1. React Developer Tools
- 安装浏览器扩展
- 检查组件树结构
- 查看组件状态和Props
- 分析性能问题

### 2. 控制台调试
```jsx
// 状态调试
console.log('当前状态:', state);

// Props调试
console.log('组件Props:', props);

// 性能调试
console.time('操作耗时');
// 某些操作
console.timeEnd('操作耗时');
```

### 3. 网络请求调试
```jsx
// API调用调试
import { getBlogs } from '../services/blogService';

async function loadData() {
  try {
    console.log('开始获取数据...');
    const data = await getBlogs();
    console.log('获取数据成功:', data);
  } catch (error) {
    console.error('获取数据失败:', error);
  }
}
```

### 4. 样式调试
```css
/* 添加调试边框 */
.debug {
  border: 1px solid red;
}

/* 可视化布局 */
.container {
  background-color: rgba(255, 0, 0, 0.1); /* 半透明背景 */
}
```

---

## 性能优化

### 1. 组件优化
```jsx
// 使用React.memo避免不必要的重新渲染
import React from 'react';

const MemoizedComponent = React.memo(function ComponentName({ prop1, prop2 }) {
  // 组件逻辑
});

// 使用useCallback缓存函数
const handleClick = useCallback(() => {
  // 事件处理逻辑
}, [dependency]);

// 使用useMemo缓存计算结果
const expensiveValue = useMemo(() => {
  return computeExpensiveValue(dependency);
}, [dependency]);
```

### 2. 代码分割
```jsx
// 使用动态导入进行代码分割
import { lazy, Suspense } from 'react';

const LazyComponent = lazy(() => import('./components/LazyComponent'));

function App() {
  return (
    <Suspense fallback={<div>Loading...</div>}>
      <LazyComponent />
    </Suspense>
  );
}
```

### 3. 图片优化
```jsx
// 使用适当的图片格式
// WebP格式提供更好的压缩率
<img src="image.webp" alt="描述" />

// 添加loading属性优化加载
<img src="image.jpg" alt="描述" loading="lazy" />
```

---

## 常见问题和解决方案

### 1. 样式不生效
**问题**: CSS Modules类名不生效
**解决方案**: 
- 检查类名拼写是否正确
- 确认样式文件是否正确导入
- 检查是否有其他样式覆盖

### 2. 组件重新渲染过多
**问题**: 组件频繁重新渲染
**解决方案**:
- 使用React.memo进行优化
- 检查依赖数组是否正确
- 避免在渲染函数中创建新对象

### 3. API请求失败
**问题**: 网络请求报错
**解决方案**:
- 检查网络连接
- 确认API端点是否正确
- 添加错误处理逻辑
- 查看浏览器控制台错误信息

### 4. 样式冲突
**问题**: 组件样式相互影响
**解决方案**:
- 使用CSS Modules避免全局污染
- 检查类名是否重复
- 使用更具体的选择器

### 5. 响应式布局问题
**问题**: 移动端显示异常
**解决方案**:
- 检查媒体查询断点设置
- 使用相对单位（rem, %）
- 测试不同设备尺寸

---

## 部署指南

### 1. 构建项目
```bash
# 生产环境构建
npm run build

# 构建输出目录
dist/
```

### 2. 本地预览
```bash
# 预览构建结果
npm run preview
```

### 3. Docker部署
```bash
# 构建Docker镜像
docker build -t techspace-frontend .

# 运行容器
docker run -p 3000:3000 techspace-frontend
```

### 4. 静态部署
构建后的 `dist/` 目录可以部署到任何静态文件服务器：
- Nginx
- Apache
- GitHub Pages
- Netlify
- Vercel

---

## 扩展开发

### 1. 添加新组件
1. 在 `src/components/` 目录创建新文件夹
2. 创建组件文件和样式文件
3. 遵循组件开发规范
4. 在需要的地方导入使用

### 2. 添加新页面
1. 在 `src/pages/` 目录创建新文件夹
2. 创建页面组件和样式文件
3. 在路由配置中添加新路由
4. 更新导航菜单（如需要）

### 3. 添加新服务
1. 在 `src/services/` 目录创建服务文件
2. 定义API接口函数
3. 添加错误处理逻辑
4. 在组件中导入使用

---

## 相关资源

### 官方文档
- [React官方文档](https://react.dev/)
- [Vite官方文档](https://vitejs.dev/)
- [React Router文档](https://reactrouter.com/)

### 学习资源
- [MDN Web文档](https://developer.mozilla.org/)
- [CSS-Tricks](https://css-tricks.com/)
- [JavaScript.info](https://javascript.info/)

### 工具推荐
- **VS Code**: 推荐编辑器
- **React Developer Tools**: React调试工具
- **ESLint**: 代码质量检查
- **Prettier**: 代码格式化

---

*本指南会根据项目发展持续更新，如有建议或问题请联系。*