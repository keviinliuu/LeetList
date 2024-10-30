import ReactDOM from 'react-dom/client'
import axios from 'axios'
import { createBrowserRouter, RouterProvider } from 'react-router-dom';

import Landing from './pages/Landing.tsx'
import Login from './pages/Login.tsx'
import Register from './pages/Register.tsx'
import Home from './pages/Home.tsx'
import CreateList from './pages/CreateList.tsx'
import './main.css'

// env variables must start with VITE_ to be discoverable
axios.defaults.baseURL = `http://localhost:${import.meta.env.VITE_ENDPOINT}/query`;

const router = createBrowserRouter([
  {
    path: '/',
    element: <Landing />,
  },
  {
    path: 'login',
    element: <Login />
  },
  {
    path: 'register',
    element: <Register />
  },
  {
    path: 'home',
    element: <Home />
  },
  {
    path: 'create',
    element: <CreateList />
  }
])

ReactDOM.createRoot(document.getElementById('root')!).render(<RouterProvider router={router} />);
