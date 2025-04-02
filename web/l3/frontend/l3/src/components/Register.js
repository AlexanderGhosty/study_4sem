import { useState } from 'react';
import { useNavigate } from 'react-router-dom';

function Register() {
  const [username, setUsername] = useState('');
  const [email, setEmail]     = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  const handleRegister = async (event) => {
    event.preventDefault();

    try {
      const response = await fetch("http://localhost:8080/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username: username, email: email, password: password })
      });

      if (response.ok) {
        // Регистрация прошла успешно, уведомляем пользователя и перенаправляем на страницу логина
        alert("Регистрация прошла успешно! Теперь вы можете войти в систему.");
        navigate('/login');
      } else {
        // Обработка ошибок – читаем сообщение ошибки из ответа (если есть)
        const data = await response.json();
        alert("Ошибка регистрации: " + (data.message || response.statusText));
      }
    } catch (error) {
      console.error("Ошибка при регистрации:", error);
      alert("Ошибка регистрации. Попробуйте еще раз позже.");
    }
  };

  return (
    <div className="register-page">
      <h2>Регистрация</h2>
      <form onSubmit={handleRegister}>
        <div>
          <label>Username:</label><br />
          <input 
            type="text" 
            value={username} 
            onChange={e => setUsername(e.target.value)} 
            required 
          />
        </div>
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
        <button type="submit">Зарегистрироваться</button>
      </form>
    </div>
  );
}

export default Register;
