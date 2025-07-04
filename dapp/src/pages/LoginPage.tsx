import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import TopBar from '../components/TopBar';
import Footer from '../components/Footer';
import PageWrapper from '../components/PageWrapper';
import ErrorLine from '../components/ErrorLine';

const primary = '#5E64FF';
const greyBg = '#F4F4F6';
const greyText = '#8C929C';
const error = '#F04438';

const LoginPage: React.FC = () => {
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [showPwd, setShowPwd] = useState(false);
  const [emailErr, setEmailErr] = useState('');
  const [pwdErr, setPwdErr] = useState('');
  const [emailFocus, setEmailFocus] = useState(false);
  const [pwdFocus, setPwdFocus] = useState(false);

  const errorIcon = '/icons/close.svg';

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    let valid = true;
    if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) {
      setEmailErr('Invalid email');
      valid = false;
    } else {
      setEmailErr('');
    }
    if (password.length < 6) {
      setPwdErr('Password too short');
      valid = false;
    } else {
      setPwdErr('');
    }
    if (!valid) return;

    // TODO real auth
    alert('Logged in!');
  };

  return (
    <div style={{ minHeight: '100vh', background: '#FFFFFF', display: 'flex', flexDirection: 'column' }}>
      <TopBar />

      <main style={{ flex: 1, display: 'flex', alignItems: 'center', justifyContent: 'center' }}>
        <PageWrapper>
        <form
          onSubmit={handleSubmit}
          style={{
            width: 480,
            padding: '48px 32px',
            borderRadius: 24,
            background: '#FFFFFF',
            border: '1px solid var(--outline-base_em, #F4F4F6)',
            boxShadow: '0px 1px 1px -0.5px var(--elevationshadow), 0px 3px 3px -1.5px var(--elevationshadow), 0px 24px 24px -12px var(--elevationshadow)',
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            gap: 24,
          }}
        >
          <h1 style={{ fontFamily: 'Coolvetica, sans-serif', fontSize: 64, lineHeight: '48px', color: primary, margin: 0, letterSpacing: '-0.64px' }}>swarmind</h1>

          {/* Email field */}
          <div style={{ width: '100%', display: 'flex', flexDirection: 'column', gap: 4 }}>
            <label style={{ fontSize: 13, fontWeight: 500, color: greyText }}>Email</label>
            <input
              type="email"
              placeholder="Enter your email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              style={{
                width: '100%',
                height: 48,
                borderRadius: 12,
                background: emailFocus ? '#FFFFFF' : greyBg,
                border: emailErr
                  ? `1px solid ${error}`
                  : emailFocus
                    ? `1px solid ${primary}`
                    : '1px solid transparent',
                boxShadow: emailErr && emailFocus
                  ? '0 0 0 3px rgba(240,68,56,0.2)'
                  : emailFocus
                    ? '0 0 0 3px rgba(94,100,255,0.2)'
                    : 'none',
                padding: '0 12px',
                fontSize: 15,
                outline: 'none',
                boxSizing: 'border-box',
                transition: 'border 0.15s, box-shadow 0.15s, background 0.15s',
              }}
              onFocus={() => setEmailFocus(true)}
              onBlur={() => setEmailFocus(false)}
            />
            <div style={{ width: 416, display: 'flex', flexDirection: 'column', alignItems: 'flex-start', margin: '2px 0 0 0' }}>
              {emailErr && <ErrorLine>{emailErr}</ErrorLine>}
            </div>
          </div>

          {/* Password field */}
          <div style={{ width: '100%', display: 'flex', flexDirection: 'column', gap: 4 }}>
            <label style={{ fontSize: 13, fontWeight: 500, color: greyText }}>Password</label>
            <div style={{ position: 'relative' }}>
              <input
                type={showPwd ? 'text' : 'password'}
                placeholder="Input text"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                style={{
                  width: '100%',
                  height: 48,
                  borderRadius: 12,
                  background: pwdFocus ? '#FFFFFF' : greyBg,
                  border: pwdErr
                    ? `1px solid ${error}`
                    : pwdFocus
                      ? `1px solid ${primary}`
                      : '1px solid transparent',
                  boxShadow: pwdErr && pwdFocus
                    ? '0 0 0 3px rgba(240,68,56,0.2)'
                    : pwdFocus
                      ? '0 0 0 3px rgba(94,100,255,0.2)'
                      : 'none',
                  padding: '0 44px 0 12px',
                  boxSizing: 'border-box',
                  fontSize: 15,
                  outline: 'none',
                  transition: 'border 0.15s, box-shadow 0.15s, background 0.15s',
                }}
                onFocus={() => setPwdFocus(true)}
                onBlur={() => setPwdFocus(false)}
              />
              <button
                type="button"
                onClick={() => setShowPwd((p) => !p)}
                style={{
                  position: 'absolute',
                  top: '50%',
                  right: 12,
                  transform: 'translateY(-50%)',
                  background: 'none',
                  border: 'none',
                  cursor: 'pointer',
                  padding: 0,
                }}
              >
                {showPwd ? (
                  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke={pwdErr ? error : greyText} strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                    <path d="M17.94 17.94A10.94 10.94 0 0 1 12 20c-5 0-9.27-3-11-8 1.13-2.9 3-5.23 5.32-6.64" />
                    <path d="M3 3l18 18" />
                  </svg>
                ) : (
                  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke={pwdErr ? error : greyText} strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
                    <circle cx="12" cy="12" r="3" />
                    <path d="M2.05 12a10.97 10.97 0 0 1 19.9 0 10.97 10.97 0 0 1-19.9 0z" />
                  </svg>
                )}
              </button>
            </div>
            <div style={{ width: 416, display: 'flex', flexDirection: 'column', alignItems: 'flex-start', margin: '2px 0 0 0' }}>
              {pwdErr && <ErrorLine>{pwdErr}</ErrorLine>}
            </div>
            <div style={{ width: '100%', textAlign: 'right' }}>
              <button type="button" onClick={() => navigate('/forgot-password') } style={{ border: 'none', background: 'none', color: primary, fontSize: 12, cursor: 'pointer' }}>
                Forgot password?
              </button>
            </div>
          </div>

          <button
            type="submit"
            style={{
              width: '100%',
              height: 48,
              borderRadius: 12,
              background: primary,
              color: '#FFFFFF',
              fontWeight: 600,
              fontSize: 15,
              border: 'none',
              boxShadow: '0px 1px 1px -0.5px var(--elevationshadow), 0px 3px 3px -1.5px var(--elevationshadow), inset 0 3px 4px -3px rgba(255,255,255,0.56), inset 0 0 8px -2px rgba(255,255,255,0.48)',
              cursor: 'pointer',
            }}
          >
            Continue
          </button>

          <p style={{ fontSize: 13, color: greyText }}>
            Don't have an account?{' '}
            <button type="button" onClick={() => navigate('/signup')} style={{ border: 'none', background: 'none', color: primary, cursor: 'pointer', fontSize: 13 }}>
              Sign up
            </button>
          </p>
        </form>
        </PageWrapper>
      </main>

      <Footer />
    </div>
  );
};

export default LoginPage; 