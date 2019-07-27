import React, { Component } from 'react';
import Table from './Table.js';
import SearchBox from  './SearchBox.js'
import '../bootstrap-4.1.3-dist/css/bootstrap.min.css';

class TestSearch extends Component {
  
   getInitialState: function() {
        return {
            searchResults: []
        }
    },
    

    showResults: function(response){
        this.setState({
            searchResults: response.results
        })
    },
    
    search: function(URL){
        $.ajax({
            type: "GET",
            dataType: 'json',
            url: URL,
            success: function(response){
                this.showResults(response);
            }.bind(this)
        });
    },

    render: function(){
        return (
            <div>
				Test xxxxx
                <SearchBox search={this.search} />
                <Results searchResults={this.state.searchResults} />
            </div>
        );
    },
}

export default TestSearch;