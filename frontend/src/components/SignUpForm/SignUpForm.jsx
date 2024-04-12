"use client";
import React, { useState, useEffect } from "react";
import { Form, Input, Button, Row, Col, message, Modal, Space } from "antd";
// import { Amplify } from "aws-amplify"; // Importar o Auth do aws-amplify
// import { signUp } from "aws-amplify/auth";

// import awsmobile from "../../aws-exports";

//Amplify.configure({
//  Auth: {
//  region: awsmobile.region,
// userPoolId: awsmobile.userPoolId,
//  userPoolWebClientId: awsmobile.userPoolWebClientId,
//},
//})

const SignUpForm = () => {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [user, setUser] = useState({});
  const [verificationCode, setVerificationCode] = useState("");

  const [form] = Form.useForm(); // Adicione esta linha

  useEffect(() => {
    if (isModalOpen) {
      form.resetFields();
    }
  }, [isModalOpen, form]);

  const handleOk = async () => {
    try {
      await handleVerificationCode(verificationCode);
    } catch (error) {
      console.log("Failed:", error);
    }
  };

  const handleVerificationCode = async (code) => {
    try {
      console.log("Código de Verificação:", code); // Exibe corretamente o código de verificação
      const { email } = user;
      const url = "http://localhost:8080/users/confirmation";
      const data = { email, code };
      const response = await fetch(url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });
      if (!response.ok) {
        throw new Error("Problemas no Código de Verificação");
      }
      const responseData = await response.json();
      console.log("Verificação Realizada com Sucesso", responseData);
      message.success("Sua conta já está disponível para o Login !");
      setIsModalOpen(false);

      window.location.href = "/login"; // Redireciona para a página de login
      return responseData;
    } catch (error) {
      console.error("Erro ao validar Código:", error);
      throw error;
    }
  };

  const handleCancel = () => {
    setIsModalOpen(false);
  };

  const handleSignUp = async (values) => {
    try {
      console.log("Dados do formulário:", values);
      const { name, email, password } = values;

      setUser({
        name,
        email,
        password
      });

      const url = "http://localhost:8080/users/signup";
      const data = { name, email, password };
      const response = await fetch(url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
      });
      if (!response.ok) {
        throw new Error("Erro ao criar usuário");
      }
      const responseData = await response.json();
      console.log("Usuário criado com sucesso:", responseData);
      message.success("Usuário Criado com Sucesso !");
      setIsModalOpen(true); // Abra o modal após o cadastro bem-sucedido
      return responseData;
    } catch (error) {
      console.error("Erro ao criar usuário:", error);
      throw error;
    }
  };

  console.log(user);

  const onFinishFailed = (errorInfo) => {
    console.log("Failed:", errorInfo);
  };

  return (
    <>
      <Form
        name="basic"
        style={{ color: "white" }}
        initialValues={{
          remember: true,
        }}
        onFinish={handleSignUp}
        onFinishFailed={onFinishFailed}
        autoComplete="off"
      >
        <h4>Nome Completo:</h4>
        <Form.Item
          style={{ width: "99%" }}
          name="name"
          rules={[
            {
              required: true,
              message: "Por favor insira seu nome completo!",
            },
          ]}
        >
          <Input />
        </Form.Item>
        <h4>Email:</h4>
        <Form.Item
          style={{ width: "99%" }}
          name="email"
          rules={[
            {
              required: true,
              message: "Por favor insira seu email!",
            },
          ]}
        >
          <Input />
        </Form.Item>
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
            <Form.Item
              style={{ width: "98%" }}
              name="password"
              rules={[
                {
                  required: true,
                  message: "Por favor digite sua senha!",
                },
              ]}
            >
              <Input.Password />
            </Form.Item>
          </Col>
          <Col span={12}>
            <Form.Item
              style={{ width: "98%" }}
              name="confirmPassword"
              rules={[
                {
                  required: true,
                  message: "Por favor confirme sua senha!",
                },
              ]}
            >
              <Input.Password />
            </Form.Item>
          </Col>
        </Row>
        <Form.Item
          wrapperCol={{
            span: 16,
          }}
        >
          <Button type="primary" htmlType="submit" style={{ width: "100%" }}>
            Submit
          </Button>
        </Form.Item>
      </Form>
      <Modal
        title="Verificação de Conta"
        open={isModalOpen}
        onOk={handleOk}
        onCancel={handleCancel}
        footer={[
          <Button 
            type="Dashed"
            key="back"
            >
            Cancelar
          </Button>,
          <Button 
            key="submit"
            onClick={handleOk}
            style={{backgroundColor: '#FFA13A !important'}}
          >
            Enviar Código
          </Button>,
        ]}
      >
        <p>
          Seu registro foi criado com sucesso {user.name}, pedimos que você
          nos confirme o código de verificação que foi enviado para o seu email
        </p>
        <Form
          form={form} // Passe o objeto de formulário para o formulário interno
          name="verification"
          style={{ color: "white" }}
          initialValues={{
            remember: true,
          }}
          autoComplete="off"
        >
          <Space direction="vertical">
          <h5>Insira o Código de Verificação enviado para Seu Email:</h5>
          <Form.Item name="verificationCode">
            <Input
              placeholder="Digite o código de verificação"
              onChange={(e) => setVerificationCode(e.target.value)}
            />
          </Form.Item>
          </Space>
        </Form>
      </Modal>
    </>
  );
};

export default SignUpForm;
