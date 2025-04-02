import { useNavigate } from 'react-router-dom';

function Dashboard() {
  const navigate = useNavigate();

  const handleLogout = () => {
    localStorage.removeItem('token');             
    navigate('/login');                           
  };

  return (
    <div className="container mt-5">
      <div className="card shadow">
        <div className="card-header">
          <h1>Добро пожаловать, пользователь!</h1>
        </div>
        <div className="card-body">
          <p>Это ваша личная панель управления. Здесь можно разместить полезную информацию или действия.</p>
          <button onClick={handleLogout} className="btn btn-danger">Выйти</button>
        </div>
      </div>
    </div>
  );
}

export default Dashboard;
