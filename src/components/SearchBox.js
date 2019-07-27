import React from 'react';


class SearchBox extends Component {
	
render: function(){
        return (
            <div>
                <input type="text" ref="query" />
                <select ref="category">
                    <option value="software">Apps</option>
                    <option value="movie">Films</option>
                </select>
                <input type="submit" onClick={this.createAjax} />
            </div>
        );
    },

    createAjax: function(){
        var query    = ReactDOM.findDOMNode(this.refs.query).value;
        var category = ReactDOM.findDOMNode(this.refs.category).value;
        var URL      = 'https://itunes.apple.com/search?term=' + query +'&country=us&entity=' + category;
        this.props.search(URL)
    }

}

export default SearchBox;