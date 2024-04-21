import { useState } from 'react';
import { fetchData } from '../functions/utils';
import { Summary, SummaryProps } from '../components/Summary';

// show summary of videos
export default function SummaryPage() {
    const [summaryProps, setSummaryProps] = useState<SummaryProps>({ summary: [], watch_count: 0 });
    const [watchCount, setWatchCount] = useState(0);
    const [summary, setSummary] = useState<Summary[]>([]); // [ { tag_name: 'tag', count: 1 }, ... ]
    const [date, setDate] = useState('');
    const [startDate, setStartDate] = useState('');
    const [endDate, setEndDate] = useState('');

    const handleSummaryShowOnDate = () => {
        fetchData(`http://localhost:8088/summary?date=${date}`)
            .then(data => {
                setSummary(data);
            })
            .catch(error => {
                console.error(error);
                setSummary([]);
            });

        fetchData(`http://localhost:8088/history?date=${date}`)
            .then(data => {
                setWatchCount(data.length);
            })
            .catch(error => {
                console.error(error);
                setWatchCount(0);
            });
    };

    const handleSummaryShowOnDuration = () => {
        let params: string[] = [];
        if (startDate) {
            params.push(`startDate=${startDate}`);
        }
        if (endDate) {
            params.push(`endDate=${endDate}`);
        }

        fetchData(`http://localhost:8088/summary?${params.join('&')}`)
            .then(data => {
                setSummary(data);
            })
            .catch(error => {
                console.error(error);
                setSummary([]);
            });

        fetchData(`http://localhost:8088/history?${params.join('&')}`)
            .then(data => {
                setWatchCount(data.length);
            })
            .catch(error => {
                console.error(error);
                setWatchCount(0);
            });
    };

    return (
        <div>
            <h1>Summary</h1>
            <p> サマリーを表示する特定の日付を選択：
                <input type='date' value={date} onChange={(e) => setDate(e.target.value)} />
                <button onClick={handleSummaryShowOnDate}>Show</button>
            </p>
            <p> サマリーを表示する期間を選択：
                <input type='date' value={startDate} onChange={(e) => setStartDate(e.target.value)} />
                〜
                <input type='date' value={endDate} onChange={(e) => setEndDate(e.target.value)} />
                <button onClick={handleSummaryShowOnDuration}>Show</button>
            </p>
            <Summary summary={summary} watch_count={watchCount} />
        </div>
    );
}