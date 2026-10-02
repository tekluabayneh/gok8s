import axios from "axios"
import https from "node:https"
import type { KubeconfType } from "../../types/configtypes.js"
import chalk from "chalk";


//TODO: 
//this function should be turned into facoty pattern
//so anyone who gets it won't call it as singleton pattern instade it will use as dependency injection so 
//everytime commands used it, it won't have side effect 

async function createFactoryApi(conf: KubeconfType, agent: https.Agent) {
  const api = axios.create({
    baseURL: conf.BASE_URL,
    httpsAgent: agent,
    timeout: 20_000,
    withCredentials: true
  })

  api.interceptors.request.use(async (config) => {
    const publicRoutes = ["/healthz", "/livez", "/readyz", "/version"]


    if (publicRoutes.some(route => config.url?.includes(route))) {
      return config;
    }

    if (conf.token) {
      config.headers.Authorization = `Bearer ${conf.token}`;
    }

    return config;
  }, (error) => {
    if (axios.isAxiosError(error)) {
      console.log(error.response?.data.message)
    } else {
      console.log(chalk.red("something went wrong", error))
    }
    return Promise.reject(error);
  });


  return api

}

export default createFactoryApi

