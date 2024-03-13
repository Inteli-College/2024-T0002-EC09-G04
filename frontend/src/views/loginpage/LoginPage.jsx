import React from 'react';
import { Col, Row, Image } from "antd";
import LoginForm from '@/components/LoginForm/LoginForm';

import logo from '../../assets/img/id_visual.png';

function onFinish(values) {
  console.log('Success:', values);
};

function onFinishFailed(errorInfo) {
  console.log('Failed:', errorInfo);
};

const LoginPage = ({ onFinish, onFinishFailed }) => {
  return (
    <Row justify='center'>
      <Col span={16}>
        <Image preview={false}
          width={200}
          src={logo}
        />
        <LoginForm onFinish={onFinish} onFinishFailed={onFinishFailed} />
      </Col>
    </Row>
  );
};

export default LoginPage;