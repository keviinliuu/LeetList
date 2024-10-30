import { Link, useNavigate } from 'react-router-dom';
import { useState } from 'react';
import axios from 'axios';
import Logo from '../assets/logo.png'

export default function Login() {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('')
    const [error, setError] = useState<string | null>(null);
    const navigate = useNavigate();

    const handleLogin = async () => {
        const query = `
            mutation Login($email: String!, $password: String!) {
                login(email: $email, password: $password) {
                    token
                    user {
                        ID
                        email
                    }
                }
            }
        `;

        const variables = { email, password };

        try {
            const response = await axios.post('', {query, variables});
            console.log(response);
            const result = response.data;

            if(result.errors) {
                setError(result.errors[0].message);
            }
            else {
                const { token, user } = result.data.login;
                localStorage.setItem('token', token);
                localStorage.setItem('user', JSON.stringify(user));
                navigate('/home')
                console.log('Login successful', user);
            }
        } catch (err) {
            console.error('Login error:', err);
            setError("Failed to login. Please try again.");
        }
    }

    return (
        <div className='flex flex-col min-h-screen w-screen bg-richBlack items-center justify-center'>
            <div className='w-[32rem] h-[40rem] bg-white rounded-lg items-center font-main flex flex-col gap-y-2 pt-10'>
                <img className='w-12 pb-5' src={Logo}></img>
                
                <div className='text-4xl'>
                    Login to LeetList
                </div>

                <div className='pt-10 text-xl'>
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
                </div>

                <div className='pt-10 pb-10 text-xl'>
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
                </div>

                <button 
                    className='transition ease-in-out delay-50 hover:bg-cerulean text-2xl h-14 w-32 bg-prussian text-white rounded-lg'
                    onClick={handleLogin}
                >
                    Login
                </button>

                <div className='p-3'>New user? <Link to='/register' className='text-cerulean hover:underline transition ease-in-out'>Register here!</Link></div>

                {error && <div className='text-red-500'>{error}</div>}
            </div>
        </div>
    )
}