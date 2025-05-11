import { Form, Input, message } from "antd";
import { SignInAdmin } from "../../../services/https";
import './index.css'


function SignInPages() {

    const [messageApi, contextHolder] = message.useMessage();

    const onFinish = async (values: { email: string; password: string }) => {
        const { email, password } = values;

        try {
            const adminValues: SignInAdminInterface = { email, password };
            const res = await SignInAdmin(adminValues);

            if (res.status === 200) {
                messageApi.success("Admin sign-in successful");
                localStorage.setItem("isLogin", "true");
                localStorage.setItem("page", "admin-dashboard");
                localStorage.setItem("token_type", res.data.token_type);
                localStorage.setItem("token", res.data.token);
                localStorage.setItem("profile", res.data.profile);
                localStorage.setItem("id", res.data.id);
                localStorage.setItem("role", "admin");

                setTimeout(() => {
                    location.href = "/dashboard";
                }, 2000);
            } else {
                messageApi.error(res.data.error);
            }
        } catch {
            messageApi.error("Sign-in failed. Please try again.");
        }
    };

    return (
        <>
            {contextHolder}
            <div className="container">
                <div className="left">
                    <div className="header">
                        <h2 className="animation a1">Welcome To NEXT SUT</h2>
                        <h4 className="animation a2">Log in to your account using email and password</h4>
                    </div>
                    <div className="form">
                        <Form
                            name="basic"
                            onFinish={onFinish}
                            autoComplete="off"
                            layout="vertical"
                        >
                            <Form.Item
                                label="Email"
                                name="email"
                                rules={[
                                    { required: true, message: "Please input your email!" },
                                    { type: "email", message: "Please enter a valid email!" },
                                ]}
                            >
                                <Input />
                            </Form.Item>
                            <Form.Item
                                label="Password"
                                name="password"
                                rules={[
                                    { required: true, message: "Please input your password!" },
                                ]}
                            >
                                <Input.Password />
                            </Form.Item>

                            <Form.Item>

                                <button className="button-login" type="submit" >
                                    <span>LOG IN</span>
                                </button>
                            </Form.Item>
                        </Form>
                    </div>
                </div>
                <div className="right"></div>
            </div>
        </>
    );
}

export default SignInPages;