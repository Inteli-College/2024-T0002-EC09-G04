import React from 'react';
import { Form, Input, Checkbox, Button, Row, Col } from "antd";
import FormItem from "antd/lib/form/FormItem";
import InputPassword from "antd/lib/input/Password";


const LoginForm = ({ onFinish, onFinishFailed }) => {
  return (
    <Col span={10}>
    <Form
      style={{
        color:'white',
        fontFamily: "'Poppins', sans-serif",  
      }}
      name="basic"
      initialValues={{
        remember: true,
      }}
      onFinish={onFinish}
      onFinishFailed={onFinishFailed}
      autoComplete="off"
    >
      <h4>Email:</h4>
      <FormItem
      id="Email"
        name="Email"
        rules={[
          {
            required: true,
            message: 'Por favor insira seu email!',
          },
        ]}
      >
        <Input />
      </FormItem>
      <h4>Senha:</h4>
      <FormItem
      id="password"
        name="password"
        rules={[
          {
            required: true,
            message: 'Por favor digite sua senha!',
          },
        ]}
      >
        <InputPassword />
      </FormItem>
      <Row justify='center'>
      <FormItem
        wrapperCol={{
          span: 16,
        }}
      >
        <Button type="primary" id="submit" htmlType="submit" width='full' style={{ width: 650 }}>
          Submit
        </Button>
      </FormItem>
      </Row>
      <p style={{textAlign:'center'}}>Esqueceu a Senha? <a>Resetar</a></p>
        <p style={{textAlign:'center'}}>Ainda não tem seus acessos? <a href='/signup'>Cadastre-se</a></p>
    </Form>
    </Col>
  );
};

export default LoginForm;