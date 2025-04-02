// src/components/ProductCard.js
import React from 'react';
import styles from './ProductCard.module.css';

function ProductCard({ title, description, price }) {
  return (
    <div className={`card h-100 shadow-sm ${styles.card}`}>
      <div className={`card-body ${styles.cardBody}`}>
        <h5 className={styles.cardTitle}>{title}</h5>
        <p className="card-text">{description}</p>
        <p className="card-text">
          <strong>Цена: {price} руб.</strong>
        </p>
        <button className={styles.btnBuy}>Купить</button>
      </div>
    </div>
  );
}

export default ProductCard;
