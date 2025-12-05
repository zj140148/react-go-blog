import React, { useEffect, useState } from 'react';
import ReactMarkdown from 'react-markdown';
import remarkGfm from 'remark-gfm';
import styles from './About.module.css';
import PageContainer from '../../components/Layout/PageContainer/PageContainer.jsx';
import PageTitle from '../../components/Layout/PageTitle/PageTitle.jsx';
import MdCard from '../../components/About/mdCard/mdCard.jsx';
import PageWrapper from '../../components/Layout/PageWrapper/PageWrapper.jsx';

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || '/api';
function About() {
  const [mdContent, setMdContent] = useState('');

  useEffect(() => {
    fetch(API_BASE_URL + '/about-md') // 后端接口
      .then(res => res.json())
      .then(data => {
        if (data.success && data.data && data.data.content) {
          setMdContent(data.data.content);
        } else {
          setMdContent('加载失败');
        }
      })
      .catch(err => {
        console.error('Failed to fetch markdown:', err);
        setMdContent('加载失败');
      });
  }, []);

  return (
    <PageWrapper>
      <div className={styles.about}>
        <PageContainer>
          <PageTitle>关于我</PageTitle>
          <MdCard>
            <div className="markdown-content">
            
              <ReactMarkdown remarkPlugins={[remarkGfm]}>
                {mdContent || '加载中...'}
              </ReactMarkdown>
            </div>
          </MdCard>
        </PageContainer>
      </div>
    </PageWrapper>
  );
}

export default About;