import React from 'react';
interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {}

const Sensor: React.FC<InputProps> = ({ placeholder, ...props }) => {


  return (
    <div className="relative w-full">
      <h1>Você é muito feioso</h1>
    </div>
  );
};

export default Sensor;