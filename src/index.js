import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './components/searchuser/App';
//import App from './components/searchuser/TestsearchAllUserbyID';
//import App from './components/searchuser/TestsearchAllUserbyFirstName';
//import App from './components/searchuser/TestsearchAllUserbyLastName';
//import App from './components/searchuser/TestsearchAllUserbyEmail';
//import App from './components/searchuser/TestsearchAllUserbyGender';
//import App from './components/searchuser/TestsearchAllUserbyAge';
//import App from './components/searchuser/TestsearchAllUserbyAgeBetween';

//import App from './components/manageuser/TestaddUser';
//import App from './components/manageuser/TesteditUser';
//import App from './components/manageuser/TestdelteUser';


import * as serviceWorker from './serviceWorker';

ReactDOM.render(<App />, document.getElementById('root'));

//ReactDOM.render(<TestSearch />, document.getElementById('root'));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: https://bit.ly/CRA-PWA
serviceWorker.unregister();