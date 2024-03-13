import React from 'react';
import Image from 'next/image';
import { Col, Row } from "antd";
import SignUpForm from '@/components/SignUpForm/SignUpForm';

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

const validateConfirmPassword = (_, value, callback) => {
  const { getFieldValue } = form;

  if (value && value !== getFieldValue('password')) {
    callback('As senhas não são iguais!');
  } else {
    callback();
  }
};

const SignUpPage = ({ onFinish, onFinishFailed, validatePassword, validateConfirmPassword}) => {
  return (
    <Row justify='center'>
      <Col span={12}>
        <Image width={100} src={logo}/>
          <SignUpForm onFinish={onFinish} onFinishFailed={onFinishFailed} validatePassword={validatePassword} validateConfirmPassword={validateConfirmPassword}/>
      </Col>
    </Row>
  );
};

export default SignUpPage;