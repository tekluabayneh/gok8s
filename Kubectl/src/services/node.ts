
import axios, { AxiosInstance, AxiosResponse } from "axios"
import renderToTerminal, { ResourceMap } from "../utils/render-to-terminal.js"
import chalk from "chalk"
type ActionType = "create" | "delete" | "get"


// this function must flexable to handle all kinds pods service
export const nodeService = async (api: AxiosInstance | undefined, url: string, manifest: JSON | null, type: ActionType) => {
  try {
    if (!api) {
      console.log(chalk.red("something went wrong Opps!"))
      return
    }

    let res: AxiosResponse

    switch (type) {
      case "get": {
        res = await api.get(url)
        renderToTerminal(res.data.items ?? [res.data], "Node")
        if (res?.data?.item ? res?.data?.items?.length === 0 : false) {
          console.log(`no resource are found`)
        }
      }
        
      case "create":
    }

  } catch (error) {
    if (axios.isAxiosError(error)) {
      console.log(chalk.red(error.response?.data.message))
    } else {
      console.log(chalk.red("something went wrong", error))
    }
  }

}
