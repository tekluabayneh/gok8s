import axios from "axios"
import fs from "fs"
import os from "os"
import { yamlToJson } from "../utils/yamlToJson.js"
import chalk from "chalk"
import { parse } from "yaml"

//FIRE: 
//one must thing i have to taken care of is all resouce does not need same 
//RBAC so things change and that also need to be taken care of
//HOT: 
// 
//WARM: 
//
//COOL: 
//
//ICE: 
//
//DONE: reading and extracting yaml file to json already is done so i could use that with some modificatin to the function 
let BASE_URL = ""


// get base url form mykbectl config 
// TODO 
// read conf file from ~/.kube/config this files will provide all the config that kubectl needs 
const api = axios.create({ baseURL: BASE_URL, timeout: 20000, withCredentials: true })

//TODO 
//
//read file from ~/.kube/config DONE:

//make sure the directory and the files exist  DONE:
//and make sure file is not empty  DONE:

//make sure files has properate fiels that are need for apiserver 

//make sure any error are handled well 
// const RootPath = fs.readdir(, options, callback)


let token = ""
const defaultPath = os.homedir() + "/.kube/config"
fs.stat(defaultPath, (error, stats) => {
  if (error) {
    console.log(chalk.red(error))
    return
  }

  // change to json and extract out what needs to be send through the api  
  console.log(stats.isFile())
  if (stats.size > 0) {
    const yaml = fs.readFileSync(defaultPath, "utf-8")
    const jsonFile = parse(yaml)

    console.log("file form .kube", jsonFile)
  }

})


api.interceptors.request.use(async (config) => {
  const publicRoutes = ["/fake one herer"];


  if (publicRoutes.some(route => config.url?.includes(route))) {
    return config;
  }


  // we should be able to get token form config ~/.kube/config  file 

  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }


  return config;
}, (error) => {
  return Promise.reject(error);
});

export default api
