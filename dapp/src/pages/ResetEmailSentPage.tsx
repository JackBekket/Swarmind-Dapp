import React from 'react';
import { useNavigate } from 'react-router-dom';
const primary = '#5E64FF';
const greyText = '#8C929C';

const ResetEmailSentPage: React.FC = () => {
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
        {/* mail icon placeholder */}
        <div style={{ width: 48, height: 48, borderRadius: 12, background: 'rgba(94,100,255,0.1)', display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
          <span style={{ color: primary, fontSize: 24 }}>✉️</span>
        </div>
        <h2 style={{ margin: 0 }}>Check your email</h2>
        <p style={{ textAlign: 'center', color: greyText, fontSize: 14, marginTop: -12 }}>
          Please check the email address for instructions to reset your password
        </p>
        <button
          type="button"
          onClick={() => navigate('/forgot-password')}
          style={{ border: 'none', background: 'none', color: primary, fontSize: 14, cursor: 'pointer' }}
        >
          ← Resend email
        </button>
      </div>
    </div>
  );
};

export default ResetEmailSentPage; 