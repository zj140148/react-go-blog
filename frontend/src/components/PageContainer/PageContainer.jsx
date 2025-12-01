import styles from './PageContainer.module.css'

export default function PageContainer({ children }) {
  return (
    <div className={styles.PageContainer}>
    {children}   {/* 每个页面不同的部分 */}
    </div>
  );
}