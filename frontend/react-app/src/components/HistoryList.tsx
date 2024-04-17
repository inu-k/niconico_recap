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
                return (
                    <div className="history-item-container">
                        <div className="history-item-thumbnail">
                            <img src={history.thumbnail_url} alt="thumbnail" />
                        </div>
                        <div className="history-item-content">
                            <p><a href={url} target="_blank" rel="noreferrer">{history.title}</a></p>
                            <p className="history-item-video-info"><a href={`http://localhost:3030/videos/${history.video_id}`}>動画情報</a></p>
                            <p>{formatDate(history.watch_date)} 視聴</p>
                        </div>
                    </div>);
            })}
        </p>
    </div>);
}