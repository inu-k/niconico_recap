import { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import { HistoryList, VideoHistoryProps } from '../components/HistoryList';
import { fetchData } from '../functions/utils';

// show history of watch_videos
export default function HistoryPage() {
    const [historyProps, setHisotryProps] = useState<VideoHistoryProps>({ history: [] });
    const [date, setDate] = useState('');

    useEffect(() => {
        fetchData('http://localhost:8088/history')
            .then(data => {
                setHisotryProps({ history: data });
            })
            .catch(error => {
                console.error(error);
                setHisotryProps({ history: [] });
            });
        console.log('useEffect', historyProps);
    }, []);

    const handleHistoryShow = () => {
        fetchData(`http://localhost:8088/history/${date}`)
            .then(data => {
                setHisotryProps({ history: data });
            })
            .catch(error => {
                console.error(error);
                setHisotryProps({ history: [] });
            });
    };

    return (
        <div>
            <h1>History</h1>
            <p> 履歴を表示する日付を選択：
                <input type='date' value={date} onChange={(e) => setDate(e.target.value)} />
                <button onClick={handleHistoryShow}>Show</button>
            </p>
            <HistoryList history={historyProps.history} />
        </div>
    );
}