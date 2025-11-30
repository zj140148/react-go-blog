# 组件API文档

## 布局组件

### Layout组件
**文件路径**: `src/components/Layout/Layout.jsx`

页面布局容器组件，提供整体页面结构。

**Props**: 无

**使用示例**:
```jsx
import Layout from '../components/Layout/Layout';

function App() {
  return (
    <Layout>
      {/* 页面内容 */}
    </Layout>
  );
}
```

---

### Header组件
**文件路径**: `src/components/Header/Header.jsx`

网站头部导航组件，包含logo和导航菜单。

**Props**: 无

**使用示例**:
```jsx
import Header from '../components/Header/Header';

function Home() {
  return (
    <div>
      <Header />
      {/* 其他内容 */}
    </div>
  );
}
```

---

### Footer组件
**文件路径**: `src/components/Footer/Footer.jsx`

网站页脚组件，显示版权信息。

**Props**: 无

**使用示例**:
```jsx
import Footer from '../components/Footer/Footer';

function Home() {
  return (
    <div>
      {/* 页面内容 */}
      <Footer />
    </div>
  );
}
```

---

### Body组件
**文件路径**: `src/components/Body/Body.jsx`

主体内容容器组件，用于包裹页面主要内容。

**Props**: 无

**使用示例**:
```jsx
import Body from '../components/Body/Body';

function Home() {
  return (
    <Body>
      {/* 主体内容 */}
    </Body>
  );
}
```

**Body样式:**
```css
.Body {
    padding: 2rem 0;
    width: 100%;
    display: flex;
    flex-direction: column;
    align-items: center;
    /* 水平居中 */
    justify-content: flex-start;
    /* 垂直方向靠上对齐 */
    background-color: #1e8a73;
    min-height: calc(100vh - 80px);
    /* 减去导航栏高度 */
    box-sizing: border-box;
    background-color: transparent;
    /* 或 rgba(255,255,255,0.0) */
    min-height: calc(100vh - 80px);
    box-sizing: border-box;
}
```

---

### PageContainer组件
**文件路径**: `src/components/PageContainer/PageContainer.jsx`

页面内容容器组件，提供统一的页面内容样式。

**Props**: 无

**样式特性**:
- 背景颜色: `rgba(255, 255, 255, 0.7)`
- 圆角边框: `12px`
- 阴影效果
- 模糊背景: `backdrop-filter: blur(10px)`(暂时注释，待优化)
```css
.PageContainer {
  border-radius: 24px;
  padding: 3rem;
  margin: 0 auto;
  width: 100%;
  max-width: 80%;
  border: 1px solid rgba(255, 255, 255, 0.18);
  box-sizing: border-box;
  background: rgba(255, 255, 255, 0.7);
  /* backdrop-filter: blur(12px) saturate(180%); */
  /* -webkit-backdrop-filter: blur(12px) saturate(180%); */
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.3), inset 0 0 20px rgba(255, 255, 255, 0.05);
  position: relative;
  overflow: hidden;
}
```
**使用示例**:
```jsx
import PageContainer from '../components/PageContainer/PageContainer';

function Home() {
  return (
    <PageContainer>
      <h1>页面标题</h1>
      <p>页面内容</p>
    </PageContainer>
  );
}
```

---

## 内容组件

### Card组件
**文件路径**: `src/components/Card/Card.jsx`

通用卡片组件，用于展示内容块。

**Props**:
| 属性名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| title | string | 是 | 卡片标题 |
| items | array | 是 | 卡片内容项数组 |

**使用示例**:
```jsx
import Card from '../components/Card/Card';

function Home() {
  const skills = ["React", "Vite", "CSS Modules"];
  
  return (
    <Card title="技术栈" items={skills} />
  );
}
```

**样式特性**:
- 标题颜色: `#f8fafc`
- 内容颜色: `#e2e8f0`
- 背景颜色: `rgba(20, 53, 50, 0.85)`
- hover背景颜色: `rgba(22, 54, 51, 0.9)`
```css
.Card {
  background: rgba(20, 53, 50, 0.85);
  border-radius: 16px;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1), 0 2px 4px -1px rgba(0, 0, 0, 0.06);
  padding: 2rem;
  margin-bottom: 1.5rem;
  border: 1px solid rgba(148, 163, 184, 0.2);
  transition: all 0.4s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  text-align: left;
  outline: none;
  backdrop-filter: blur(10px);
  -webkit-backdrop-filter: blur(10px);
}

.Card:hover {
  transform: translateY(-8px) scale(1.02);
  background: rgba(22, 54, 51, 0.9);
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.4), 0 0 0 1px rgba(94, 234, 212, 0.3);
}

/* 卡片标题 */
.CardTitle {
  font-size: 1.5rem;
  font-weight: 700;
  color: #f8fafc;
  margin-bottom: 1rem;

}

/* 卡片内容 */
.CardContent {
  color: #e2e8f0;
  line-height: 1.7;
}

.CardContent p {
  margin-bottom: 0.5rem;
}
```
---

### mdCard组件
**文件路径**: `src/components/mdCard/mdCard.jsx`

