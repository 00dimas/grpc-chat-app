import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App.tsx'; // Impor yang benar
import { reportWebVitals } from './reportWebVitals.ts'; // Perbaikan impor

// Menentukan elemen root untuk aplikasi
const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement);

// Merender aplikasi ke dalam root
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

// Menyertakan fungsionalitas pelaporan web vitals
reportWebVitals(console.log); // Melakukan pemanggilan dengan memberikan parameter yang sesuai
