import React from 'react';
import './App.css';
import Login from './pages/Login';
import Home from './pages/Home';
import Register from './pages/Register';
import Logout from './pages/Logout';
import Nav from './components/Nav';
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
        <main className="form-signin">
          <Routes>
              <Route path="/" element={<Home />} />
              <Route path="/login" element={<Login  />} />
              <Route path="/Register" element={<Register />}/>
              <Route path="/logout" element={<Logout />}/>
            </Routes>
        </main>
      </BrowserRouter>
    </div>
  );
};

export default App;
