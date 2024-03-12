import React from 'react';
import Sensor from '@/components/Sensor';
import '../../app/globals.css';

export default function sensor(){
  return (
    <div className="flex h-screen justify-center items-center bg-custom-blue">
      <div className="flex flex-col gap-4">
        <Sensor></Sensor>
      </div>
    </div>
  );
};

