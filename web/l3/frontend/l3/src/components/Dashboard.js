import { useNavigate } from 'react-router-dom';

function Dashboard() {
  const navigate = useNavigate();

  const handleLogout = () => {
    localStorage.removeItem('token');             
    navigate('/login');                           
  };

  return (
    <div>
      <h1>Добро пожаловать, пользователь!</h1>
      {/* ... содержимое личного кабинета ... */}
      <button onClick={handleLogout}>Выйти</button>
    </div>
  );
}

export default Dashboard;
