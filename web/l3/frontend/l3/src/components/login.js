import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

function Login() {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  const handleLogin = async (event) => {
    event.preventDefault(); // отменяем перезагрузку формы
    
    // Формируем запрос к /login
    try {
      const response = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({ email: email, password: password })
      });
      if (response.ok) {
        const data = await response.json();
        const token = data.token;
        // Сохраняем токен в localStorage
        localStorage.setItem('token', token);
        // Перенаправляем пользователя на защищенную страницу (например, Dashboard)
        navigate('/dashboard'); 
      } else {
        // Неверные данные или ошибка – можно показать сообщение
        alert("Ошибка входа: " + response.statusText);
      }
    } catch (err) {
      console.error("Login error:", err);
      alert("Ошибка запроса, попробуйте позже");
    }
  };

  return (
    <div className="login-page">
      <h2>Вход</h2>
      <form onSubmit={handleLogin}>
        <div>
          <label>Email:</label><br />
          <input 
            type="email" 
            value={email} 
            onChange={e => setEmail(e.target.value)} 
            required 
          />
        </div>
        <div>
          <label>Пароль:</label><br />
          <input 
            type="password" 
            value={password} 
            onChange={e => setPassword(e.target.value)} 
            required 
          />
        </div>
        <button type="submit">Войти</button>
      </form>
    </div>
  );
}

export default Login;
