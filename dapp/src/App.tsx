// App.tsx
import React from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import DepositPage from './pages/DepositPage';
import ModelManagerPage from './pages/ModelManagerPage';
import WorkerManagerPage from './pages/WorkerManagerPage';

const App: React.FC = () => {
  return (
    <Router>
      <Routes>
        <Route path="/deposit" element={<DepositPage />} />
        <Route path="/models" element={<ModelManagerPage />} />
        <Route path="/workers" element={<WorkerManagerPage />} />
        <Route path="/" element={<div>Main page</div>} />
      </Routes>
    </Router>
  );
};

export default App;