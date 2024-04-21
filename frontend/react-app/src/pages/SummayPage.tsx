import { useState } from 'react';
import { fetchData } from '../functions/utils';
import { Summary, SummaryProps } from '../components/Summary';

// show summary of videos
export default function SummaryPage() {
    const [summaryProps, setSummaryProps] = useState<SummaryProps>({ summary: [] });
    const [date, setDate] = useState('');
    const [startDate, setStartDate] = useState('');
    const [endDate, setEndDate] = useState('');

    const handleSummaryShowOnDate = () => {
        fetchData(`http://localhost:8088/summary?date=${date}`)
            .then(data => {
                setSummaryProps({ summary: data });
            })
            .catch(error => {
                console.error(error);
                setSummaryProps({ summary: [] });
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
                setSummaryProps({ summary: data });
            })
            .catch(error => {
                console.error(error);
                setSummaryProps({ summary: [] });
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
            <Summary summary={summaryProps.summary} />
        </div>
    );
}