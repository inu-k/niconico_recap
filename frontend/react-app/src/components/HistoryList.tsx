import { formatDate } from "../functions/utils";

export interface VideoHistory {
    video_id: string;
    watch_date: string;
    title: string;
}

export interface VideoHistoryProps {
    history: VideoHistory[];
}

export function HistoryList({ history }: VideoHistoryProps) {
    return (<div>
        <p>
            {history.map((history, index) => {
                const url = `https://www.nicovideo.jp/watch/${history.video_id}`;
                return (
                    <div>
                        <img className="history-thumbnail" src="https://nicovideo.cdn.nimg.jp/thumbnails/43639677/43639677.83203883.M" alt="masuo" />
                        <a href={url} target="_blank" rel="noreferrer">{history.title}</a> {formatDate(history.watch_date)}
                    </div>);
            })}
        </p>
    </div>);
}