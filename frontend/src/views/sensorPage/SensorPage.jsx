"use client"; // This is a client component 👈🏽
import React, { useState } from 'react';
import { Col, Row, Button} from "antd";
import SensorForm from '@/components/Sensor/SensorForm';
import SensorText from '@/components/Sensor/SensorText';
import FormItem from "antd/lib/form/FormItem";

const SensorPage = ({ onFinish, onFinishFailed }) => {

    return (
        <div className="bg-white flex flex-col justify-center items-center min-h-screen">
            <Row justify='center' style={{ width: '100%' }}>
                <Col span={10}>
                    <SensorText/>
                </Col>
                <Col span={12}>
                    <h1 className='font-bold text-2xl font-mono	my-4 text-custom-purple'>Adicione Novos Sensores</h1>
                    <SensorForm onFinish={onFinish} onFinishFailed={onFinishFailed}/>
                </Col>
            </Row>
        
        </div>
    );
};

export default SensorPage;
