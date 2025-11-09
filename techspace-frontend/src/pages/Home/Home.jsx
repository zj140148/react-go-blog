import React from 'react';
import styles from './Home.module.css';
function Home() {
  return (
    <div className={styles.home}>
      <div className={styles.pageContainer}>
        <h1 className={styles.pageTitle}>欢迎来到我的技术博客</h1>
        <p className={styles.pageSubtitle}>
          这里记录了我的技术学习历程、项目经验和心得体会。
          希望通过这个平台与大家分享知识，共同进步。
        </p>
        
        <div className={styles.grid}>
          <div className={styles.card}>
            <h3 className={styles.cardTitle}>最新文章</h3>
            <div className={styles.cardContent}>
              <p>React Hooks 最佳实践</p>
              <p>Go 语言并发编程</p>
              <p>现代 CSS 布局技巧</p>
            </div>
          </div>
          
          <div className={styles.card}>
            <h3 className={styles.cardTitle}>技术栈</h3>
            <div className={styles.cardContent}>
              <p>前端：React, TypeScript, Tailwind CSS</p>
              <p>后端：Go, Gin, MySQL</p>
              <p>工具：Git, Docker, Vite</p>
            </div>
          </div>
        </div>

        
        
        <div className={styles.buttonContainer}>
          <a href="/blog" className={styles.primaryButton}>
            查看全部文章
          </a>
        </div>
        {/* 插入空白测试 */}
        <div className={styles.spacer}></div>
      </div>
    </div>
  );
}

export default Home;