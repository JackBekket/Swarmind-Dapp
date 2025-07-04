import React, { useState } from 'react';
import { sendPasswordResetEmail } from '../services/auth';
import { useNavigate } from 'react-router-dom';
import TopBar from '../components/TopBar';
import PageWrapper from '../components/PageWrapper';
import Footer from '../components/Footer';
import ErrorLine from '../components/ErrorLine';

const primary = '#5E64FF';
const greyBg = '#F4F4F6';
const greyText = '#8C929C';
const errorCol = '#F04438';
const errorIcon = '/icons/close.svg';

const ForgotPasswordPage: React.FC = () => {
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [focus, setFocus] = useState(false);
  const [err, setErr] = useState('');
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)) {
      setErr('Invalid email');
      return;
    }
    setErr('');
    setLoading(true);
    try {
      await sendPasswordResetEmail(email);
      navigate('/forgot-password/sent');
    } catch (e: any) {
      console.error(e);
      setErr(e?.message || 'Error sending email');
    } finally {
      setLoading(false);
    }
  };

  return (
    <div style={{ minHeight: '100vh', display: 'flex', flexDirection: 'column' }}>
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
            gap: 32,
          }}
        >
          <h2 style={{ fontSize: 24, margin: 0, fontWeight: 700 }}>Reset password</h2>
          <p style={{ textAlign: 'center', fontSize: 14, color: greyText, marginTop: -12 }}>
            Enter your email address and we will send you instructions to reset your password
          </p>

          <div style={{ width: '100%', display: 'flex', flexDirection: 'column', gap: 4 }}>
            <label style={{ fontSize: 13, fontWeight: 500, color: greyText }}>Email</label>
            <input
              type="email"
              placeholder="Enter your email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              onFocus={() => setFocus(true)}
              onBlur={() => setFocus(false)}
              style={{
                width: '100%',
                height: 48,
                borderRadius: 12,
                background: focus ? '#FFFFFF' : greyBg,
                border: err ? `1px solid ${errorCol}` : focus ? `1px solid ${primary}` : '1px solid transparent',
                boxShadow: focus ? '0 0 0 3px rgba(94,100,255,0.2)' : 'none',
                padding: '0 12px',
                fontSize: 15,
                boxSizing: 'border-box',
                transition: 'border 0.15s, box-shadow 0.15s, background 0.15s',
                outline: 'none',
              }}
            />
            <div style={{ width: 416, display: 'flex', flexDirection: 'column', alignItems: 'flex-start', margin: '2px 0 0 0' }}>
              {err && <ErrorLine>{err}</ErrorLine>}
            </div>
          </div>

          <button
            type='submit'
            disabled={loading}
            style={{
              width: '100%',
              height: 48,
              borderRadius: 12,
              background: primary,
              color: '#FFFFFF',
              fontWeight: 600,
              fontSize: 15,
              border: 'none',
              cursor: loading ? 'default' : 'pointer',
              opacity: loading ? 0.6 : 1,
              boxShadow: '0 1px 1px -0.5px var(--elevationshadow), 0px 3px 3px -1.5px var(--elevationshadow), inset 0 3px 4px -3px rgba(255,255,255,0.56), inset 0 0 8px -2px rgba(255,255,255,0.48)',
            }}
          >
            {loading ? 'Sending…' : 'Continue'}
          </button>

          <button
            type='button'
            onClick={() => navigate('/signup')}
            style={{
              border: 'none',
              background: 'none',
              color: primary,
              fontSize: 14,
              cursor: 'pointer',
            }}
          >
            ← Back to Sign Up
          </button>
        </form>
        </PageWrapper>
      </main>
      <Footer />
    </div>
  );
};

export default ForgotPasswordPage;
