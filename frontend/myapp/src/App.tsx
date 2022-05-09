// import React from 'react';
// import ReactDOM from 'react-dom';

import styles from './App.module.css';
import Nav from './components/Nav';
import Login from './pages/Login';
import Register from './pages/Register';
import Logout from './pages/Logout';
import Home from './pages/Home';
import AddChannel from './pages/AddChannel';
import {BrowserRouter, Route, Routes} from 'react-router-dom';
// import { useCookies } from 'react-cookie';


function App() {

  // const [ cookies, setCookie, removeCookie ] = useCookies();

  // const RemoveCookieHandler = (keyword: string): boolean => {
  //   removeCookie(keyword);
  //   return true;
  // }

  // const lookCookieHandler = (keyword: string) => {   
  //     return cookies;
  // }


  // const setCookieHandler = (keyword: string, value: string): boolean => {
  //     setCookie(keyword, value);
  //     return true;
  // }



  return (
    <div className="App">
      <BrowserRouter>
          <Nav />
        <main id={styles.mainArea}>
          <Routes>
              <Route path="/login" element={<Login />} />
              <Route path="/Register" element={<Register />}/>
              <Route path="/logout" element={<Logout/>}/>
              <Route path="/" element={<Home />} />
              <Route path="/search" element={<AddChannel />} />
            </Routes>
        </main>
      </BrowserRouter>
    </div>
  );
};

export default App;
