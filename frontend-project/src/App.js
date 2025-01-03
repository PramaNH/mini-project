import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Login from './pages/Login';
import Authenticated from './pages/Authenticated';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/authenticated" element={<Authenticated />} />
        <Route path="/" element={<Login />} />
      </Routes>
    </Router>
  );
}

export default App;
