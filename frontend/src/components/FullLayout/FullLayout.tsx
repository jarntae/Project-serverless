import React, { useEffect, useState } from "react";
import { Routes, Route, useLocation } from "react-router-dom";
import "../../App.css";
import { Breadcrumb, Layout } from "antd";

import SignInPages from "../../page/authentication/Login/index.tsx";

const { Content } = Layout;

const FullLayout: React.FC = () => {
  const location = useLocation();
  const [checkLogin, setCheckLogin] = useState(false);

  useEffect(() => {
    setCheckLogin(location.pathname !== "/");
  }, [location.pathname]);


  return (
    <>
      <Layout style={{ minHeight: "100vh", backgroundColor: "#F5F5F5" }}>
        {checkLogin}
        <Layout style={{ backgroundColor: "#f7f8fc", minHeight: "100vh" }}>
          <Content style={{ marginTop: "0px" }}>
            <Breadcrumb style={{ marginTop: "0px" }} />
            <div>
              <Routes>
                <Route path="/" element={<SignInPages />} />
                {/* <Route path="/dashboard" element={<Dashboard />} /> */}
                
              </Routes>
            </div>
          </Content>
        </Layout>
      </Layout>
    </>
  );
};

export default FullLayout;