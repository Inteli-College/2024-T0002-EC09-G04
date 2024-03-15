import React, { useState } from 'react';
import { Form, Input, Button } from "antd";
import axios from 'axios';
import FormItem from "antd/lib/form/FormItem";


const SensorForm = () => {
  const [formData, setFormData] = useState({
    Nome: '',
    lat: '',
    long: '',
    tipo: '',
    min: '',
    max: ''
  });

  const onFinish = async (values) => {
    try {
      // Envia os dados para a API
      await axios.post('sua_api_aqui', formData);
      console.log('Dados enviados com sucesso!');
    } catch (error) {
      console.error('Erro ao enviar os dados:', error);
    }
  };

  const onFinishFailed = (errorInfo) => {
    console.log('Failed:', errorInfo);
  };

  const handleChange = (event) => {
    const { name, value } = event.target;
    setFormData({
      ...formData,
      [name]: value
    });
  };

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
        maxWidth: 400,
      }}
      initialValues={{
        remember: true,
      }}
      onFinish={onFinish}
      onFinishFailed={onFinishFailed}
      autoComplete="off"
    >
      <p>Qual o nome do sensor?</p>
      <Form.Item
        name="Nome"
        rules={[
          {
            required: true,
            message: 'Por favor, digite o nome do seu sensor!',
          },
        ]}
      >
        <Input className='w-full' name="Nome" onChange={handleChange} />
      </Form.Item>


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
        <Input className='w-full' name="Nome" onChange={handleChange} />

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
        <Input className='w-full' name="Nome" onChange={handleChange} />
    </FormItem>

    <p>Insira o tipo do sensor:</p>
    <FormItem
      name="tipo"
      rules={[
        {
          required: true,
          message: 'Por favor, insira o tipo de sensor!',
        },
      ]}
    >
        <Input className='w-full' name="Nome" onChange={handleChange} />
    </FormItem>

    <p>Insira o mínimo de captura do sensor:</p>
    <FormItem
      name="min"
      rules={[
        {
          required: true,
          message: 'Por favor, insira o valor mínimo do sensor',
        },
      ]}
    >
        <Input className='w-full' name="Nome" onChange={handleChange} />
    </FormItem>

    <p>Insira o máximo de captura do sensor:</p>
    <FormItem
      name="max"
      rules={[
        {
          required: true,
          message: 'Por favor, insira o valor máximo do sensor',
        },
      ]}
    >
      <Input className='w-full' name="Nome" onChange={handleChange} />
    </FormItem>
      <Form.Item
        wrapperCol={{
          span: 16,
        }}
      >
        <Button type="primary" htmlType="submit" style={{ width: '100%', backgroundColor: '#FFA13A', color: 'white' }}>
          Criar sensor!
        </Button>
      </Form.Item>
    </Form>
  );
};

export default SensorForm;
