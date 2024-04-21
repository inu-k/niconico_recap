import './App.css';
import { Routes, Route, Link } from 'react-router-dom';
import HistoryPage from './pages/HistoryPage';
import VideoInfoPage from './pages/VideoInfoPage';


function App() {
  return (
    <div className="App">
      <header className="App-header">
        <p>
          hi
        </p>

        <Link to="/history">History</Link>

      </header>

      <Routes>
        <Route path='/history' element={<HistoryPage />} />
        <Route path='/videos/:videoId' element={<VideoInfoPage />} />
      </Routes>
    </div>
  );
}

export default App;
