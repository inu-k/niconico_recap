import './App.css';
import { Routes, Route, Link } from 'react-router-dom';
import HistoryPage from './pages/HistoryPage';
import VideoInfoPage from './pages/VideoInfoPage';
import SummaryPage from './pages/SummayPage';
import VideoSearchPage from './pages/VideoSearchPage';


function App() {
  return (
    <div className="App">
      <header className="App-header">
        <div className='header-title'>niconico_recap</div>

        <div className='header-links'>
          <Link className='header-link' to="/history">視聴履歴</Link>
          <Link className='header-link' to="/summary">サマリー</Link>
          <Link className='header-link' to="/search">動画を検索</Link>
        </div>

      </header>

      <Routes>
        <Route path='/history' element={<HistoryPage />} />
        <Route path='/summary' element={<SummaryPage />} />
        <Route path='/videos/:videoId' element={<VideoInfoPage />} />
        <Route path='/search' element={<VideoSearchPage />} />
      </Routes>
    </div>
  );
}

export default App;
