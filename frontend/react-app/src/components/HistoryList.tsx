import { formatDate } from "../functions/utils";

export interface VideoHistory {
    video_id: string;
    watch_date: string;
    title: string;
    thumbnail_url: string;
}

export interface VideoHistoryProps {
    history: VideoHistory[];
}

export function HistoryList({ history }: VideoHistoryProps) {
    return (<div>
        <p>
            {history.map((history, index) => {
                const url = `https://www.nicovideo.jp/watch/${history.video_id}`;
                if (index === 0) {
                    console.log('thumbnail:', history.thumbnail_url)
                }
                return (
                    <div>
                        <img className="history-thumbnail" src={history.thumbnail_url} alt="thumbnail" />
                        <a href={url} target="_blank" rel="noreferrer">{history.title}</a> {formatDate(history.watch_date)}
                    </div>);
            })}
        </p>
    </div>);
}