import React from 'react';
import ProductCard from './ProductCard';

function ProductList({ products }) {
  return (
    <div className="row">
      {products.map((product) => (
        <div className="col-md-4 mb-4" key={product.id}>
          <ProductCard 
            title={product.title}
            description={product.description}
            price={product.price}
          />
        </div>
      ))}
    </div>
  );
}

export default ProductList;
