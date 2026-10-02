import chalk from "chalk"
import buildAgent from "./agent.js"
import createFactoryApi from "./client.js"
import loadConfig from "./config.js"

const creaetApi = async () => {
  const resOfConfig = await loadConfig()
  if (!resOfConfig) {
    console.log(chalk.red("configuration files is not loaded"))
    return
  }

  return createFactoryApi(resOfConfig, await buildAgent(resOfConfig))
}

export default creaetApi
