'use client';
import React from 'react';
import { Form, Input, Checkbox, Button, Row, Col } from "antd";
import FormItem from "antd/lib/form/FormItem";
import InputPassword from "antd/lib/input/Password";
import { Amplify } from "aws-amplify"
import awsmobile from '../../aws-exports'

const SignUpForm = ({}) => {

Amplify.configure({ ...awsmobile, ssr: true });

async function onFinish(values) {
  try {
    // Aqui você pode chamar a função para cadastrar o usuário com os dados fornecidos
    console.log('Dados do formulário:', values);

    // Enviar os dados para o Cognito
    const { email, password } = values; // Supondo que os campos de email e senha tenham esses nomes
    const user = await Auth.signUp({
      username: email,
      password: password
      // outros atributos opcionais do usuário podem ser adicionados aqui, como nome, sobrenome, etc.
    });
    console.log('Usuário cadastrado com sucesso:', user);
  } catch (error) {
    console.error('Erro ao cadastrar usuário:', error);
  }
}

const onFinishFailed = (errorInfo) => {
  console.log('Failed:', errorInfo);
};

const validatePassword = (_, value) => {
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

const validateConfirmPassword = (_, value, callback, form) => {
  const { getFieldValue } = form;

  if (value && value !== getFieldValue('password')) {
    callback('As senhas não são iguais!');
  } else {
    callback();
  }
}; 

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
