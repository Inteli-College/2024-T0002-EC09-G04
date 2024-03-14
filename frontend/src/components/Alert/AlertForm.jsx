import React from 'react';
import { Form, Input, Button, Select } from "antd";
import FormItem from "antd/lib/form/FormItem";


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
      maxWidth: 400,
    }}
    initialValues={{
      remember: true,
    }}
    onFinish={onFinish}
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
      <Input style={{ width: '100%' }} />
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
      <Input style={{ width: '100%' }} />
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
      <Button type="primary" htmlType="submit" style={{ width: '100%', backgroundColor: '#FFA13A !important', color: 'white' }}>
        Enviar Alerta !
      </Button>
    </FormItem>
  </Form>
  );
};

export default LoginForm;