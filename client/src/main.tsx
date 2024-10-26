import ReactDOM from 'react-dom/client'
import axios from 'axios'
import { createBrowserRouter, RouterProvider } from 'react-router-dom';

import Home from './pages/Home.tsx'
import Login from './pages/Login.tsx'
import './main.css'

axios.defaults.baseURL = `http://localhost:${import.meta.env.GRAPHQL_PORT}/query/`;

const router = createBrowserRouter([
  {
    path: '/',
    element: <Home />,
  },
  {
    path: 'login',
    element: <Login />
  }
])

ReactDOM.createRoot(document.getElementById('root')!).render(<RouterProvider router={router} />);
