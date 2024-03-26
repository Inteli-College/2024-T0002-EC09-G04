'use client';
import React, { useState }  from 'react';
import Image from 'next/image';
import { Col, Row } from 'antd';
import SignUpForm from '@/components/SignUpForm/SignUpForm';

import logo from '../../assets/img/id_visual.png';

const SignUpPage = ({}) => {
  return (
    <div style={{ backgroundColor: '#2E329B', height: '100vh', padding: '10%' }}>
      <Col>
      </Col>
      <Col>
        <Row justify='center'>
          <Image width={100} src={logo} alt="Logo Projeto" />
        </Row>
        <Row justify='center'>
          <SignUpForm/>
        </Row>
        <p style={{textAlign:'center', color:'white'}}>Já possui cadastro? <a href='/login'>Faça o Login.</a></p>
      </Col>
    </div>
  );
};

export default SignUpPage;
