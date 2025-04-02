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
        alert("Регистрация прошла успешно! Теперь вы можете войти в систему.");
        navigate('/login');
      } else {
        const data = await response.json();
        alert("Ошибка регистрации: " + (data.message || response.statusText));
      }
    } catch (error) {
      console.error("Ошибка при регистрации:", error);
      alert("Ошибка регистрации. Попробуйте еще раз позже.");
    }
  };

  return (
    <div className="container mt-5">
      <div className="row justify-content-center">
        <div className="col-md-6">
          <div className="card shadow">
            <div className="card-body">
              <h2 className="card-title text-center mb-4">Регистрация</h2>
              <form onSubmit={handleRegister}>
                <div className="mb-3">
                  <label className="form-label">Username:</label>
                  <input 
                    type="text" 
                    className="form-control" 
                    value={username} 
                    onChange={e => setUsername(e.target.value)} 
                    required 
                  />
                </div>
                <div className="mb-3">
                  <label className="form-label">Email:</label>
                  <input 
                    type="email" 
                    className="form-control" 
                    value={email} 
                    onChange={e => setEmail(e.target.value)} 
                    required 
                  />
                </div>
                <div className="mb-3">
                  <label className="form-label">Пароль:</label>
                  <input 
                    type="password" 
                    className="form-control" 
                    value={password} 
                    onChange={e => setPassword(e.target.value)} 
                    required 
                  />
                </div>
                <button type="submit" className="btn btn-success w-100">Зарегистрироваться</button>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}

export default Register;
