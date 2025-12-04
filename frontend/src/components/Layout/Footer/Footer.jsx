import styles from './Footer.module.css';
import { FaGithub } from "react-icons/fa"

function Footer() {
  return (
    <footer className={styles.footer}>
      <div className={styles['footer-content']}>
        <p>
          &copy; {new Date().getFullYear()} Wh1stle 的小站. All rights reserved.
        </p>
        <p>
          <a href="https://github.com/Wh1stle05/react-go-blog" target="_blank" rel="noopener noreferrer">
            <FaGithub size={16} />
            项目地址
          </a>
        </p>
      </div>
    </footer>
  );
}

export default Footer;