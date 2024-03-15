'use client'
import React, { useState } from 'react';
import { Form, Input, Button, Select } from "antd";
import axios from 'axios'; // Importe o Axios para fazer solicitações HTTP
import FormItem from "antd/lib/form/FormItem";

const Alert = () => {
  const [successMessage, setSuccessMessage] = useState("");
  const onFinishFailed = () => {}; // Defina onFinishFailed

  const handleSubmit = async (values) => {
    try {
      // Extrai os valores do formulário
      const { Latitude, Longitude, Acidente } = values;

      // Cria um objeto com os dados no formato esperado pelo backend
      const data = {
        "latitude": parseFloat(Latitude),
        "longitude": parseFloat(Longitude),
        "option": Acidente,
      };

      // Envia os dados para o backend usando Axios
      const response = await axios.post('http://localhost:8080/alerts', data);

      // Verifica se a solicitação foi bem-sucedida
      if (response.status === 200) {
        setSuccessMessage('Alerta enviado com sucesso!'); // Define a mensagem de sucesso no estado
        // Faça qualquer outra ação necessária após o envio bem-sucedido
      }
    } catch (error) {
      console.error('Erro ao enviar alerta:', error);
      // Trate o erro conforme necessário
    }
  };

  return (
    <div>
      {successMessage && <p style={{ color: 'green' }}>{successMessage}</p>} {/* Exibe a mensagem de sucesso se houver */}
      <Form
        name="basic"
        labelCol={{
          span: 8,
        }}
        wrapperCol={{
          span: 16,
        }}
        style={{
          maxWidth: 400,
        }}
        initialValues={{
          remember: true,
        }}
        onFinish={handleSubmit} // Use a função handleSubmit para lidar com o envio do formulário
        onFinishFailed={onFinishFailed}
        autoComplete="off"
      >
        <p>Indique sua Latitude:</p>
        <FormItem
          name="Latitude"
          rules={[
            {
              required: true,
              message: 'Por favor, insira sua latitude!',
            },
          ]}
        >
          <Input name="Latitude" style={{ width: '100%' }} />
        </FormItem>

        <p>Indique sua Longitude:</p>
        <FormItem
          name="Longitude"
          rules={[
            {
              required: true,
              message: 'Por favor, digite sua longitude!',
            },
          ]}
        >
          <Input name="Longitude" style={{ width: '100%' }} />
        </FormItem>

        <p>Qual o tipo de alerta ? </p>
        <FormItem
          name="Acidente"
          rules={[
            {
              required: true,
              message: 'Por favor, informe o acidente!',
            },
          ]}
        >
          <Select
            defaultValue=""
            style={{ width: '100%' }}
            options={[
              {
                value: 'Alagamento',
                label: 'Alagamento',
              },
              {
                value: 'Congestionamento',
                label: 'Congestionamento',
              },
              {
                value: 'Deslizamento',
                label: 'Deslizamento',
              },
              {
                value: 'Incêndio',
                label: 'Incêndio',
              },
              {
                value: 'Obra',
                label: 'Obra',
              },
            ]}
          />
        </FormItem>

        <FormItem
          wrapperCol={{
            span: 16,
          }}
        >
          <Button type="primary" id="submit" htmlType="submit" style={{ width: '100%', backgroundColor: '#FFA13A !important', color: 'white' }}>
            Enviar Alerta !
          </Button>
        </FormItem>
      </Form>
    </div>
  );
};

export default Alert;
