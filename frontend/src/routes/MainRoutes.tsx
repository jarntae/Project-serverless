import { lazy } from "react";
import type { RouteObject } from "react-router-dom";
import Loadable from "../third-party/Loadable";

import FullLayout from "../components/FullLayout/FullLayout.tsx";
import Dashboard from "../page/Dashboard/index.tsx";


const SignInPages = Loadable(lazy(() => import("../page/authentication/Login/index.tsx")));


const MainRoutes = (isLoggedIn : boolean): RouteObject => {
  
  return {
    path: "/",
    element: isLoggedIn ? <FullLayout /> : <SignInPages />,
    children: [
      {
        path: "/dashbord",
        element: <Dashboard />,
      },
    ],
  };
};

export default MainRoutes;