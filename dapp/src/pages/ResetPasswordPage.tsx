import React, { useState } from 'react';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { confirmPasswordReset } from '../services/auth';
import { zxcvbn } from '@zxcvbn-ts/core';
import TopBar from '../components/TopBar';
import PageWrapper from '../components/PageWrapper';
import Footer from '../components/Footer';
import ErrorLine from '../components/ErrorLine';

const primary = '#5E64FF';
const greyBg = '#F4F4F6';
const greyText = '#8C929C';
const errorCol = '#F04438';

const red = '#F04438';
const yellow = '#F8B400';
const green = '#22C55E';

const barEmpty = 'rgba(0,0,0,0.12)';

const strengthLabels = [
  { text: 'Very Weak', emoji: '😟' },
  { text: 'Weak', emoji: '😕' },
  { text: 'Strong', emoji: '🙂' },
  { text: 'Very Strong', emoji: '💪' },
];

const getStrength = (pwd: string) => {
  if (!pwd) return { level: 0, entropy: 0 };
  const res = zxcvbn(pwd);
  const entropy = Math.log2(res.guesses);
  let level = 1;
  if (res.score === 2) level = 2;
  if (res.score === 3) level = 3;
  if (res.score === 4) level = 4;
  return { level, entropy };
};

const errorIcon = '/icons/close.svg';

