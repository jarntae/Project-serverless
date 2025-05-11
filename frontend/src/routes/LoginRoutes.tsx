import { lazy } from "react";
import type { RouteObject } from "react-router-dom";
import MinimalLayout from "../components/MinimalLayout/MinimalLayout";
import Loadable from "../third-party/Loadable";

const LoginPages = Loadable(lazy(() => import("../page/authentication/Login/index")))

const LoginRoutes = (): RouteObject => {
  return {
    path: "/",
    element: <MinimalLayout />,
    children: [
      {
        path: "/",
        element: <LoginPages/>,
      },
      {
        path: "/login",
        element: <LoginPages />,
      },
    ],
  };
};

export default LoginRoutes;