import axios from "axios"
import os from "node:os"

import type KubectlconfigType from "../../types/configtypes.d.js"
import { yamlToJson } from "../utils/yaml-to-json.js"

import chalk from "chalk"
import { promises as stPromise } from "node:fs"
import https from "node:https"


let currentContext = ""
let certificateAuthorityData = ""
let insecureSkipTlsVerify = false
let clientCertificateData = ""
let clientKeyData = ""
let BASE_URL = ""
let token = ""
let contextUser = ""
let contextCluster = ""

async function loadConfig() {
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

  for (const ctx of jsonFile.contexts) {
    if (ctx.name === currentContext) {
      contextUser = ctx.context.user
      contextCluster = ctx.context.cluster
    }
  }


  for (const key in jsonFile) {
    if (!Object.hasOwn(jsonFile, key)) {
      continue
    }


    // extract out token, client-key-data, certificateAuthorityData from usres with a given current-context data
    if (key === "users") {
      for (const usr of jsonFile[key]) {
        if (usr.name === contextUser) {
          token = usr?.user?.token ?? ""
          clientCertificateData = usr.user["client-certificate-data"] ?? ""
          clientKeyData = usr.user["client-key-data"] ?? ""
        }
      }
    }

    // find name that mach the contextCluster and update the certificateAuthorityData and insecure-skip-tls-verify
    if (key === "clusters") {
      for (const clus of jsonFile[key]) {
        if (clus.name === contextCluster) {
          BASE_URL = clus.cluster.server
          certificateAuthorityData = clus.cluster["certificate-authority-data"] ?? ""
          insecureSkipTlsVerify = clus.cluster["insecure-skip-tls-verify"] ?? false
        }
      }
    }
  }

}

await loadConfig()
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
  httpsAgent: await buildAgent(),
  timeout: 20_000,
  withCredentials: true
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

