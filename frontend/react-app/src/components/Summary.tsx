import { Line } from 'react-chartjs-2';
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
} from 'chart.js';

ChartJS.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend
);

const options = {
    maintainAspectRatio: true,
}

export interface Summary {
    tag_name: string;
    count: number;
}

export interface SummaryProps {
    summary: Summary[];
    watch_count: number;
}

export function Summary({ summary, watch_count }: SummaryProps) {
    return (
        <div>

            {/* <div style={{ width: '50%', height: '50%' }}>
                <Line
                    height={10}
                    width={20}
                    data={{
                        labels: ['1', '2', '3', '4', '5'],
                        datasets: [{
                            label: '視聴回数',
                            data: [12, 19, 3, 5, 2],
                            backgroundColor: 'rgba(255, 99, 132, 0.8)',
                        }]
                    }}
                    options={options}
                />
            </div> */}

            <p>総視聴回数：{watch_count}回</p>
            <p>
                {summary.map((summary, index) => {
                    let percentage = (watch_count === 0) ? 0 : (summary.count / watch_count * 100).toFixed(2);
                    return (
                        <div className="summary-item-container">
                            <div className="summary-item-content">
                                <p>{summary.tag_name}: {summary.count} 回 ({percentage}%)</p>
                            </div>
                        </div>);
                })}
            </p>
        </div>);
}