import React, {  useState } from 'react';
import styles from "./AddChannel.module.css";
import SearchBar from '../components/SearchBar';
import SearchResult from '../components/SearchResult';


const AddChannel = () => {
    const [ SearchItem, setKeyword ] = useState<string>('')

    



    return (
        <div id={styles.mainArea}>
                <div id="styles.searchBox">
                    <SearchBar setStateProps={setKeyword}/>
                </div>
                <div id={styles.searchResult}>
                    <SearchResult searchResultProps={SearchItem} />
                </div>
        </div>
    )
}

export default AddChannel;