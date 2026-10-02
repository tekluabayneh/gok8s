import axios, { AxiosResponse, AxiosInstance } from "axios"
import renderToTerminal, { ResourceMap } from "../utils/render-to-terminal.js"
import chalk from "chalk"
type ActionType = "create" | "get" | "delete"


// this function must flexable to handle all kinds pods service
export const PodService = async (api: AxiosInstance | undefined, url: string, namespace: string | undefined, manifest: JSON | null, type: ActionType, name?: string) => {
  try {
    if (!api) {
      console.log(chalk.red("something went wrong Opps!"))
      return
    }
    let res: AxiosResponse

    switch (type) {
      case "get":
        res = await api.get(url)
        renderToTerminal(res.data.items ?? [res.data], "Pod")
        if (res?.data?.item ? res?.data?.items?.length === 0 : false) {
          console.log(`no resource are found in the ${namespace ?? "default"} namespace`)
        }
        return
      case "create":
        res = await api.post(url, manifest)
        console.log("res", res.data)
        return
      case "delete":
        res = await api.delete(url)
        console.log(res)
        if (res.status == 200) {
          console.log(chalk.green(`pod "${name}" deleted from ${namespace} namespace`))
        } else {
          console.log(chalk.red(`${res.data.message}`))
        }
        return
      default:
        console.log("default")
    }

  } catch (error) {
    if (axios.isAxiosError(error)) {
      console.log(chalk.red(error.response?.data.message))
    } else {
      console.log(chalk.red("something went wrong", error))
    }
  }

}
