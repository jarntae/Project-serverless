import { lazy } from "react";
import { RouteObject } from "react-router-dom";
import Loadable from "../third-party/Loadable";

import FullLayout from "../components/FullLayout/FullLayout.tsx";


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