import React from 'react';
import styles from './Blog.module.css';
function Blog() {
  return (
    <div className={styles.blog}>
      <h1 className={styles['blog-title']}>博客文章</h1>
      <div className={styles['blog-content']}>
        <ul>
          <li>文章1：React入门指南</li>
          <li>文章2：Go语言基础</li>
          <li>文章3：项目开发经验分享</li>
        </ul>
        <div className={styles.spacer}></div>
      </div>
    </div>
  );
}

export default Blog;