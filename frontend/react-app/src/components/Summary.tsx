export interface Summary {
    tag_name: string;
    count: number;
}

export interface SummaryProps {
    summary: Summary[];
}

export function Summary({ summary }: SummaryProps) {
    return (<div>
        <p>
            {summary.map((summary, index) => {
                return (
                    <div className="summary-item-container">
                        <div className="summary-item-content">
                            <p>{summary.tag_name}: {summary.count} 回</p>
                        </div>
                    </div>);
            })}
        </p>
    </div>);
}