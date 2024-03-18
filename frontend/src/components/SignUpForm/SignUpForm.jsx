import React from 'react';
import { Form, Input, Checkbox, Button, Row, Col } from "antd";
import FormItem from "antd/lib/form/FormItem";
import InputPassword from "antd/lib/input/Password";

const SignUpForm = ({ onFinish, onFinishFailed, validateConfirmPassword, validatePassword }) => {
  return (
    <Form
      name="basic"
      style={{color:'white'}}
      initialValues={{
        remember: true,
      }}
      onFinish={onFinish}
      onFinishFailed={onFinishFailed}
      autoComplete="off"
    >
      <h4>Email:</h4>
      <FormItem
        style={{width: '99%'}}
        name="email"
        rules={[
          {
            required: true,
            message: 'Por favor insira seu email!',
          },
        ]}
      >
        <Input />
      </FormItem>
      <Row>
        <Col span={12}>
          <h4 level={5}>Senha:</h4>
        </Col>
        <Col>
          <h4 level={5}>Confirme sua senha:</h4>
        </Col>
      </Row>
      <Row>
        <Col span={12}>
          <FormItem
            style={{width: '98%'}}
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
        </Col>
        <Col span={12}>
          <FormItem
            style={{width: '98%'}}
            name="confirmPassword"
            dependencies={['password']}
            rules={[
              {
                required: true,
                message: 'Por favor confirme sua senha!',
              },
              {
                validator: validateConfirmPassword,
              },
            ]}
          >
            <InputPassword />
          </FormItem>
        </Col>
      </Row>
      <FormItem
        wrapperCol={{
          span: 16,
        }}
      >
        <Button type="primary" htmlType="submit" width="full" style={{ width: 650 }}>
          Submit
        </Button>
      </FormItem>
    </Form>
  );
};

export default SignUpForm;
