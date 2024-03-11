import React from 'react';
interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {}

const Input: React.FC<InputProps> = ({ placeholder, ...props }) => {
  // Determina a URL da imagem com base no texto do placeholder
  let imageUrl = '';
  if (placeholder === 'USUÁRIO') {
    imageUrl = '/user.png';
  } else {
    imageUrl = '/lock.png';
  }

  return (
    <div className="relative w-full">
      {imageUrl && (
        <img
          src={imageUrl}
          className="absolute left-0 bottom-0 mb-4 ml-4 h-8 w-8"
          alt=""
        />
      )}
      <input
        placeholder={placeholder}
        className="bg-custom-blue text-white border-4 border-white rounded-lg p-4 pl-16 w-full outline-none focus:border-black" 
        {...props}
      />
    </div>
  );
};

export default Input;
