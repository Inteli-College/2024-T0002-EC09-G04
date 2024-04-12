import React, { useState } from 'react';
import { Form, Input, Button, Select, message, Row, Col } from "antd";
import axios from 'axios';
import dynamic from 'next/dynamic';

const MapComponentWithNoSSR = dynamic(() => import('@/components/Map/Map'), {
  ssr: false, // Isso desativa a renderização do lado do servidor para este componente
});

const SensorForm = () => {
  const [form] = Form.useForm();
  const [showMap, setShowMap] = useState(false);

  const handleLocationSelect = (latlng) => {
    form.setFieldsValue({
      latitude: latlng.lat,
      longitude: latlng.lng,
    });
    console.log('Localização Selecionada:', latlng);
  };

  const onFinish = async (values) => {
    try {
      const { name, latitude, longitude, sensors } = values;

      const params = sensors.reduce((acc, sensor) => {
        const { obj, min, max, z } = sensor;
        acc[obj] = {
          min: parseFloat(min),
          max: parseFloat(max),
          z: parseFloat(z),
        };
        return acc;
      }, {});

      const payload = {
        name,
        latitude: parseFloat(latitude),
        longitude: parseFloat(longitude),
        params,
      };

      console.log('Payload:', payload);

      const response = await axios.post('http://localhost:8080/sensors', payload);

      if (response.status === 201) {
        message.success('Sensor criado com sucesso!');
      }
    } catch (error) {
      console.error('Falha:', error);
      message.error('Falha ao criar sensor.');
    }
  };

  const onFinishFailed = (errorInfo) => {
    console.log('Falha:', errorInfo);
  };

  return (
    <div>
      <Form
        form={form}
        name="dynamic_sensor_form"
        initialValues={{ remember: true }}
        onFinish={onFinish}
        onFinishFailed={onFinishFailed}
        autoComplete="off"
        layout="vertical"
      >
        <Form.Item
          name="name"
          label="Nome do Sensor"
          rules={[{ required: true, message: 'Por favor, insira o nome do seu sensor!' }]}
          style={{ margin: '10px 0' }}
        >
          <Input />
        </Form.Item>

        <Form.Item
          name="latitude"
          label="Latitude"
          rules={[{ required: true, message: 'Por favor, insira a latitude!' }]}
          style={{ margin: '10px 0' }}
        >
          <Input />
        </Form.Item>

        <Form.Item
          name="longitude"
          label="Longitude"
          rules={[{ required: true, message: 'Por favor, insira a longitude!' }]}
          style={{ margin: '10px 0' }}
        >
          <Input />
        </Form.Item>

        {!showMap && <Button onClick={() => setShowMap(true)} style={{ margin: '10px 0' }}>Selecionar no mapa</Button>}
        {showMap && (
          <>
            <MapComponentWithNoSSR onLocationSelect={handleLocationSelect} />
            <Button onClick={() => setShowMap(false)} style={{ margin: '10px 0' }}>
              Fechar mapa
            </Button>
          </>
        )}

        <Form.List
          name="sensors"
          rules={[
            {
              validator: async (_, sensors) => {
                if (!sensors || sensors.length < 1) {
                  return Promise.reject(new Error('Por favor, adicione pelo menos um tipo de sensor.'));
                }
              },
            },
          ]}
        >
          {(fields, { add, remove }) => (
            <>
              {fields.map((field, index) => (
                <Row key={field.key} gutter={16} style={{ marginBottom: 20 }}>
                  <Col span={6}>
                    <Form.Item
                      label="Tipo de Sensor"
                      name={[field.name, 'obj']}
                      fieldKey={[field.fieldKey, 'obj']}
                      rules={[{ required: true, message: 'Por favor, selecione o tipo de sensor.' }]}
                    >
                      <Select placeholder="Selecione um tipo de sensor">
                        <Select.Option value="CO2">CO2</Select.Option>
                        <Select.Option value="CO">CO</Select.Option>
                        <Select.Option value="NO2">NO2</Select.Option>
                        <Select.Option value="MP10">MP10</Select.Option>
                        <Select.Option value="MP25">MP2.5</Select.Option>
                      </Select>
                    </Form.Item>
                  </Col>

                  <Col span={6}>
                    <Form.Item
                      label="Valor Mínimo"
                      name={[field.name, 'min']}
                      fieldKey={[field.fieldKey, 'min']}
                      rules={[{ required: true, message: 'Por favor, insira o valor mínimo.' }]}
                    >
                      <Input type="number" />
                    </Form.Item>
                  </Col>

                  <Col span={6}>
                    <Form.Item
                      label="Valor Máximo"
                      name={[field.name, 'max']}
                      fieldKey={[field.fieldKey, 'max']}
                      rules={[{ required: true, message: 'Por favor, insira o valor máximo.' }]}
                    >
                      <Input type="number" />
                    </Form.Item>
                  </Col>

                  <Col span={6}>
                    <Form.Item
                      label="Intervalo de Confiança"
                      name={[field.name, 'z']}
                      fieldKey={[field.fieldKey, 'z']}
                      rules={[{ required: true, message: 'Por favor, insira o intervalo de confiança.' }]}
                    >
                      <Input type="number" />
                    </Form.Item>
                  </Col>
                  
                  <Col span={24}>
                    <Button style={{ backgroundColor: '#ff5232', color: 'white' }} onClick={() => remove(field.name)}>
                      Remover
                    </Button>
                  </Col>
                </Row>
              ))}

              <Form.Item>
                <Button type="dashed" onClick={() => add()} block icon="+">
                  Adicionar Tipo de Sensor
                </Button>
              </Form.Item>
            </>
          )}
        </Form.List>

        <Form.Item>
          <Button type="primary" htmlType="submit" style={{ width: '100%', backgroundColor: '#FFA13A !important', color: 'white' }}>
            Enviar
          </Button>
        </Form.Item>
      </Form>
    </div>
  );
};

export default SensorForm;
