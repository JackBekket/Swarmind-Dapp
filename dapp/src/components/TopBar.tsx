import React from 'react';

const primary = '#5E64FF';

const TopBar: React.FC = () => (
  <header
    style={{
      height: 56,
      borderBottom: '1px solid #F4F4F6',
      display: 'flex',
      alignItems: 'center',
      padding: '0 16px',
      flexShrink: 0,
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
);

export default TopBar; 