import axios from "axios"
import os from "os"
import { yamlToJson } from "../utils/yamlToJson.js"
import chalk from "chalk"
import KubectlconfigType from "../../types/configtypes.d.js"
import { promises as stPromise } from "fs"
import https from "https"

let currentContext = ""
let certificateAuthorityData = ""
let insecureSkipTlsVerify = false
let clientCertificateData = ""
let clientKeyData = ""
let BASE_URL = ""
let token = ""
let contextUser = ""
let contextCluster = ""

async function LoadConfig() {
  const defaultPath = os.homedir() + "/.kube/config"
  let stats
  try {
    stats = await stPromise.stat(defaultPath)
  } catch (error) {
    console.log(chalk.red(error))
  }

  if (stats && stats?.size < 1) {
    console.log(chalk.red("file is empty"))
    return
  }

  const jsonFile = await yamlToJson<KubectlconfigType>(defaultPath)
  if (!jsonFile) return


  currentContext = jsonFile["current-context"]

  for (const ctx of jsonFile["contexts"]) {
    if (ctx.name == currentContext) {
      contextUser = ctx.context.user
      contextCluster = ctx.context.cluster
    }
  }

  for (const key in jsonFile) {
    if (key == "users") {
      for (const usr of jsonFile[key]) {
        if (usr.name == contextUser) {
          token = usr?.user?.token ?? ""
          clientCertificateData = usr.user["client-certificate-data"] ?? ""
          clientKeyData = usr.user["client-key-data"] ?? ""
        }
      }
    }

    if (key == "clusters") {
      for (const clus of jsonFile[key]) {
        if (clus.name == contextCluster) {
          BASE_URL = clus.cluster.server
          certificateAuthorityData = clus.cluster["certificate-authority-data"] ?? ""
          insecureSkipTlsVerify = clus.cluster["insecure-skip-tls-verify"] ?? false
        }
      }
    }
  }
}

await LoadConfig()
async function buildAgent() {
  const agentOpts: https.AgentOptions = {}

  if (insecureSkipTlsVerify) {
    agentOpts.rejectUnauthorized = false
  } else if (certificateAuthorityData) {
    agentOpts.ca = Buffer.from(certificateAuthorityData, "base64")
  }

  if (clientCertificateData && clientKeyData) {
    agentOpts.cert = Buffer.from(clientCertificateData, "base64")
    agentOpts.key = Buffer.from(clientKeyData, "base64")
  }

  return new https.Agent(agentOpts)
}

const api = axios.create({
  baseURL: BASE_URL,
  timeout: 20000,
  withCredentials: true,
  httpsAgent: await buildAgent()
})


api.interceptors.request.use(async (config) => {
  const publicRoutes = ["/healthz", "/livez", "/readyz", "/version"]


  if (publicRoutes.some(route => config.url?.includes(route))) {
    return config;
  }

  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }

  return config;
}, (error) => {
  return Promise.reject(error);
});

export default api

