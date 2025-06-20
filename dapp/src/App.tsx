// App.tsx
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import DepositPage from './pages/DepositPage';

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path="/deposit" element={<DepositPage />} />
        <Route path="/" element={<div>Main page</div>} />
      </Routes>
    </Router>
  );
};

export default App;