import os from "node:os"

import type { KubectlconfigType } from "../../types/configtypes.d.js"

import type { KubeconfType } from "../../types/configtypes.js"
import { yamlToJson } from "../utils/yaml-to-json.js"

import chalk from "chalk"
import { promises as stPromise } from "node:fs"


// loadconfig shoudl return those credential not load them to global varible 
async function loadConfig(): Promise<KubeconfType | void> {
  let conf = {
    "currentContext": "",
    "certificateAuthorityData": "",
    "insecureSkipTlsVerify": false,
    "contextUser": "",
    "contextCluster": "",
    "token": "",
    "BASE_URL": "",
    "clientCertificateData": "",
    "clientKeyData": ""
  }

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


  conf.currentContext = jsonFile["current-context"]

  for (const ctx of jsonFile.contexts) {
    if (ctx.name === conf.currentContext) {
      conf.contextUser = ctx.context.user
      conf.contextCluster = ctx.context.cluster
    }
  }


  for (const key in jsonFile) {
    if (!Object.hasOwn(jsonFile, key)) {
      continue
    }


    // extract out token, client-key-data, certificateAuthorityData from usres with a given current-context data
    if (key === "users") {
      for (const usr of jsonFile[key]) {
        if (usr.name === conf.contextUser) {
          conf.token = usr?.user?.token ?? ""
          conf.clientCertificateData = usr.user["client-certificate-data"] ?? ""
          conf.clientKeyData = usr.user["client-key-data"] ?? ""

        }
      }
    }

    // find name that mach the contextCluster and update the certificateAuthorityData and insecure-skip-tls-verify
    if (key === "clusters") {
      for (const clus of jsonFile[key]) {
        if (clus.name === conf.contextCluster) {
          conf.BASE_URL = clus.cluster.server
          conf.certificateAuthorityData = clus.cluster["certificate-authority-data"] ?? ""
          conf.insecureSkipTlsVerify = clus.cluster["insecure-skip-tls-verify"] ?? false
        }
      }
    }
  }
  return conf
}

export default loadConfig
