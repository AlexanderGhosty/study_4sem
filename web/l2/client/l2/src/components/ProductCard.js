import React from 'react';
import styles from './ProductCard.module.css';

function ProductCard({ title, description, price }) {
  return (
    <div className="card h-100">
      <div className="card-body">
        {/* Пример использования CSS-модуля */}
        <h5 className={styles.cardTitle}>{title}</h5>
        <p>{description}</p>
        <p><strong>Цена: {price} руб.</strong></p>
      </div>
    </div>
  );
}

export default ProductCard;
