// routes/LoginRoutes.tsx
import { lazy } from "react";
import { useRoutes } from "react-router-dom";
import MinimalLayout from "../components/MinimalLayout/MinimalLayout";
import Loadable from "../third-party/Loadable";

const LoginPages = Loadable(lazy(() => import("../page/authentication/Login/index")));

const ConfigRoutes = () => {
  const routes = useRoutes([
    {
      path: "/",
      element: <MinimalLayout />,
      children: [
        { path: "/", element: <LoginPages /> },
        { path: "/login", element: <LoginPages /> },
      ],
    },
  ]);

  return routes;
};

export default ConfigRoutes;
