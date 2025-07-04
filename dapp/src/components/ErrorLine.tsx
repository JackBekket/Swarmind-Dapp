import React from 'react';

const ErrorLine: React.FC<{ children: React.ReactNode }> = ({ children }) => (
  <div
    style={{
      width: 416,
      minHeight: 24,
      display: 'flex',
      alignItems: 'center',
      gap: 2,
      padding: '0 4px',
      color: '#F04438',
      fontSize: 15,
      fontWeight: 400,
      borderRadius: 6,
      background: 'transparent',
      margin: '2px 0',
    }}
  >
    <img src="/icons/close.svg" alt="error" style={{ width: 18, height: 18, flexShrink: 0 }} />
    {children}
  </div>
);

export default ErrorLine; 