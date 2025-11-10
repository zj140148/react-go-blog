import React from 'react';
import styles from './Home.module.css';
import PageContainer from '../../components/PageContainer/PageContainer.jsx';
import PageTitle from '../../components/PageTitle/PageTitle.jsx';
import Body from '../../components/Body/Body.jsx';
import Card from '../../components/Card/Card.jsx';

function Home() {
  return (
    <Body>
      <PageContainer>
        <PageTitle>欢迎来到我的技术博客</PageTitle>
        <p className={styles.pageSubtitle}>
          这里记录了我的技术学习历程、项目经验和心得体会。
          希望通过这个平台与大家分享知识，共同进步。
        </p>

        <div className={styles.grid}>
          <Card
            title="最新文章"
            items={[
              "React Hooks 最佳实践",
              "Go 语言并发编程",
              "现代 CSS 布局技巧"
            ]}
          />

          <Card
            title="技术栈"
            items={[
              "前端：React, TypeScript, Tailwind CSS",
              "后端：Go, Gin, MySQL",
              "工具：Git, Docker, Vite"
            ]}
          />
        </div>



        <div className={styles.buttonContainer}>
          <a href="/blog" className={styles.primaryButton}>
            查看全部文章
          </a>
        </div>
        {/* 插入空白测试 */}
        <div className={styles.spacer}></div>

      </PageContainer>
    </Body>
  );
}

export default Home;