import React, { useState, useEffect } from 'react';
import Header from './components/Header';
import ProductList from './components/ProductList';

function App() {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    fetch('http://localhost:8080/products')
      .then(response => response.json())
      .then(data => {
        setProducts(data); 
      })
      .catch(error => {
        console.error('Ошибка при загрузке:', error);
      });
  }, []);

  return (
    <div>
      <Header />
      <div className="container mt-4">
        <h1>Список продуктов</h1>
        {/* Передаём products в компонент ProductList через props */}
        <ProductList products={products} />
      </div>
    </div>
  );
}

export default App;
