import React from 'react';
import { Form, Input, Checkbox, Button } from "antd";
import FormItem from "antd/lib/form/FormItem";
import InputPassword from "antd/lib/input/Password";
import Title from "antd/lib/typography/title";

const LoginForm = ({onFinish, onFinishFailed}) =>{
  return (
    <Form
    name="basic"
    labelCol={{
      span: 8,
    }}
    wrapperCol={{
      span: 16,
    }}
    style={{
      maxWidth: 600,
    }}
    initialValues={{
      remember: true,
    }}
    onFinish={onFinish}
    onFinishFailed={onFinishFailed}
    autoComplete="off"
  >
    <Title level={5}>Email:</Title>
    <FormItem
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
    <Title level={5}>Senha:</Title>
    <FormItem
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

    <FormItem
      name="remember"
      valuePropName="checked"
      wrapperCol={{
        span: 16,
      }}
    >
      <Checkbox>Remember me</Checkbox>
    </FormItem>

    <FormItem
      wrapperCol={{
        span: 16,
      }}
    >
      <Button type="primary" htmlType="submit" width='full'  style={{ width: '100%' }}>
        Submit
      </Button>
    </FormItem>
  </Form>
  );
};

export default LoginForm;