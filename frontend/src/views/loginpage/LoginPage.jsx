import React from 'react';
import Image from 'next/image';
import { Col, Row } from "antd";
import LoginForm from '@/components/LoginForm/LoginForm';

import logo from '../../assets/img/id_visual.png';

function onFinish(values) {
  console.log('Success:', values);
};

function onFinishFailed(errorInfo) {
  console.log('Failed:', errorInfo);
};

function validatePassword(_, value){
  // Verifica se a senha tem pelo menos uma letra maiúscula, uma letra minúscula e um número
  const upperCaseRegex = /[A-Z]/;
  const lowerCaseRegex = /[a-z]/;
  const numberRegex = /[0-9]/;

  if (
    !upperCaseRegex.test(value) ||
    !lowerCaseRegex.test(value) ||
    !numberRegex.test(value)
  ) {
    return Promise.reject(
      'A senha deve conter pelo menos uma letra maiúscula, uma letra minúscula e um número!'
    );
  }

  return Promise.resolve();
};

const LoginPage = ({ onFinish, onFinishFailed, validatePassword}) => {
  return (
    <Row justify='center'>
      <Col span={16}>
        <Image width={100} src={logo}/>
        <LoginForm onFinish={onFinish} onFinishFailed={onFinishFailed}/>
      </Col>
    </Row>
  );
};

export default LoginPage;