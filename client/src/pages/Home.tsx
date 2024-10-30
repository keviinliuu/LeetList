import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';

export default function Home() {
    const [user, setUser] = useState<any | null>(null);
    const navigate = useNavigate();

    useEffect(() => {
        const user = localStorage.getItem('user');
        
        if(!user) {
            
        }

        // if(!user) {
        //     navigate('/')
        // }

        // const storedUser = localStorage.getItem('user');
        // if (storedUser) {
        //     setUser(JSON.parse(storedUser));
        // }
    }, []);

    return (
        <div className='flex flex-col min-h-screen w-screen bg-richBlack text-cerulean text-3xl font-main p-10'>
            {/* <div>Hello, {user.email}!</div> */}
            <div>Your Lists:</div>
            {/* If you have lists, map through them */}
            {/* {user.lists && user.lists.length > 0 ? (
                <ul>
                    {user.lists.map((list: any) => (
                        <li key={list.ID}>
                            <div className="text-white text-xl">{list.title}</div>
                        </li>
                    ))}
                </ul>
            ) : (
                <div>No lists available</div>
            )} */}
        </div>
    );
}
