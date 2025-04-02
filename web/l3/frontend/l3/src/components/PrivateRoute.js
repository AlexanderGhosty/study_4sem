import { Navigate } from 'react-router-dom';

function PrivateRoute({ children }) {
  const token = localStorage.getItem('token');
  if (!token) {
    // Если токена нет, перенаправляем на страницу логина
    return <Navigate to="/login" replace />;
  }
  // Если токен есть, рендерим дочерний компонент
  return children;
}

export default PrivateRoute;
