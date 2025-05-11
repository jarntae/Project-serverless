import { useRoutes, RouteObject } from "react-router-dom";

import MainRoutes from "./MainRoutes";

function ConfigRoutes() {

  const isLoggedIn = localStorage.getItem("isLogin") === "true";

  let routes: RouteObject[] = [];


  if (isLoggedIn) {

    routes = [MainRoutes(isLoggedIn)];

  } else {

    routes = [MainRoutes()];

  }


  return useRoutes(routes);

}


export default ConfigRoutes;