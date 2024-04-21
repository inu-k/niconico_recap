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