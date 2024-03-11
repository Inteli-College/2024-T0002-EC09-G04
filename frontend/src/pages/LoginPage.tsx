import React from 'react';
import Input from '../components/LoginInput';

const LoginPage: React.FC = () => {
  return (
    <div className="flex h-screen justify-center items-center bg-custom-blue">
      <div className="flex flex-col gap-4">
        <Input placeholder="USUÁRIO" type="text" name="username" />
        <Input placeholder="SENHA" type="password" name="password" />
      </div>
    </div>
  );
};

export default LoginPage;
