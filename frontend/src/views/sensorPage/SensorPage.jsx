"use client"; // This is a client component 👈🏽
import React, { useState } from 'react';
import { Col, Row, Button} from "antd";
import SensorForm from '@/components/Sensor/SensorForm';
import SensorText from '@/components/Sensor/SensorText';
import FormItem from "antd/lib/form/FormItem";
import MapComponent from '@/components/Map/Map';

const SensorPage = ({ onFinish, onFinishFailed }) => {

        const [location, setLocation] = useState({ lat: -23.55052, lng: -46.633308 });
    
        const handleLocationSelect = (latlng) => {
            setLocation(latlng);
            console.log('Localização selecionada:', latlng);
        };

    return (
        <div className="bg-white flex flex-col justify-center items-center min-h-screen">
            <Row justify='center' style={{ width: '100%' }}>
                <Col span={10}>
                    <SensorText/>
                </Col>
                <Col span={12}>
                    <h1 className='font-bold text-2xl font-mono	my-4 text-custom-purple'>Adicione Novos Sensores</h1>
                    <MapComponent onLocationSelect={handleLocationSelect} />
                    <p>Latitude: {location.lat}, Longitude: {location.lng}</p>
                    <SensorForm onFinish={onFinish} onFinishFailed={onFinishFailed}/>
                </Col>
            </Row>
        
        </div>
    );
};

export default SensorPage;
