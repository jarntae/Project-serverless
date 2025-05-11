import { useRoutes, type RouteObject } from "react-router-dom";
import MainRoutes from "../routes/MainRoutes";
import LoginRoutes from "../routes/LoginRoutes";


function ConfigRoutes() {
  const isLoggedIn = localStorage.getItem("isLogin") === "true";

  let routes: RouteObject[] = [];

  if (isLoggedIn) {
    routes = [MainRoutes(isLoggedIn), LoginRoutes()];
  } else {
    routes = [LoginRoutes()];
  }

  return useRoutes(routes);
}

export default ConfigRoutes;
