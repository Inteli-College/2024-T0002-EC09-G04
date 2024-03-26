
import React from "react";
import Link from "next/link";
import Image from "next/image";
import { Button, Layout, Menu, Row, Col, Space } from "antd";
import MenuItem from "antd/lib/menu/MenuItem";

const Header = Layout;
const Content = Layout;
const Footer = Layout;

import logo from "../../assets/img/id_visual.png";
import header from "../../assets/img/header.png";

const HomePage = () => {
  return (
    <>
      <Row
        justify="end"
        style={{ backgroundColor: "#2E329B", padding: 12 }}
        align="middle"
      >
        <Col span={4}>
          <Image src={logo} width={75} />
        </Col>
        <Col span={19} style={{ display: "flex", justifyContent: "end" }}>
          <Button
            type="link"
            style={{
              fontSize: 16,
              backgroundColor: "#FFA13A",
              color: "#fff",
              fontWeight: 700,
              fontFamily: "'Space Mono', monospace",
              fontSize: 18,
            }}
            href="/login"
          >
            CADASTRE-SE
          </Button>
        </Col>
      </Row>
      <Layout className="layout">
        <Row
          style={{
            padding: "0 50px",
            padding: 36,
            height: "81vh",
            backgroundColor: "#fff",
          }}
          align="middle"
        >
          <Col span={12} style={{ padding: 100, alignItems: "center" }}>
            <Space direction="vertical" size={20}>
              <Row>
                <h1
                  style={{
                    fontFamily: "'Space Mono', monospace",
                    fontSize: 64,
                    fontWeight: 700,
                  }}
                >
                  Orbit City
                </h1>
              </Row>
              <Row>
                <p
                  style={{ fontFamily: "Montserrat, sans-serif", fontSize: 18 }}
                >
                  Um software que simula o monitoramento de fatores climáticos
                  através de dispositivos IoT, por meio do qual seria possível o
                  acesso a um dashboard contendo informações relevantes acerca
                  dos indicadores monitorados pela Prodam.
                </p>
              </Row>
              <Row>
                <Button
                  style={{
                    fontSize: 16,
                    backgroundColor: "#FFA13A",
                    color: "#fff",
                    fontWeight: 700,
                    fontFamily: "'Space Mono', monospace",
                    fontSize: 18,
                  }}
                  type='link'
                  href="/login"
                >
                  CADASTRE-SE
                </Button>
              </Row>
            </Space>
          </Col>
          <Col span={10}>
            <Image src={header} />
          </Col>
        </Row>
      </Layout>
    </>
  );
};

export default HomePage;
