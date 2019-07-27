import React, { Component } from 'react';
import SearchPanel from './SearchPanel.js';
import Table from './Table.js';

import '../../bootstrap-4.1.3-dist/css/bootstrap.min.css';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      users: []
    }
  }

  search(URL) {
	 console.log(URL)
  }
  
  componentDidMount() {
        fetch('http://localhost:10000/allusers')
        .then( response => response.json())
        .then(
            (result) => {
                this.setState({
                    isLoaded : true,
                    users : result
                });
            },
            (error) => {
                this.setState({
                    isLoaded: true,
                    error
                })
            },
        )
    }
  
  render() {
    return (
      <div className="App">
        <SearchPanel search={this.search} />
		<Table users={ this.state.users } />
      </div>
    );
  }
  
}

export default App;