import React from 'react';
import { Form, Input, Checkbox, Button, Row, Col } from "antd";
import FormItem from "antd/lib/form/FormItem";
import InputPassword from "antd/lib/input/Password";

const SignUpForm = ({ onFinish, onFinishFailed, validateConfirmPassword, validatePassword }) => {
  return (
    <Form
      name="basic"
      initialValues={{
        remember: true,
      }}
      onFinish={onFinish}
      onFinishFailed={onFinishFailed}
      autoComplete="off"
    >
      <h5 level={5}>Email:</h5>
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
          <h5 level={5}>Senha:</h5>
        </Col>
        <Col>
          <h5 level={5}>Confirme sua senha:</h5>
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
        <Button type="primary" htmlType="submit" width="full" style={{ width: '100%' }}>
          Submit
        </Button>
      </FormItem>
    </Form>
  );
};

export default SignUpForm;
