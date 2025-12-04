import React, { useState, useEffect, use } from 'react';
import styles from './Blog.module.css';
import PageContainer from '../../components/Layout/PageContainer/PageContainer.jsx';
import PageTitle from '../../components/Layout/PageTitle/PageTitle.jsx';
import Card from '../../components/Blog/Card/Card.jsx';
import Body from '../../components/Layout/Body/Body.jsx';
import Pagination from '../../components/Blog/Pagination/Pagination.jsx';
import { getBlogsPaginated } from '../../services/blogService.js';
import PageWrapper from '../../components/Layout/PageWrapper/PageWrapper.jsx';
import BlogLink from '../../components/Blog/BlogLink/BlogLink.jsx';
function Blog() {
  const [blogs, setBlogs] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(0);
  const [total, setTotal] = useState(0);
  const pageSize = 6; // 每页显示6个卡片


  const fetchBlogs = async (page = 1) => {
    try {
      setLoading(true);
      const response = await getBlogsPaginated(page, pageSize);
      setBlogs(response.data);
      setCurrentPage(response.page);
      setTotalPages(response.total_pages);
      setTotal(response.total);
      setError(null);
    } catch (err) {
      setError(err.message);
      setBlogs([]);
    } finally {
      setLoading(false);
    }
  };


  useEffect(() => {
    fetchBlogs(currentPage);
  }, [currentPage]);

  if (loading) {
    return (
      <PageWrapper>
        <Body>
          <PageContainer>
            <PageTitle>博客文章</PageTitle>
            <div>加载中...</div>
          </PageContainer>
        </Body>
      </PageWrapper>
    );
  }

  if (error) {
    return (
      <PageWrapper>
        <Body>
          <PageContainer>
            <PageTitle>博客文章</PageTitle>
            <div>错误: {error}</div>
          </PageContainer>
        </Body>
      </PageWrapper>
    );
  }

  return (
    <PageWrapper key={currentPage}>
      <Body>
        <PageContainer>
          <PageTitle>博客文章</PageTitle>

          <div className={styles.grid}>
            {blogs.map((blog) => (
              <BlogLink id={blog.id}>
                <Card
                  title={blog.title}
                  items={[
                    blog.summary,
                    `作者: ${blog.author} | 浏览量: ${blog.view_count}`
                  ]}
                />
              </BlogLink>
            ))}
          </div>

          {/* 分页组件 */}
          {!loading && !error && blogs.length > 0 && (
            <Pagination
              currentPage={currentPage}
              totalPages={totalPages}
              onPageChange={setCurrentPage}
              loading={loading}
            />
          )}

          {/* 空状态 */}
          {!loading && !error && blogs.length === 0 && (
            <div className={styles.emptyState}>
              暂无博客文章
            </div>
          )}
        </PageContainer>
      </Body>
    </PageWrapper>
  );
}

export default Blog;