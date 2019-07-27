import React, { Component } from 'react';
import '../../bootstrap-4.1.3-dist/css/bootstrap.min.css';

class TestaddUser extends Component {
  constructor(props) {
    super(props);
    this.state = {
      user: []
    }
  }
  	
	componentDidMount() {
		fetch('http://localhost:10000/addnewuser', {
			method: 'POST',
			body: JSON.stringify({
				Id: '0',
				Firstname: 'Test_Firstname',
				Lastname: 'Test_Lastname',
				Email: 'TestEmail@mail.com',
				Gender: 'Female',
				Age: '20'
			}),
			headers: {
				 'Access-Control-Allow-Origin': '*',
				 'Content-Type': 'application/json',
			}
		}).then(response => {
				return response.json()
			}).then(json => {
				this.setState({
					user:json
				});
			});
	}
  
  render() {
    return (
			<div>
				<p><b>New User Create</b></p>
				<p>Id : {this.state.user.Id}</p>
				<p>Firstname : {this.state.user.Firstname}</p>
				<p>Lastname : {this.state.user.Lastname}</p>
				<p>Email : {this.state.user.Email}</p>
				<p>Gender : {this.state.user.Gender}</p>
				<p>Age : {this.state.user.Age}</p>
			</div>
    );
  }
  
}

export default TestaddUser;