import React from 'react';

const Table = ({ users }) => {
  return (
    <table className="table">
      <thead>
        <tr>
          <th>Id</th>
          <th>FirstName</th>
          <th>LastName</th>
          <th>Email</th>
          <th>Gender</th>
          <th>Age</th>
        </tr>
      </thead>
      <tbody>
         { (users.length > 0) ? users.map( (user, index) => {
           return (
            <tr key={ index }>
              <td>{ user.Id }</td>
              <td>{ user.Firstname }</td>
              <td>{ user.Lastname }</td>
              <td>{ user.Email }</td>
              <td>{ user.Gender }</td>
              <td>{ user.Age }</td>
            </tr>
          )
         }) : <tr><td colSpan="5">Loading...</td></tr> }
      </tbody>
    </table>
  );
}

export default Table