const ResetPasswordPage: React.FC = () => {
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const oobCode = searchParams.get('oobCode');

  const [pwd, setPwd] = useState('');
  const [pwd2, setPwd2] = useState('');
  const [focus1, setFocus1] = useState(false);
  const [focus2, setFocus2] = useState(false);
  const [show1, setShow1] = useState(false);
  const [show2, setShow2] = useState(false);
  const [pwdErr, setPwdErr] = useState('');
  const [matchErr, setMatchErr] = useState('');
  const [loading, setLoading] = useState(false);

  if (!oobCode) {
    return <p>Invalid or missing reset token</p>;
  }

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    let valid = true;
    if (pwd !== pwd2) {
      setMatchErr('Passwords do not match');
      valid = false;
    } else setMatchErr('');
    if (getStrength(pwd).level < 3) {
      setPwdErr('Choose a stronger password');
      valid = false;
    } else setPwdErr('');
    if (!valid) return;
    setLoading(true);
    try {
      await confirmPasswordReset(oobCode, pwd);
      navigate('/password-changed');
    } catch (e: any) {
      console.error(e);
      setPwdErr(e?.message || 'Error resetting password');
    } finally {
      setLoading(false);
    }
  };

  const { level: strength, entropy } = getStrength(pwd);

  return (
    <div style={{ minHeight: '100vh', display: 'flex', flexDirection:'column' }}>
      <TopBar />
      <div style={{flex:1, display:'flex', alignItems:'center', justifyContent:'center'}}>
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
          <h2 style={{ margin: 0 }}>Reset password</h2>
          <p style={{ textAlign: 'center', color: greyText, fontSize: 14, marginTop: -12 }}>
            Enter a new password below to change your password
          </p>

          {/* new password field */}
          <div style={{ width: '100%', display: 'flex', flexDirection: 'column', gap: 4 }}>
            <label style={{ fontSize: 13, fontWeight: 500, color: greyText }}>New password</label>
            <div style={{ position: 'relative' }}>
              <input
                type={show1 ? 'text' : 'password'}
                placeholder="Enter new password"
                value={pwd}
                onChange={(e) => setPwd(e.target.value)}
                onFocus={() => setFocus1(true)}
                onBlur={() => setFocus1(false)}
                style={{
                  width: '100%',
                  height: 48,
                  borderRadius: 12,
                  background: focus1 ? '#FFFFFF' : greyBg,
                  border: pwdErr ? `1px solid ${errorCol}` : focus1 ? `1px solid ${primary}` : '1px solid transparent',
                  boxShadow: focus1 && pwdErr ? '0 0 0 3px rgba(240,68,56,0.2)' : focus1 ? '0 0 0 3px rgba(94,100,255,0.2)' : 'none',
                  padding: '0 44px 0 12px',
                  boxSizing: 'border-box',
                  fontSize: 15,
                  outline: 'none',
                  transition: 'border 0.15s, box-shadow 0.15s, background 0.15s',
                }}
              />
              <button
                type="button"
                onClick={() => setShow1((p) => !p)}
                style={{ position: 'absolute', right: 12, top: '50%', transform: 'translateY(-50%)', background: 'none', border: 'none', cursor: 'pointer' }}
              >
                {show1 ? (
                  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke={pwdErr ? errorCol : greyText} strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M17.94 17.94A10.94 10.94 0 0 1 12 20c-5 0-9.27-3-11-8 1.13-2.9 3-5.23 5.32-6.64" /><path d="M3 3l18 18" /></svg>
                ) : (
                  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke={pwdErr ? errorCol : greyText} strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><circle cx="12" cy="12" r="3" /><path d="M2.05 12a10.97 10.97 0 0 1 19.9 0 10.97 10.97 0 0 1-19.9 0z" /></svg>
                )}
              </button>
            </div>
            {/* strength bar */}
            <div style={{ width: '100%', display: 'flex', flexDirection: 'column', gap: 4 }}>
              <div style={{ display: 'flex', gap: 8 }}>
                {[0, 1, 2, 3].map((i) => {
                  let color = barEmpty;
                  if (strength === 1 && i === 0) color = red;
                  if (strength === 2 && i < 2) color = yellow;
                  if (strength === 3 && i < 3) color = green;
                  if (strength === 4) color = green;
                  return (
                    <div
                      key={i}
                      style={{ flex: 1, height: 6, borderRadius: 1000, background: color }}
                    />
                  );
                })}
              </div>
              <span style={{ fontSize: 13, color: greyText, alignSelf: 'flex-end' }}>
                {pwd
                  ? `${strengthLabels[strength - 1]?.text} ${strengthLabels[strength - 1]?.emoji}`
                  : 'Choose a strong password'}
              </span>
            </div>
          </div>

          {/* repeat */}
          <div style={{ width: '100%', display: 'flex', flexDirection: 'column', gap: 4 }}>
            <label style={{ fontSize: 13, fontWeight: 500, color: greyText }}>Repeat new password</label>
            <div style={{ position: 'relative' }}>
              <input
                type={show2 ? 'text' : 'password'}
                placeholder="Re-enter new password"
                value={pwd2}
                onChange={(e) => setPwd2(e.target.value)}
                onFocus={() => setFocus2(true)}
                onBlur={() => setFocus2(false)}
                style={{
                  width: '100%',
                  height: 48,
                  borderRadius: 12,
                  background: focus2 ? '#FFFFFF' : greyBg,
                  border: matchErr ? `1px solid ${errorCol}` : focus2 ? `1px solid ${primary}` : '1px solid transparent',
                  boxShadow: focus2 && matchErr ? '0 0 0 3px rgba(240,68,56,0.2)' : focus2 ? '0 0 0 3px rgba(94,100,255,0.2)' : 'none',
                  padding: '0 44px 0 12px',
                  boxSizing: 'border-box',
                  fontSize: 15,
                  outline: 'none',
                  transition: 'border 0.15s, box-shadow 0.15s, background 0.15s',
                }}
              />
              <button
                type="button"
                onClick={() => setShow2((p) => !p)}
                style={{ position: 'absolute', right: 12, top: '50%', transform: 'translateY(-50%)', background: 'none', border: 'none', cursor: 'pointer' }}
              >
                {show2 ? (
                  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke={matchErr ? errorCol : greyText} strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M17.94 17.94A10.94 10.94 0 0 1 12 20c-5 0-9.27-3-11-8 1.13-2.9 3-5.23 5.32-6.64" /><path d="M3 3l18 18" /></svg>
                ) : (
                  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke={matchErr ? errorCol : greyText} strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><circle cx="12" cy="12" r="3" /><path d="M2.05 12a10.97 10.97 0 0 1 19.9 0 10.97 10.97 0 0 1-19.9 0z" /></svg>
                )}
              </button>
            </div>
          </div>

          <div style={{ width: 416, display: 'flex', flexDirection: 'column', alignItems: 'flex-start', margin: '2px 0 0 0' }}>
            {pwdErr && <ErrorLine>{pwdErr}</ErrorLine>}
          </div>

          <div style={{ width: 416, display: 'flex', flexDirection: 'column', alignItems: 'flex-start', margin: '2px 0 0 0' }}>
            {matchErr && <ErrorLine>{matchErr}</ErrorLine>}
          </div>

          <button
            type="submit"
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
            {loading ? 'Resetting…' : 'Reset password'}
          </button>
        </form>
        </PageWrapper>
      </div>
      <Footer />
    </div>
  );
};

export default ResetPasswordPage; 