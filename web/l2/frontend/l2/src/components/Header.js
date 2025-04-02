import React from 'react';
import 'bootstrap/dist/css/bootstrap.min.css';

function Header() {
  return (
    <nav className="navbar navbar-expand-lg navbar-light bg-light shadow">
      <div className="container">
        <a className="navbar-brand" href="/">
          My Store
        </a>
      </div>
    </nav>
  );
}

export default Header;
