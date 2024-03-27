import React from 'react';
import { Col, Row } from "antd";
import axios from 'axios'; // Importe o Axios para fazer solicitações HTTP
import AlertForm from '@/components/Alert/AlertForm';


const AlertPage = ({}) => {
  return (
    <div style={{ backgroundColor: 'white', minHeight: '100vh', display: 'flex', flexDirection: 'column', justifyContent: 'center', alignItems: 'center' }}>
      <Row justify='center' style={{ width: '100%' }}>
        <Col span={10}>
          <h1 style={{ 
            fontWeight: 'bold', 
            fontSize: '2.8rem', 
            fontFamily: "'Space Mono', monospace", 
            marginBottom: '1rem', 
            position: 'relative',
            width: '70%',
            color: '#2E329B'
          }}>
            Contribua com nossos dados
          </h1>
          <p style={{ 
            fontWeight: '300', 
            fontSize: '1.25rem',
            fontFamily: "'Montserrat', sans-serif",
            width: '70%',
            paddingTop: '3rem',
            boxSizing: 'border-box'
          }}>
            Para nos ajudar a ter uma contribuição na sociedade mais constante, insira sua localização e o acidente que você presenciou, para que possamos alertar outras pessoas.
          </p>
        </Col>
        <Col span={12}>
          <h1 style={{ fontWeight: 'bold', fontSize: '2rem', fontFamily: "'Space Mono', monospace", marginBottom: '1rem', color: '#2E329B' }}>Crie seu Alerta</h1>
          {/* Passe a função handleSubmit como prop para o componente AlertForm */}
          <AlertForm/>
        </Col>
      </Row>
    </div>
  );
};

export default AlertPage;