Markdown内容卡片组件，用于渲染Markdown格式的内容。样式与Card组件相同。

**Props**: 无

**使用示例**:
```jsx
import MdCard from '../components/mdCard/mdCard';

function BlogDetail() {
  return (
    <MdCard>
      {/* Markdown内容 */}
    </MdCard>
  );
}
```

---

### PageTitle组件
**文件路径**: `src/components/PageTitle/PageTitle.jsx`

页面标题组件，用于显示页面主标题。

**Props**: 无

**使用示例**:
```jsx
import PageTitle from '../components/PageTitle/PageTitle';

function About() {
  return (
    <div>
      <PageTitle />
      {/* 其他内容 */}
    </div>
  );
}
```

**样式特性**:
- 标题颜色: `#1e293b`
- 白色文本阴影: `rgba(255, 255, 255, 0.8)`
```css
.PageTitle {
  text-shadow: 0 2px 4px rgba(255, 255, 255, 0.8);
  font-size: 3rem;
  font-weight: 800;
  color: #1e293b;
  margin-bottom: 1.5rem;
  text-align: center;
  letter-spacing: -0.5px;
}
```

---

## 页面组件

### Home组件
**文件路径**: `src/pages/Home/Home.jsx`

首页组件，展示网站概览和主要信息。

**Props**: 无

**组件结构**:
- PageTitle组件
- 页面副标题
- 技术栈Card组件
- 项目介绍Card组件

**样式特性**:
- 副标题颜色: `#334155`
- 白色文本阴影: `rgba(255, 255, 255, 0.8)`

---

### Blog组件
**文件路径**: `src/pages/Blog/Blog.jsx`

博客列表页面组件，展示所有博客文章。

**Props**: 无

**功能特性**:
- 获取博客列表数据
- 渲染博客文章列表
- 支持文章分类显示

---

### BlogDetail组件
**文件路径**: `src/pages/BlogDetail/BlogDetail.jsx`

博客详情页面组件，展示单篇博客文章的详细内容。

**Props**: 无

**功能特性**:
- 根据ID获取博客详情
- 渲染Markdown格式的文章内容
- 显示文章元信息（标题、日期等）

---

### About组件
**文件路径**: `src/pages/About/About.jsx`

关于页面组件，展示个人或项目介绍信息。

**Props**: 无

---

### Contact组件
**文件路径**: `src/pages/Contact/Contact.jsx`

联系页面组件，包含联系表单。

**Props**: 无

**功能特性**:
- 联系表单（姓名、邮箱、留言）
- 表单验证
- 提交处理

**表单字段**:
| 字段名 | 类型 | 说明 |
|--------|------|------|
| name | text | 用户姓名 |
| email | email | 用户邮箱 |
| message | textarea | 留言内容 |

**样式特性**:
- 表单背景: 半透明白色
- 圆角边框: `12px`
- 阴影效果
- 现代化输入框设计
- 渐变背景提交按钮

---

## 服务组件

### blogService
**文件路径**: `src/services/blogService.js`

博客相关的API服务，提供数据获取功能。

**API方法**:

#### getBlogs()
获取所有博客文章列表。

**返回值**: Promise<Blog[]>

**使用示例**:
```jsx
import { getBlogs } from '../services/blogService';

async function loadBlogs() {
  try {
    const blogs = await getBlogs();
    console.log(blogs);
  } catch (error) {
    console.error('获取博客失败:', error);
  }
}
```

#### getBlogById(id)
根据ID获取单篇博客文章详情。

**参数**:
| 参数名 | 类型 | 说明 |
|--------|------|------|
| id | string/number | 博客文章ID |

**返回值**: Promise<Blog>

**使用示例**:
```jsx
import { getBlogById } from '../services/blogService';

async function loadBlogDetail(blogId) {
  try {
    const blog = await getBlogById(blogId);
    console.log(blog);
  } catch (error) {
    console.error('获取博客详情失败:', error);
  }
}
```

---

## 样式规范

### CSS Modules命名规范
- 组件样式类名使用驼峰命名法
- 样式文件与组件文件同名
- 使用CSS Modules避免样式冲突

### 响应式设计
- 移动端优先原则
- 使用现代CSS布局技术（Flexbox、Grid）
- 支持各种屏幕尺寸

### 颜色方案
- 深色文字: `#1e293b`, `#334155`
- 浅色背景: `rgba(255, 255, 255, 0.7)`
- 文本阴影: `rgba(255, 255, 255, 0.8)`
- 渐变背景: 使用CSS渐变创建现代化效果

---

## 最佳实践

### 组件开发
1. 保持组件职责单一
2. 合理使用Props传递数据
3. 添加必要的错误处理
4. 编写清晰的注释说明

### 状态管理
1. 优先使用React Hooks
2. 合理使用useState和useEffect
3. 避免不必要的重新渲染
4. 考虑使用Context API进行跨组件状态管理

### 性能优化
1. 使用React.memo优化组件渲染
2. 合理使用useCallback和useMemo
3. 实现代码分割和懒加载
4. 优化图片和其他静态资源

---

*本文档会随项目发展持续更新，如有疑问请参考具体组件实现或联系我。*