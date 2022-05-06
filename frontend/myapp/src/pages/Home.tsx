import  React, { useEffect, useState } from 'react';
import { useCookies } from 'react-cookie';
import { useNavigate } from 'react-router-dom';
import axios from "axios";



const url = "http://localhost:8000";

const Home = () => {
    // const [users, setUsers] = useState<string[]>([]);
    const [ cookies ] = useCookies(['clientKey']);
    const [ data, setData ] = useState<string>('');
    const navigate = useNavigate();


    const options = {
        headers: {
            'content-Type': 'application/json;charset=utf-8',
            'Access-Control-Allow-Origin': 'http://localhost:8000',
            'clientKey': cookies.clientKey,
        },
        credentials: true,
    };
    

    useEffect(() => {
        (
            
            async () => {
                try{
                    await axios.get(url, options)
                    .then((response) => {
                        if (response.data.user === '') {
                            return false;
                        }
                        console.log(response.data.user.user)
                        setData(response.data.user.user);
                    });
                }catch (e) {
                    console.log('not user')
                    navigate("/login")
                }
                
            }
        )();
    });

    


    

    return (
        <div>
            <div>
              {data}
            </div>
            <div>
                
                
            </div>
            
        </div>
    );
};

export default Home;

