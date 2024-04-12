'use client'
import React, {useState} from 'react';
import { Form, Input, Checkbox, Button, Row, Col, message } from "antd";
import FormItem from "antd/lib/form/FormItem";
import InputPassword from "antd/lib/input/Password";


const LoginForm = ({ onFinishFailed }) => {
  const [loading, setLoading] = useState(false);

  const handleSignIn = async (values) => {
    try {
      console.log("Dados do formulário:", values);
      const { email, password } = values;

      const url = "http://localhost:8080/users/signin";
      const data = { email, password };
      const response = await fetch(url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });
      if (!response.ok) {
        throw new Error("Problemas no Login");
      }
      const responseData = await response.json();
      
      // Verificar se o token existe
      if (responseData) {
        message.loading("Realizando Login...")
        setLoading(true);
        
        setTimeout(() => {
          message.success("Login Realizado com Sucesso");
        }, 500); 

      setTimeout(() => {
        window.location.href = '/alert';
      }, 2000); // 1000ms = 1 segundo
      } else {
        // Exibir mensagem de erro
        message.error("Usuário não cadastrado, cadastre-se!");
      }
    } catch (error) {
      console.error("Erro ao realizar login:", error);
      throw error;
    }
  };
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
      onFinish={handleSignIn}
      onFinishFailed={onFinishFailed}
      autoComplete="off"
    >
      <h4>Email:</h4>
      <FormItem
      id="Email"
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