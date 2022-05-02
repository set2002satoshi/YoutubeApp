import react, { useEffect, useState } from 'react';
import { useCookies } from 'react-cookie';
import { useNavigate } from 'react-router-dom';


const Logout = () => {
    const [ Cookies, setCookie, removeCookie ] = useCookies();
    const [ redirect, setRedirect ] = useState<boolean>(false);
    
    const navigate = useNavigate();
    useEffect(() => {
        try{
            removeCookie("clientKey");
            console.log(Cookies)
            
            setRedirect(true);
        }catch(e){
            return 
        }
    });
    if (redirect) {
        navigate("/login");
    }
    return <></>
}


export default Logout;
