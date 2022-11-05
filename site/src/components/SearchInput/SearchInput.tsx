import React, {Component} from 'react';
import S from "./searchInput.module.scss";

class SearchInput extends Component {
    render() {
        return (
            <span className={S.input_wrapper}>
              <input type="text" className={S.search} />
            </span>
        );
    }
}

export default SearchInput;
