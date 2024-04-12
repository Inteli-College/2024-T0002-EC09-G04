'use client'
import React, { useState } from 'react';
import { Form, Input, Button, Select, message } from "antd";
import axios from 'axios'; // Importe o Axios para fazer solicitações HTTP
import FormItem from "antd/lib/form/FormItem";
import dynamic from 'next/dynamic';

const MapComponentWithNoSSR = dynamic(() => import('@/components/Map/Map'), {
  ssr: false, // Desativa a renderização do lado do servidor para este componente
});

const Alert = () => {
  const onFinishFailed = () => {}; // Defina onFinishFailed

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

  const handleSubmit = async (values) => {
    try {
      const {Acidente } = values;

      const data = {
        "latitude": parseFloat(latitude),
        "longitude": parseFloat(longitude),
        "option": Acidente,
      };

      // Envia os dados para o backend usando Axios
      const response = await axios.post('http://localhost:8080/alerts', data);

      // Verifica se a solicitação foi bem-sucedida
      if (response.status === 201) {
        message.success("Alerta criado com sucesso!");
      }
    } catch (error) {
      console.error('Erro ao enviar alerta:', error);
      // Trate o erro conforme necessário
    }
  };

  return (
    <div>
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
                  name="Latitude"
                  rules={[
                    {
                      required: true,
                      message: 'Por favor, insira sua latitude!',
                    },
                  ]}
                >
                  <Input name="Latitude" style={{ width: '100%' }} onChange={(e) => setLatitude(e.target.value)}/>
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
                  <Input name="Longitude" style={{ width: '100%' }} onChange={(e) => setLongitude(e.target.value)} />
                </FormItem>
              </div>
            }
          </>
        )}

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
            initialvalues={{ name: "" }}
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
          <Button id="submit" htmlType="submit" style={{ width: '100%', backgroundColor: '#FFA13A !important', color: 'white' }}>
            Enviar Alerta !
          </Button>
        </FormItem>
      </Form>
    </div>
  );
};

export default Alert;
