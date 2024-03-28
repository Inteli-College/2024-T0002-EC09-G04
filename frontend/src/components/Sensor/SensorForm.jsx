import React, { useState } from 'react';
import { Form, Input, Button, Select } from "antd";
import axios from 'axios';
import FormItem from "antd/lib/form/FormItem";
import dynamic from 'next/dynamic';

const MapComponentWithNoSSR = dynamic(() => import('@/components/Map/Map'), {
  ssr: false, // Desativa a renderização do lado do servidor para este componente
});

const SensorForm = () => {
  const [successMessage, setSuccessMessage] = useState("");
  const [showMap, setShowMap] = useState(false);

  const [location, setLocation] = useState({ lat: -23.55052, lng: -46.633308 });

  const [latitude, setLatitude] = useState('');
  const [longitude, setLongitude] = useState('');

  const handleLocationSelect = (latlng) => {
    setLocation(latlng);
    setLatitude(latlng.lat);
    setLongitude(latlng.lng);
    console.log('Localização selecionada:', latlng);
  };

  const onFinishFailed = () => {};

  const handleSubmit = async (values) => {
    try {
      const { name, min, max, obj, z} = values;

      const data = 
        {
          "name": name,
          "latitude": parseFloat(latitude),
          "longitude": parseFloat(longitude),
          "params": {
            [obj]: {
              "min": parseFloat(min),
              "max": parseFloat(max),
              "z": parseFloat(z),
            },
          }
        }

      console.log(data.params)
      const response = await axios.post('http://localhost:8000/sensors', data);

      if (response.status === 201) {
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

        {!showMap && <Button onClick={() => setShowMap(true)} style={{ marginBottom: 16 }}>Indicar Local</Button>}
        
        {showMap && <Button onClick={() => setShowMap(false)} style={{ marginBottom: 16 }}>Informar Valores</Button>}

        {showMap ? (
          <div>
            <MapComponentWithNoSSR onLocationSelect={handleLocationSelect} />
          </div>
        ) : (
          <>
            {
              <div>
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
                  <Input className='w-full' onChange={(e) => setLatitude(e.target.value)}/>
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
                  <Input className='w-full' onChange={(e) => setLongitude(e.target.value)}/>
                </FormItem>
              </div>
            }
          </>
        )}

        <p>Sensor que capta:</p>
        <FormItem
          name="obj"
          rules={[
            {
              required: true,
              message: 'Por favor, informe o tipo!',
            },
          ]}
        >
          <Select
            initialvalues={{ name: "" }}
            style={{ width: '100%' }}
            options={[
              {
                value: 'CO2',
                label: 'CO2',
              },
              {
                value: 'CO',
                label: 'CO',
              },
              {
                value: 'NO2',
                label: 'NO2',
              },
              {
                value: 'MP10',
                label: 'MP10',
              },
              {
                value: 'MP25',
                label: 'MP25',
              },
            ]}
          />
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

        <p>Insira intervalo de confiança:</p>
        <FormItem
          name="z"
          rules={[
            {
              required: true,
              message: 'Por favor, insira o intervalo de confiança',
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
