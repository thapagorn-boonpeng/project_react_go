import React, { Component } from 'react';
import ReactDOM from 'react-dom';

class SearchPanel extends Component {
    constructor(props) {
      super(props);
	  this.createAjaxSearch = this.createAjaxSearch.bind(this);
    }
	
	render(){
        return (
            <div>
                <input type="text" ref="valuesearch"  />
                <select ref="keysearch">
                    <option value="searchuserbyid">searchAllUserbyID</option>
                    <option value="searchuserbyfirstname">searchAllUserbyFirstName</option>
					<option value="searchuserbylastname">searchAllUserbyLastName</option>
					<option value="searchuserbyemail">searchAllUserbyEmail</option>
					<option value="searchuserbygender">searchAllUserbyGender</option>
					<option value="searchuserbyage">searchAllUserbyAge</option>
					<option value="searchuserbyagebetween">searchAllUserbyAgeBetween</option>
                </select>
                <input type="submit" onClick={this.createAjaxSearch} />
            </div>
        );
    }

    createAjaxSearch(){
        var valuesearch = ReactDOM.findDOMNode(this.refs.valuesearch).value;
        var keysearch = ReactDOM.findDOMNode(this.refs.keysearch).value;
        var URL      = 'http://localhost:10000/' + keysearch +'/' + valuesearch;
        this.props.search(URL)
    }
}

export default SearchPanel;