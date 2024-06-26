import { title } from 'process';
import { useState, useEffect } from 'react';
import { useParams } from 'react-router-dom';
import styles from './VideoInfoPage.module.css';

export interface VideoInfo {
    video_id: string;
    title: string;
    tags: string[];
    thumbnail_url: string;
}

// show video information
export default function VideoInfoPage() {
    const [videoInfo, setVideoInfo] = useState<VideoInfo>({ video_id: '', title: '', tags: [], thumbnail_url: '' });
    let { videoId } = useParams();

    useEffect(() => {
        fetch(`http://localhost:8088/videos/${videoId}`)
            .then(response => response.json())
            .then(data => {
                setVideoInfo(data);
            })
            .catch(error => {
                console.error(error);
                setVideoInfo({ video_id: 'sm00000000', title: '', tags: [], thumbnail_url: '' });
            });
    }, []);

    if ('code' in videoInfo) {
        if (videoInfo.code === 404) {
            return <div>Video not found</div>;
        }
        else {
            return <div>Something went wrong</div>;
        }
    }
    if (!videoInfo.video_id) {
        return <div>Loading...</div>;
    }

    return (
        <div className={styles.container}>
            <h1>動画情報</h1>
            <h2>{videoInfo.title}</h2>
            <p>
                <img className="video-thumbnail" src={videoInfo.thumbnail_url} alt="thumbnail" />
                <a href={`https://www.nicovideo.jp/watch/${videoInfo.video_id}`} target="_blank" rel="noreferrer">{videoInfo.title}</a>
            </p>
            <p>タグ:</p>
            <ul>
                {videoInfo.tags.map(tag => <li key={tag}><a href={`http://localhost:3030/search?tag=${tag}`}>{tag}</a></li>)}
            </ul>
        </div>
    );
}