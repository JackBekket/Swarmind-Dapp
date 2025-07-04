import React from 'react';

const greyText = '#8C929C';

const Footer: React.FC = () => (
  <footer
    style={{
      height: 72,
      background: '#FFFFFF',
      display: 'flex',
      justifyContent: 'center',
      alignItems: 'center',
      gap: 8,
      flexShrink: 0,
    }}
  >
    <a href="#" style={{ fontSize: 12, color: greyText, textDecoration: 'none' }}>Data Policy</a>
    <div style={{ width: 1, height: 16, background: '#F4F4F6' }} />
    <a href="#" style={{ fontSize: 12, color: greyText, textDecoration: 'none' }}>User Agreement</a>
  </footer>
);

export default Footer; 