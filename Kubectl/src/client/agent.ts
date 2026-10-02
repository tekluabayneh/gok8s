
import https from "node:https"
import type { KubeconfType } from "../../types/configtypes.js"


// this function should user config as parametere 
async function buildAgent(config: KubeconfType) {
  const agentOpts: https.AgentOptions = {}

  if (config.insecureSkipTlsVerify) {
    agentOpts.rejectUnauthorized = false
  } else if (config.certificateAuthorityData) {
    agentOpts.ca = Buffer.from(config.certificateAuthorityData, "base64")
  }

  if (config.clientCertificateData && config.clientKeyData) {
    agentOpts.cert = Buffer.from(config.clientCertificateData, "base64")
    agentOpts.key = Buffer.from(config.clientKeyData, "base64")
  }

  return new https.Agent(agentOpts)
}

export default buildAgent
