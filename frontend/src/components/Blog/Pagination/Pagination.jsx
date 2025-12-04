import React from 'react';
import styles from './Pagination.module.css';

const Pagination = ({ currentPage, totalPages, onPageChange, loading = false }) => {
  // 如果总页数小于等于1，不显示分页
  if (totalPages <= 1) {
    return null;
  }

  // 生成页码数组
  const getPageNumbers = () => {
    const pages = [];
    const maxVisiblePages = 5; // 最多显示5个页码
    
    if (totalPages <= maxVisiblePages) {
      // 如果总页数少于等于最大显示页数，显示所有页码
      for (let i = 1; i <= totalPages; i++) {
        pages.push(i);
      }
    } else {
      // 总页数较多时，智能显示页码
      if (currentPage <= 3) {
        // 当前页在前面，显示 1,2,3,4,...,last
        for (let i = 1; i <= 4; i++) {
          pages.push(i);
        }
        pages.push('...');
        pages.push(totalPages);
      } else if (currentPage >= totalPages - 2) {
        // 当前页在后面，显示 1,...,last-3,last-2,last-1,last
        pages.push(1);
        pages.push('...');
        for (let i = totalPages - 3; i <= totalPages; i++) {
          pages.push(i);
        }
      } else {
        // 当前页在中间，显示 1,...,current-1,current,current+1,...,last
        pages.push(1);
        pages.push('...');
        for (let i = currentPage - 1; i <= currentPage + 1; i++) {
          pages.push(i);
        }
        pages.push('...');
        pages.push(totalPages);
      }
    }
    
    return pages;
  };

  const handlePageClick = (page) => {
    if (page !== '...' && page !== currentPage && !loading) {
      onPageChange(page);
    }
  };

  const handlePrevPage = () => {
    if (currentPage > 1 && !loading) {
      onPageChange(currentPage - 1);
    }
  };

  const handleNextPage = () => {
    if (currentPage < totalPages && !loading) {
      onPageChange(currentPage + 1);
    }
  };

  return (
    <div className={styles.pagination}>
      <button
        className={`${styles.pageButton} ${styles.navButton} ${
          currentPage === 1 || loading ? styles.disabled : ''
        }`}
        onClick={handlePrevPage}
        disabled={currentPage === 1 || loading}
      >
        ‹ 上一页
      </button>

      <div className={styles.pageNumbers}>
        {getPageNumbers().map((page, index) => (
          <button
            key={index}
            className={`${styles.pageButton} ${
              page === currentPage ? styles.active : ''
            } ${page === '...' ? styles.ellipsis : ''} ${
              loading ? styles.disabled : ''
            }`}
            onClick={() => handlePageClick(page)}
            disabled={page === '...' || loading}
          >
            {page}
          </button>
        ))}
      </div>

      <button
        className={`${styles.pageButton} ${styles.navButton} ${
          currentPage === totalPages || loading ? styles.disabled : ''
        }`}
        onClick={handleNextPage}
        disabled={currentPage === totalPages || loading}
      >
        下一页 ›
      </button>
    </div>
  );
};

export default Pagination;