
import axios, { AxiosInstance, AxiosResponse } from "axios"
import renderToTerminal, { ResourceMap } from "../utils/render-to-terminal.js"
import chalk from "chalk"
type ActionType = "create" | "delete" | "get"


// this function must flexable to handle all kinds pods service
export const NamespaceService = async (api: AxiosInstance | undefined, url: string, manifest: JSON | null | object, type: ActionType, args?: { nameOfNs: string | undefined }, filename?: string) => {
  try {
    if (!api) {
      console.log(chalk.red("something went wrong Opps!"))
      return
    }

    let res: AxiosResponse

    switch (type) {
      case "create": {
        res = await api.post(`/api/v1/namespaces`, manifest)
        if (res.data.status.phase === "Active") {
          console.log(chalk.green(`namespace/${args?.nameOfNs ?? filename} creared`))
        }

        return
      }

      case "delete": {
        res = await api.delete(url)
        if (res.status == 200) {
          console.log(chalk.green(`namespace "${args?.nameOfNs ?? filename}" deleted`))
        }

        return
      }

      case "get": {
        res = await api.get(url)
        renderToTerminal(res.data.items ?? [res.data], "Namespace")
        if (res?.data?.item ? res?.data?.items?.length === 0 : false) {
          console.log(`no resource are found `)
        }

        return
      }

      default: {
        console.log("default rached i dont why")
      }

    }

  } catch (error) {
    if (axios.isAxiosError(error)) {
      console.log(chalk.red(error.response?.data.message))
    } else {
      console.log(chalk.red("something went wrong", error))
    }
  }

}
