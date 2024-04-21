import './App.css';
import { Routes, Route, Link } from 'react-router-dom';
import HistoryPage from './pages/HistoryPage';
import VideoInfoPage from './pages/VideoInfoPage';


function App() {
  return (
    <div className="App">
      <header className="App-header">
        <p>
          niconico-recap
        </p>

        <Link className='header-link' to="/history">History</Link>
        <Link className='header-link' to="/summary">Summary</Link>

      </header>

      <Routes>
        <Route path='/history' element={<HistoryPage />} />
        <Route path='/summary' element={<div>Summary</div>} />
        <Route path='/videos/:videoId' element={<VideoInfoPage />} />
      </Routes>
    </div>
  );
}

export default App;
