import React from 'react';
import { useNavigate } from 'react-router-dom';

const primaryColor = '#5E64FF';

const AuthPage: React.FC = () => {
  const navigate = useNavigate();

  return (
    <div style={{ minHeight: '100vh', background: '#FFFFFF', display: 'flex', flexDirection: 'column' }}>
      {/* Header */}
      <header
        style={{
          height: 56,
          borderBottom: '1px solid #F4F4F6',
          display: 'flex',
          alignItems: 'center',
          padding: '0 16px',
        }}
      >
        <button
          type="button"
          style={{
            width: 40,
            height: 40,
            borderRadius: 10,
            background:
              'linear-gradient(180deg, rgba(94, 100, 255, 0.12) 0%, rgba(94, 100, 255, 0.2) 100%)',
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            border: 'none',
            cursor: 'pointer',
            boxShadow: '0px 3px 3px -1.5px var(--elevationshadow), 0px 1px 1px -0.5px var(--elevationshadow)',
          }}
        >
          <img src="/icons/icon.svg" alt="Swarmind" style={{ width: 24, height: 27 }} />
        </button>
      </header>

      {/* Main card */}
      <main style={{ flex: 1, display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
        <div
          style={{
            width: 480,
            padding: '48px 32px',
            borderRadius: 24,
            background: '#FFFFFF',
            border: '1px solid var(--outline-base_em, #F4F4F6)',
            boxShadow:
              '0px 1px 1px -0.5px rgba(0,0,0,0.04), 0px 3px 3px -1.5px rgba(0,0,0,0.04), 0px 24px 24px -12px rgba(0,0,0,0.04)',
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            gap: 32,
          }}
        >
          {/* Early access pill */}
          <span
            style={{
              background: '#EBECF0',
              borderRadius: 8,
              padding: '4px 12px',
              fontSize: 11,
              color: '#8C929C',
              fontWeight: 600,
            }}
          >
            Early Access
          </span>

          {/* Logo text */}
          <h1
            style={{
              fontFamily: 'Coolvetica, sans-serif',
              fontSize: 64,
              lineHeight: '48px',
              color: primaryColor,
              margin: 0,
              letterSpacing: '-0.64px',
            }}
          >
            swarmind
          </h1>

          {/* Tagline */}
          <p style={{ margin: 0, color: '#5B616D', fontSize: 15, fontWeight: 500 }}>
            Decentralized AI platform
          </p>

          {/* Buttons */}
          <div style={{ width: '100%', display: 'flex', flexDirection: 'column', gap: 16 }}>
            <button
              style={{
                width: '100%',
                height: 48,
                borderRadius: 12,
                background: primaryColor,
                color: '#FFFFFF',
                fontWeight: 600,
                fontSize: 15,
                border: 'none',
                boxShadow:
                  '0px 1px 1px -0.5px rgba(0,0,0,0.04), 0px 3px 3px -1.5px rgba(0,0,0,0.04), inset 0px 3px 4px -3px rgba(255,255,255,0.56), inset 0px 0px 8px -2px rgba(255,255,255,0.48)',
                cursor: 'pointer',
              }}
              onClick={() => navigate('/signup')}
            >
              Sign up
            </button>

            <button
              style={{
                width: '100%',
                height: 48,
                borderRadius: 12,
                background: '#F4F4F6',
                color: primaryColor,
                fontWeight: 600,
                fontSize: 15,
                border: 'none',
                boxShadow:
                  '0px 1px 1px -0.5px rgba(0,0,0,0.04), inset 0px 1px 3px -2px rgba(255,255,255,0.04), inset 0px -1px 3px -2px rgba(0,0,0,0.07)',
                cursor: 'pointer',
              }}
              onClick={() => navigate('/login')}
            >
              Log in
            </button>
          </div>
        </div>
      </main>

      {/* Footer */}
      <footer
        style={{
          height: 72,
          background: '#FFFFFF',
          display: 'flex',
          justifyContent: 'center',
          alignItems: 'center',
          gap: 8,
          position: 'relative',
        }}
      >
        <a href="#" style={{ fontSize: 12, color: '#8C929C', textDecoration: 'none' }}>
          Data Policy
        </a>
        <div style={{ width: 1, height: 16, background: '#F4F4F6' }} />
        <a href="#" style={{ fontSize: 12, color: '#8C929C', textDecoration: 'none' }}>
          User Agreement
        </a>
      </footer>
    </div>
  );
};

export default AuthPage; 