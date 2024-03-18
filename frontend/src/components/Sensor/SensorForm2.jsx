import React from 'react';
import { Form, Input, Button} from "antd";
import FormItem from "antd/lib/form/FormItem";


const SensorForm2 = ({onFinish, onFinishFailed}) =>{
  return (
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
    onFinish={onFinish}
    onFinishFailed={onFinishFailed}
    autoComplete="off"
  >


<p>Insira seu tipo de sensor</p>
    <FormItem
      name="tipo"
      rules={[
        {
          required: true,
          message: 'Por favor, digite o tipo do seu sensor!',
        },
      ]}
    >
      <Input className='w-full'/>
    </FormItem>


    <p>Indique o mínimo que o sensor pode captar:</p>
    <FormItem
      name="min"
      rules={[
        {
          required: true,
          message: 'Por favor, digite o mínimo de captação do seu sensor!',
        },
      ]}
    >
      <Input className='w-full' />
    </FormItem>

    <p>Indique o máximo que o sensor pode captar:</p>
    <FormItem
      name="max"
      rules={[
        {
          required: true,
          message: 'Por favor, digite o máximo de captação do seu sensor',
        },
      ]}
    >
      <Input className='w-full'/>
    </FormItem>

<p>Indique o z:</p>
<FormItem
    name="z"
    rules={[
    {
        required: true,
        message: 'Por favor, digite seu z!',
    },
    ]}
>
    <Input className='w-full'/>
</FormItem>

  </Form>
  );
};

export default SensorForm2;