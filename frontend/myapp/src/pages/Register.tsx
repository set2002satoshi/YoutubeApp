import React, { useState} from "react";
import {useNavigate} from "react-router-dom";
import axios from "axios";

const url = "http://localhost:8000/user/Cuser";


const Register = () => {
    const [user, setUser] = useState<string>('');
    const [email, setEmail] = useState<string>('');
    const [password, setPassword] = useState<string>('');
    const [redirect, setRedirect] = useState<boolean>(false);

    
    const navigate = useNavigate();
    const submit = async () => {

        const response = await axios.post(url, {
            user,
            email,
            password,
        });
        console.log(response.data);
        setRedirect(true);

        // const response = await fetch("http://localhost:8000/user/Cuser",{
        //     method: "POST",
        //     headers: { "Content-Type": "application/json"},
        //     body: JSON.stringify({
        //         user,
        //         email,
        //         password
        //     })
        // })
    //     console.log(response);
    //     setRedirect(true);
    }
    
    if (redirect){
        navigate("/login");
    }

    return (
        <form action={url} method="post">
            <h1 className="h3 mb-3 fw-normal">Please Register</h1>
            <input className="form-control"  placeholder="Name"
                onChange={e => setUser(e.target.value)}
                value={user}            
            />
            <input type="email" className="form-control"  placeholder="name@example.com"
                onChange={e => setEmail(e.target.value)}
                value={email}
            />
            <input type="password" className="form-control" placeholder="Password"
                onChange={e => setPassword(e.target.value)}
                value={password}
            />            
            <input className="w-100 btn btn-lg btn-primary" value="Register" type="button" onClick={submit} />
        </form>
    );
};

export default Register;