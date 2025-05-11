import axios from "axios";
import type { LoginInterface } from "../../interface/Login";

const apiUrl = "http://localhost:8000";
const Authorization = localStorage.getItem("token");
const Bearer = localStorage.getItem("token_type");

const requestOptions = {
  headers: {
    "Content-Type": "application/json",
    Authorization: `${Bearer} ${Authorization}`,
  },
};

async function SignIn(data: LoginInterface ) {
    return await axios
      .post(`${apiUrl}/signIn`, data, requestOptions)
      .then((res) => res)
      .catch((e) => e.response);
}

export {
    SignIn
}