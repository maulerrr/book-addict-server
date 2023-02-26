import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import Navbar from './components/UI/navbar/Navbar'
import './styles/App.css';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
    <>
        <Navbar/>
        <App/>
    </>
);

