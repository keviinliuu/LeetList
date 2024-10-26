import React, { useState } from 'react';
import axios from 'axios';

export default function Login() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('')
    const [error, setError] = useState(null);

    const handleLogin = async () => {
        const query = `
            mutation Login($email: String!, $password: String!) {
            login(email: $email, password: $password) {
                token
                user {
                    ID
                    email
                    password
                }
            }
        }
        `;

        const variables = { email, password };

        try {
            const response = await axios.post('', { query, variables });
            const result = response.data;

            if(result.errors) {
                setError(result.errors[0].message);
            }
            else {
                const { token, user } = result.data.login;
                console.log('Login successful', user);
            }
        } catch (err) {
            console.error('Login error:', err);
            // setError("Failed to login. Please try again.");
        }
    }

    return (
        <div className='flex flex-col min-h-screen w-screen bg-richBlack items-center justify-center'>
            <div className='w-[25rem] h-[30rem] bg-white rounded-lg items-center justify-center font-main flex flex-col gap-y-2'>
                <div className='text-3xl'>
                    Login to LeetList
                </div>

                <div> 
                    Email
                </div>
                <input 
                    type='text'
                    className='border-2 border-black rounded-md p-1' 
                    placeholder='Email'
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                />

                <div> 
                    Password
                </div>
                <input 
                    type='password'
                    className='border-2 border-black rounded-md p-1' 
                    placeholder='Password'
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />

                <button 
                    className='transition ease-in-out delay-50 hover:scale-110 text-xl h-10 w-20 bg-prussian text-white rounded-lg'
                    onClick={handleLogin}
                >
                    Login
                </button>
            </div>
        </div>
    )
}