import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import TopBar from '../components/TopBar';
import PageWrapper from '../components/PageWrapper';
import Footer from '../components/Footer';
import ErrorLine from '../components/ErrorLine';
import { zxcvbn } from '@zxcvbn-ts/core';

const primary = '#5E64FF';
const greyBg = '#F4F4F6';
const greyText = '#8C929C';
const errorCol = '#F04438';

const RegistrationPage: React.FC = () => {
  const navigate = useNavigate();
  const [email, setEmail] = useState('');
  const [pwd, setPwd] = useState('');
  const [pwd2, setPwd2] = useState('');
  const [emailErr, setEmailErr] = useState('');
  const [pwdErr, setPwdErr] = useState('');
  const [matchErr, setMatchErr] = useState('');
  const [agreeErr,setAgreeErr]=useState('');
  const [agree, setAgree] = useState(false);
  const [showPwd1, setShowPwd1] = useState(false);
  const [showPwd2, setShowPwd2] = useState(false);
  const [loading, setLoading] = useState(false);
  const errorIcon = '/icons/close.svg';
  const [emailFocus,setEmailFocus]=useState(false);
  const [pwd1Focus,setPwd1Focus]=useState(false);
  const [pwd2Focus,setPwd2Focus]=useState(false);

  // strength utils copied from ResetPasswordPage
  const red='#F04438';
  const yellow='#F8B400';
  const green='#22C55E';
  const barEmpty='rgba(0,0,0,0.12)';

  const strengthLabels=[{text:'Very Weak',emoji:'😟'},{text:'Weak',emoji:'😕'},{text:'Strong',emoji:'🙂'},{text:'Very Strong',emoji:'💪'}];

  const getStrength=(p:string)=>{
    if(!p) return {level:0,entropy:0};
    const res = zxcvbn(p);
    const entropy = Math.log2(res.guesses);
    let level = res.score;
    return {level,entropy};
  };

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    let valid=true;
    if(!/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(email)){
      setEmailErr('Invalid email');valid=false;
    }else setEmailErr('');

    if(pwd.length<8){setPwdErr('Password too short');valid=false;}else setPwdErr('');

    if(pwd!==pwd2){setMatchErr("Passwords don't match");valid=false;}else setMatchErr('');

    if(!agree){setAgreeErr('You must accept the terms');valid=false;}else setAgreeErr('');

    if(!valid)return;

    setLoading(true);
    setTimeout(() => {
      alert('Registered!');
      navigate('/login');
    }, 800);
  };

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
            boxShadow:
              '0px 1px 1px -0.5px var(--elevationshadow), 0px 3px 3px -1.5px var(--elevationshadow), 0px 24px 24px -12px var(--elevationshadow)',
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            gap: 32,
          }}
        >
          <h1 style={{ fontFamily: 'Coolvetica, sans-serif', fontSize: 64, lineHeight: '48px', color: primary, margin: 0, letterSpacing: '-0.64px' }}>swarmind</h1>
          <p style={{ margin: 0, color: greyText, fontSize: 15 }}>Create your account</p>

          {/* email */}
          <div style={{ width: '100%', display: 'flex', flexDirection: 'column', gap: 4 }}>
            <label style={{ fontSize: 13, fontWeight: 500, color: greyText }}>Email</label>
            <input
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              placeholder="Enter your email"
              style={{
                width: '100%',
                height: 48,
                borderRadius: 12,
                background: emailFocus ? '#FFFFFF' : greyBg,
                border: emailErr ? `1px solid ${errorCol}` : emailFocus ? `1px solid ${primary}` : '1px solid transparent',
                boxShadow: emailFocus && emailErr ? '0 0 0 3px rgba(240,68,56,0.2)' : emailFocus ? '0 0 0 3px rgba(94,100,255,0.2)' : 'none',
                padding: '0 12px',
                fontSize: 15,
                boxSizing: 'border-box',
                outline: 'none',
                transition:'border 0.15s, box-shadow 0.15s, background 0.15s',
              }}
              onFocus={()=>setEmailFocus(true)}
              onBlur={()=>setEmailFocus(false)}
            />
            {emailErr && <div style={{ width: '100%', margin: '2px 0 0 0', display: 'flex', alignItems: 'flex-start' }}><ErrorLine>{emailErr}</ErrorLine></div>}
          </div>
          {/* password */}
          <div style={{ width: '100%', display: 'flex', flexDirection: 'column', gap: 4 }}>
            <label style={{ fontSize: 13, fontWeight: 500, color: greyText }}>Password</label>
            <div style={{ position: 'relative' }}>
              <input
                type={showPwd1 ? 'text' : 'password'}
                value={pwd}
                onChange={(e) => setPwd(e.target.value)}
                placeholder="Enter password"
                style={{
                  width: '100%',
                  height: 48,
                  borderRadius: 12,
                  background: pwd1Focus ? '#FFFFFF' : greyBg,
                  border: pwdErr ? `1px solid ${errorCol}` : pwd1Focus ? `1px solid ${primary}` : '1px solid transparent',
                  boxShadow: pwd1Focus && pwdErr ? '0 0 0 3px rgba(240,68,56,0.2)' : pwd1Focus ? '0 0 0 3px rgba(94,100,255,0.2)' : 'none',
                  padding: '0 44px 0 12px',
                  fontSize: 15,
                  boxSizing: 'border-box',
                  outline: 'none',
                  transition:'border 0.15s, box-shadow 0.15s, background 0.15s',
                }}
                onFocus={()=>setPwd1Focus(true)}
                onBlur={()=>setPwd1Focus(false)}
              />
              <button
                type="button"
                onClick={() => setShowPwd1((p) => !p)}
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
                {showPwd1 ? (
                  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke={pwdErr ? errorCol : greyText} strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M17.94 17.94A10.94 10.94 0 0 1 12 20c-5 0-9.27-3-11-8 1.13-2.9 3-5.23 5.32-6.64" /><path d="M3 3l18 18" /></svg>
                ) : (
                  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke={pwdErr ? errorCol : greyText} strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><circle cx="12" cy="12" r="3" /><path d="M2.05 12a10.97 10.97 0 0 1 19.9 0 10.97 10.97 0 0 1-19.9 0z" /></svg>
                )}
              </button>
            </div>
            {pwdErr && <div style={{ width: '100%', margin: '2px 0 0 0', display: 'flex', alignItems: 'flex-start' }}><ErrorLine>{pwdErr}</ErrorLine></div>}
          </div>
          {/* strength bar */}
          {(
            <div style={{ width:'100%', display:'flex', flexDirection:'column', gap:0, marginTop: -10 }}>
              <div style={{ display:'flex', gap:8 }}>
                {[0,1,2,3].map((i)=>{
                  const strength=getStrength(pwd).level;
                  let color=barEmpty;
                  if(strength===1 && i===0) color=red;
                  if(strength===2 && i<2) color=yellow;
                  if(strength===3 && i<3) color=green;
                  if(strength===4) color=green;
                  return <div key={i} style={{flex:1,height:6,borderRadius:1000,background:color}} />
                })}
              </div>
              <span style={{ fontSize:13,color:greyText, alignSelf:'flex-end' }}>
                {strengthLabels[getStrength(pwd).level-1]?.text} {strengthLabels[getStrength(pwd).level-1]?.emoji}
              </span>
            </div>
          )}
          {/* repeat */}
          <div style={{ width: '100%', display: 'flex', flexDirection: 'column', gap: 4 }}>
            <label style={{ fontSize: 13, fontWeight: 500, color: greyText }}>Repeat password</label>
            <div style={{ position: 'relative' }}>
              <input
                type={showPwd2 ? 'text' : 'password'}
                value={pwd2}
                onChange={(e) => setPwd2(e.target.value)}
                placeholder="Re-enter password"
                style={{
                  width: '100%',
                  height: 48,
                  borderRadius: 12,
                  background: pwd2Focus ? '#FFFFFF' : greyBg,
                  border: matchErr ? `1px solid ${errorCol}` : pwd2Focus ? `1px solid ${primary}` : '1px solid transparent',
                  boxShadow: pwd2Focus && matchErr ? '0 0 0 3px rgba(240,68,56,0.2)' : pwd2Focus ? '0 0 0 3px rgba(94,100,255,0.2)' : 'none',
                  padding: '0 44px 0 12px',
                  fontSize: 15,
                  boxSizing: 'border-box',
                  outline: 'none',
                  transition:'border 0.15s, box-shadow 0.15s, background 0.15s',
                }}
                onFocus={()=>setPwd2Focus(true)}
                onBlur={()=>setPwd2Focus(false)}
              />
              <button
                type="button"
                onClick={() => setShowPwd2((p) => !p)}
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
                {showPwd2 ? (
                  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke={matchErr ? errorCol : greyText} strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><path d="M17.94 17.94A10.94 10.94 0 0 1 12 20c-5 0-9.27-3-11-8 1.13-2.9 3-5.23 5.32-6.64" /><path d="M3 3l18 18" /></svg>
                ) : (
                  <svg width="20" height="20" viewBox="0 0 24 24" fill="none" stroke={matchErr ? errorCol : greyText} strokeWidth="2" strokeLinecap="round" strokeLinejoin="round"><circle cx="12" cy="12" r="3" /><path d="M2.05 12a10.97 10.97 0 0 1 19.9 0 10.97 10.97 0 0 1-19.9 0z" /></svg>
                )}
              </button>
            </div>
            {matchErr && <div style={{ width: '100%', margin: '2px 0 0 0', display: 'flex', alignItems: 'flex-start' }}><ErrorLine>{matchErr}</ErrorLine></div>}
          </div>

          {/* agreement checkbox */}
          <div style={{ width: '100%', display: 'flex', flexDirection: 'row', alignItems: 'left', gap: 2, fontSize: 13, color: greyText, marginTop: -25, marginBottom: 0, padding: 0, paddingLeft: 7 }}>
            <input type="checkbox" checked={agree} onChange={e=>setAgree(e.target.checked)} style={{ marginTop: 2, width: 13, height: 13 ,marginLeft: 3}} />
            <span>
              I have read and accepted the <a href="#" style={{color:primary,textDecoration:'none'}}>Private Policy</a> and <a href="#" style={{color:primary,textDecoration:'none'}}>User Agreement</a>
            </span>
          </div>
          {agreeErr && <div style={{ width: '100%', margin: '-20px 0 0 0', display: 'flex', alignItems: 'flex-start' }}><ErrorLine>{agreeErr}</ErrorLine></div>}

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
              boxShadow:
                '0 1px 1px -0.5px var(--elevationshadow), 0 3px 3px -1.5px var(--elevationshadow), inset 0 3px 4px -3px rgba(255,255,255,0.56), inset 0 0 8px -2px rgba(255,255,255,0.48)',
            }}
          >
            {loading ? 'Creating…' : 'Continue'}
          </button>

          <p style={{ fontSize: 13, color: greyText }}>
            Already have an account?{' '}
            <button
              type="button"
              onClick={() => navigate('/login')}
              style={{ border: 'none', background: 'none', color: primary, cursor: 'pointer', fontSize: 13 }}
            >
              Log in
            </button>
          </p>
        </form>
        </PageWrapper>
      </div>
      <Footer />
    </div>
  );
};

export default RegistrationPage; 