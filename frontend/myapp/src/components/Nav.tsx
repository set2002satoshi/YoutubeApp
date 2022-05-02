import React from 'react';
import {Link} from 'react-router-dom';

const Nav = () => {
    return (
        <nav className="navbar navbar-expand navbar-dark bg-dark" aria-label="Second navbar example">
        <div className="container-fluid">
          <Link to="/" className="navbar-brand">Home</Link>


          <div className="collapse navbar-collapse" >
            <ul className="navbar-nav me-auto">

              <li className="nav-item">
                <Link to="/login" className="nav-link active">login</Link>
              </li>
              <li className="nav-item">
                <Link to="/register" className="nav-link">Register</Link>
              </li>

              <li className="nav-item">
                <Link to="/logout" className="nav-link">Logout</Link>
              </li>

            </ul>
          </div>
        </div>
      </nav>
    );
};

export default Nav;