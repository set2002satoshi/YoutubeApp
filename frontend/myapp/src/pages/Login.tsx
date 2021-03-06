import styles from './Form.module.css';
import { useState } from 'react';
import {useNavigate} from "react-router-dom";
import { useCookies } from 'react-cookie';
import axios from "axios";

const url = "http://localhost:8000/user/login";


const Login = () => {
    const [email, setEmail] = useState<string>('');
    const [password, setPassword] = useState<string>('');
    const [redirect, setRedirect] = useState<boolean>(false);
    const [cookies, setCookie] = useCookies();
    const navigate = useNavigate();

    const options = {
        headers: { 
            'Content-Type': 'application/json;charset=utf-8',  
            'Access-Control-Allow-Origin': 'http://localhost:8000', 
        },
        credentials: true,
    };

    const params = {
        email,
        password,
    };



    const submit = async () => {
        
        
        await axios.post(url, params, options)
        .then((response) => {
            setCookie("clientKey", response.data);
            console.log(cookies);
        });

        setRedirect(true);
    }
    
    if (redirect){
        navigate("/");
    }



    return (
        <form id={styles.mainArea} action={url} method="post">
            <input type="email" className="form-control"  placeholder="name@example.com"
                onChange={e => setEmail(e.target.value)}
                value={email}
            />
            <input type="password" className="form-control" placeholder="Password"
                onChange={e => setPassword(e.target.value)}
                value={password}
            />            
            <input className="w-100 btn btn-lg btn-primary" value="Login" type="button" onClick={submit} />
        </form>
    );
};

export default Login;

