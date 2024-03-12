import React from 'react';
import Image from 'next/image';
import {Input}  from 'antd'
import './../../app/globals.css';

import logo from '../../assets/img/id_visual.png';

export default function Login() {
  return (
    <div className="flex h-screen justify-center items-center bg-custom-blue">
      <div className="flex flex-col gap-4">
        <Image
          src={logo}
          width={100}
          height={100}
          alt="Logo da empresa"
        />
        <Input placeholder="USUÁRIO" type="text" name="username" />
        <Input placeholder="SENHA" type="password" name="password" />
      </div>
    </div>
  );
}
