import React from 'react';
import styles from './Blog.module.css';
import PageContainer from '../../components/PageContainer/PageContainer.jsx';
import PageTitle from '../../components/PageTitle/PageTitle.jsx';
import Card from '../../components/Card/Card.jsx';
import Body from '../../components/Body/Body.jsx';
function Blog() {
  return (
    <Body>
      <PageContainer>
        <PageTitle>博客文章</PageTitle>

        <div className={styles.grid}>
          <Card
            title="文章1：React入门指南"
            items={[
              "这是一篇关于React入门的指南，介绍了React的基本概念、组件化开发、JSX语法等。",
              "文章适合初学者，帮助他们快速上手React开发。"
            ]}
          />
          <Card
            title="文章2：go入门指南"
            items={[
              "这是一篇关于Go语言入门的指南，介绍了Go的基本语法、并发编程、标准库等。",
              "文章适合初学者，帮助他们快速上手Go语言开发。"
            ]}
          />
        </div>
      </PageContainer>

    </Body>
  )
}

export default Blog;