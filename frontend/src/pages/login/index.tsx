import React from 'react';
import Input from '../../components/LoginInput/LoginInput';
import '../../app/globals.css';

export default function login(){
  return (
    <div className="flex h-screen justify-center items-center bg-custom-blue">
      <div className="flex flex-col gap-4">
        <Input placeholder="USUÁRIO" type="text" name="username" />
        <Input placeholder="SENHA" type="password" name="password" />
      </div>
    </div>
  );
};

