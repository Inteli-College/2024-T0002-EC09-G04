import React from "react";
import Image from "next/image";
import styles from "./page.module.css";
import  RootLayout  from './layout';
import HomePage from '../views/homepage/HomePage';
import LoginPage from  '../views/loginpage/LoginPage';

function Home() {
  return (
    <RootLayout>
      <LoginPage/>
    </RootLayout>
  );
}

export default Home;
