import React, { useState } from 'react';
import { Form, Input, Button } from "antd";
import axios from 'axios';
import FormItem from "antd/lib/form/FormItem";

const SensorForm = () => {
  const [successMessage, setSuccessMessage] = useState("");
  const onFinishFailed = () => {};

  const handleSubmit = async (values) => {
    try {
      const { name, latitude, longitude, min, max} = values;

      const data = 
        {
          "name": name,
          "latitude": latitude,
          "longitude": longitude,
          "params": {
            "co2": {
              "min": max,
              "max": min,
              "z": 0
            },
            "co":{
              "min": 0,
              "max": 0,
              "z": 0
            },
            "no2": {
              "min": 0,
              "max": 0,
              "z": 0
            },
            "mp10": {
              "min": 0,
              "max": 0,
              "z": 0
            },
            "mp25": {
              "min": 0,
              "max": 0,
              "z": 0
            },
            "rad": {
              "min": 0,
              "max": 0,
              "z": 0
            }
          }
        }

      console.log(data.params)
      const response = await axios.post('http://localhost:8080/sensors', data);

      if (response.status === 200) {
        setSuccessMessage('Sensor criado com sucesso!');
      }
    } catch (error) {
      console.error('Erro ao criar sensor:', error);
    }
  };

  return (
    <div>
      {successMessage && <p style={{ color: 'green' }}>{successMessage}</p>}
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
        onFinish={handleSubmit}
        onFinishFailed={onFinishFailed}
        autoComplete="off"
      >
        <p>Qual o nome do sensor?</p>
        <FormItem
          name="name"
          rules={[
            {
              required: true,
              message: 'Por favor, digite o nome do seu sensor!',
            },
          ]}
        >
          <Input className='w-full' />
        </FormItem>

        <p>Indique sua Latitude:</p>
        <FormItem
          name="latitude"
          rules={[
            {
              required: true,
              message: 'Por favor, insira sua latitude!',
            },
          ]}
        >
          <Input className='w-full' />
        </FormItem>

        <p>Indique sua Longitude:</p>
        <FormItem
          name="longitude"
          rules={[
            {
              required: true,
              message: 'Por favor, digite sua longitude!',
            },
          ]}
        >
          <Input className='w-full' />
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
          <Input className='w-full' />
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
          <Input className='w-full' />
        </FormItem>

        <FormItem
          wrapperCol={{
            span: 16,
          }}
        >
          <Button type="primary" htmlType="submit" style={{ width: '100%', backgroundColor: '#FFA13A', color: 'white' }}>
            Criar sensor!
          </Button>
        </FormItem>
      </Form>
    </div>
  );
};

export default SensorForm;
