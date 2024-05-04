import { useState, useEffect } from 'react';
import { useLocation } from 'react-router-dom';
import { HistoryList, VideoHistoryProps } from '../components/HistoryList';
import { fetchData } from '../functions/utils';

// show history of watch_videos
export default function HistoryPage() {
    const [historyProps, setHisotryProps] = useState<VideoHistoryProps>({ history: [] });
    const [date, setDate] = useState('');
    const [startDate, setStartDate] = useState('');
    const [endDate, setEndDate] = useState('');

    // 1週間前
    const today = new Date();
    const lastWeek = new Date(today.getFullYear(), today.getMonth(), today.getDate() - 7);

    useEffect(() => {
        fetchData(`http://localhost:8088/history?startDate=${lastWeek.toISOString().split('T')[0]}`)
            .then(data => {
                setHisotryProps({ history: data });
            })
            .catch(error => {
                console.error(error);
                setHisotryProps({ history: [] });
            });
    }, []);

    const handleHistoryShowOnDate = () => {
        fetchData(`http://localhost:8088/history?date=${date}`)
            .then(data => {
                setHisotryProps({ history: data });
            })
            .catch(error => {
                console.error(error);
                setHisotryProps({ history: [] });
            });
    };

    const handleHistoryShowOnDuration = () => {
        let params: string[] = [];
        if (startDate) {
            params.push(`startDate=${startDate}`);
        }
        if (endDate) {
            params.push(`endDate=${endDate}`);
        }

        fetchData(`http://localhost:8088/history?${params.join('&')}`)
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
            <p> 履歴を表示する特定の日付を選択：
                <input type='date' value={date} onChange={(e) => setDate(e.target.value)} />
                <button onClick={handleHistoryShowOnDate}>Show</button>
            </p>
            <p> 履歴を表示する期間を選択：
                <input type='date' value={startDate} onChange={(e) => setStartDate(e.target.value)} />
                〜
                <input type='date' value={endDate} onChange={(e) => setEndDate(e.target.value)} />
                <button onClick={handleHistoryShowOnDuration}>Show</button>
            </p>
            <HistoryList history={historyProps.history} />
        </div>
    );
}