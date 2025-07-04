import React from 'react';
import { useNavigate } from 'react-router-dom';
const primary = '#5E64FF';

const PasswordChangedPage: React.FC = () => {
  const navigate = useNavigate();
  return (
    <div style={{ minHeight: '100vh', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
      <div
        style={{
          width: 480,
          padding: '48px 32px',
          borderRadius: 24,
          background: '#FFFFFF',
          boxShadow: '0 1px 1px -0.5px rgba(0,0,0,0.04), 0 3px 3px -1.5px rgba(0,0,0,0.04), 0 24px 24px -12px rgba(0,0,0,0.04)',
          display: 'flex',
          flexDirection: 'column',
          alignItems: 'center',
          gap: 32,
          border: '1px solid #F4F4F6',
        }}
      >
        <div style={{ width: 48, height: 48, borderRadius: 12, background: 'rgba(34,197,94,0.15)', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
          <span style={{ fontSize: 24, color: '#22C55E' }}>✔</span>
        </div>
        <h2 style={{ margin: 0 }}>Password changed</h2>
        <p style={{ fontSize: 14, textAlign: 'center', marginTop: -12 }}>Your password has been changed successfully</p>
        <button
          type="button"
          onClick={() => navigate('/signup')}
          style={{
            width: '100%',
            height: 48,
            borderRadius: 12,
            background: primary,
            color: '#FFFFFF',
            fontWeight: 600,
            fontSize: 15,
            border: 'none',
            cursor: 'pointer',
          }}
        >
          Back to Sign Up
        </button>
      </div>
    </div>
  );
};

export default PasswordChangedPage; 