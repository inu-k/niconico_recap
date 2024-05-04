export interface SearchResult {
    video_id: string;
    title: string;
    thumbnail_url: string;
}

export interface SearchResultProps {
    results: SearchResult[];
}

export function SearchResultList({ results }: SearchResultProps) {
    return (<div>
        <p>
            {results.map((result, index) => {
                const url = `https://www.nicovideo.jp/watch/${result.video_id}`;
                return (
                    <div className="search-result-item-container">
                        <div className="search-result-item-thumbnail">
                            <img src={result.thumbnail_url} alt="thumbnail" />
                        </div>
                        <div className="search-result-item-content">
                            <div><a href={url} target="_blank" rel="noreferrer">{result.title}</a></div>
                            <div className="search-result-item-video-info"><a href={`http://localhost:3030/videos/${result.video_id}`}>動画情報</a></div>
                        </div>
                    </div>);
            })}
        </p>
    </div>);
}