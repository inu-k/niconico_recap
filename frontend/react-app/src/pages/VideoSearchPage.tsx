import { useEffect, useState, useRef } from "react";
import { SearchResultList, SearchResultProps } from "../components/SearchResultList";
import { fetchData } from "../functions/utils";
import { useLocation } from "react-router-dom";

// search videos page
export default function VideoSearchPage() {
    const search = useLocation().search;
    const params = new URLSearchParams(search);
    const inputRef = useRef<HTMLInputElement>(null);
    const [searchText, setSearchText] = useState(''); // title or tag
    const [searchedText, setSearchedText] = useState(''); // title or tag, 実際に検索された文字列
    const [searchResult, setSearchResult] = useState<SearchResultProps>({ results: [] });

    useEffect(() => {
        if (inputRef.current) {
            setSearchText(inputRef.current.value);
        }

        let tag = params.get('tag');
        let title = params.get('title');
        if (tag) {
            setSearchText(tag);
            handleVideoSearchByTag(tag);
        }
        if (title) {
            setSearchText(title);
            handleVideoSearchByTitle(title);
        }

        console.log('searchText:', searchText);
    }, []);

    const handleVideoSearchByTitle = (title: string) => {
        fetchData(`http://localhost:8088/videos?title=${title}`)
            .then(data => {
                setSearchedText(title);
                setSearchResult({ results: data });
            })
            .catch(error => {
                console.error(error);
                setSearchResult({ results: [] });
            });
    }

    const handleVideoSearchByTag = (tag: string) => {
        fetchData(`http://localhost:8088/videos?tag=${tag}`)
            .then(data => {
                setSearchedText(tag);
                setSearchResult({ results: data });
            })
            .catch(error => {
                console.error(error);
                setSearchResult({ results: [] });
            });
    }

    return (
        <div className="video-search-page-container">
            <h1>視聴したビデオを検索</h1>
            <input className="video-search-page-box" ref={inputRef} value={searchText} onChange={(e) => { setSearchText(e.target.value); }} type='text' placeholder='タイトルかタグを入力' />
            <button className="video-search-page-button" onClick={() => handleVideoSearchByTitle(searchText)}>タイトル検索(部分一致)</button>
            <button className="video-search-page-button" onClick={() => handleVideoSearchByTag(searchText)}>タグ検索(完全一致)</button>

            <div>
                {searchedText && (<h2>{searchedText} の検索結果: {searchResult.results.length}件</h2>)}
                <SearchResultList results={searchResult.results} />
            </div>
        </div>
    );
}