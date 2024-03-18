import React from 'react';

export default function SensorText() {
    return (
      <div>
        <div>
            <h1 style={{ 
            fontWeight: 'bold', 
            fontSize: '2.8rem', 
            fontFamily: "'Space Mono', monospace", 
            marginBottom: '1rem', 
            position: 'relative',
            width: '70%',
            zIndex: 1, 
            color: '#2E329B'
        }}>
            Expandindo nosso sistema
        </h1>
        <p style={{ 
            fontWeight: '300', 
            fontSize: '1.25rem',
            fontFamily: "'Montserrat', sans-serif",
            width: '70%',
            paddingTop: '3rem',
            boxSizing: 'border-box'
        }}>
            Aqui você pode adicionar novos sensores para serem monitorados pelo nosso sistema. 
        </p>
    </div>
      </div>
    );
  }


