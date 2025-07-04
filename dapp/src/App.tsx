// App.tsx
import React from 'react';
import { BrowserRouter as Router, Routes, Route, useLocation } from 'react-router-dom';
import { AnimatePresence } from 'framer-motion';
import DepositPage from './pages/DepositPage';
import ModelManagerPage from './pages/ModelManagerPage';
import WorkerManagerPage from './pages/WorkerManagerPage';
import AuthPage from './pages/AuthPage';
import LoginPage from './pages/LoginPage';
import ForgotPasswordPage from './pages/ForgotPasswordPage';
import ResetEmailSentPage from './pages/ResetEmailSentPage';
import ResetPasswordPage from './pages/ResetPasswordPage';
import PasswordChangedPage from './pages/PasswordChangedPage';
import RegistrationPage from './pages/RegistrationPage';

const AnimatedRoutes: React.FC = () => {
  const location = useLocation();

  return (
    <AnimatePresence mode="wait">
      <Routes location={location} key={location.pathname}>
        <Route path="/deposit" element={<DepositPage />} />
        <Route path="/models" element={<ModelManagerPage />} />
        <Route path="/workers" element={<WorkerManagerPage />} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/auth" element={<AuthPage />} />
        <Route path="/forgot-password" element={<ForgotPasswordPage />} />
        <Route path="/forgot-password/sent" element={<ResetEmailSentPage />} />
        <Route path="/reset-password" element={<ResetPasswordPage />} />
        <Route path="/password-changed" element={<PasswordChangedPage />} />
        <Route path="/signup" element={<RegistrationPage />} />
        <Route path="/" element={<div>Main page</div>} />
      </Routes>
    </AnimatePresence>
  );
};

const App: React.FC = () => {
  return (
    <Router>
      <AnimatedRoutes />
    </Router>
  );
};

export default App;