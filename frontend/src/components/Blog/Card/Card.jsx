import styles from './Card.module.css'

export default function Card({ title, items }) {
  return (
    <div className={styles.Card}>
      <h1 className={styles.CardTitle}>{title}</h1>
      <div className={styles.CardContent}>
        {items.map((item, index) => (
          <p key={index}>{item}</p>
        ))}
      </div>
    </div>
  );
}