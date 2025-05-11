import React from "react";
import { Routes, Route } from "react-router-dom";
import "../../App.css";
import { Breadcrumb, Layout, theme, message } from "antd";

// import Bar from "../Header/index"
// import BarMember from "../Header/member.tsx"
// import BarEmployee from "../Header/employee.tsx"

import Dashboard from "../../page/Dashboard/";
import LoginPages from "../../page/authentication/Login";

// import Room from "../Pages/Room/index.tsx"
// import FacilityManager from "../Pages/Facility/Manager/index.tsx";
// import FacilityBooking from "../Pages/Facility/Member/index.tsx";

// import Employee from "../Pages/Employee/employee";
// import EmployeeCreate from "../Pages/Employee/create/createEmployee";
// import EmployeeEdit from "../Pages/Employee/edit/editEmployee";
// import ProfileEdit from "../Pages/ProfileEdit/profileEdit";
// import ChangePassword from "../Pages/ProfileEdit/changePassword";

// import Member from "../Pages/Member/member"
// import MemberCreate from "../Pages/Member/create/createMember"
// import MemberEdit from "../Pages/Member/edit/editMember"


const { Content } = Layout;

const FullLayout: React.FC = () => {
    const [_messageApi, contextHolder] = message.useMessage();
    //ใช้ underscore ( _ ) แก้ปัญหา Declared but not use
    
    // const renderSider = () => {
    //   if (role === "Manager") {
    //     return <Bar />;
    //   } else if (role === "Member") {
    //     return <BarMember />;
    //   } else {
    //     return <BarEmployee />;
    //   }
    // };
  
    const {
      token: { colorBgContainer: _colorBgContainer },
    } = theme.useToken();
  
    return (
      <>
        {contextHolder}
        <Layout style={{ minHeight: "100vh", backgroundColor: "#F5F5F5" }}>
          {/* {renderSider()} */}
  
          <Layout style={{ backgroundColor: "#434343", minHeight: "100vh" }}>
            <Content style={{ margin: "0 30px" }}>
              <Breadcrumb style={{ margin: "16px 0" }} />
              <div
                style={{
                  padding: 24,
                  minHeight: "93%",
                  maxHeight: "93%",
                  background: "#f5f5f5",
                  borderRadius: "8px",
                  boxShadow: "0 4px 12px rgba(0, 0, 0, 0.1)", 
                }}
              >
                <Routes>
                    <Route path="/" element={<LoginPages/>}/>
                    <Route path="/dashboard" element={<Dashboard />} />
                </Routes>
              </div>
            </Content>
          </Layout>
        </Layout>
      </>
    );
  };
  
  export default FullLayout;
  