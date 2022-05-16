import styles from './SearchBar.module.css';
import React, { FC, Dispatch, SetStateAction, useState } from 'react';


const SearchBar: FC<{ setStateProps: Dispatch<SetStateAction<string>> }> = ({setStateProps}) => {
    const [ searchData, setSearchWord ] = useState<string>("")

    const submit = () => {
        console.log(searchData);
        setStateProps(searchData);
    }


    return (
        <div>   
                <div id={styles.berStyle}>
                    <div id={styles.siroitu} className="form-outline">
                        <input 
                            type="search" 
                            onChange={e => setSearchWord(e.target.value)}
                            id="form1" 
                            className="form-control" 
                        />
                        {/* <label className="form-label" htmlFor="form1">Search</label> */}
                    </div>
                    <button type="button" className="btn btn-primary" onClick={submit}>
                        
                        検索
                    </button>
                </div>       
        
        </div>


        
   
    )
}

export default SearchBar;