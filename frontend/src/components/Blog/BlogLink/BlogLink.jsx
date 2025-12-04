import { Link } from 'react-router-dom';
import styles from './BlogLink.module.css';


export default function BlogLink({ children, id }) {
  return (
    <Link key={id} to={`/blog/${id}`} className={styles.BlogLink}>
      {children}
    </Link>
  );
}