import axios, { AxiosInstance, AxiosResponse } from "axios"
import renderToTerminal from "../utils/render-to-terminal.js"
import chalk from "chalk"
type ActionType = "create" | "delete" | "get"

//TODO: 
//first proper message need to be loged 
//second reptitive logs need to be remove and one certralized logs need to handle them 


export const podService = async (api: AxiosInstance | undefined, url: string, namespace: string | undefined, manifest: JSON | null, type: ActionType, name?: string) => {
  try {
    if (!api) {
      console.log(chalk.red("something went wrong Opps!"))
      return
    }

    let res: AxiosResponse

    switch (type) {
      case "create": {
        res = await api.post(url, manifest)
        if (res.status == 201) {
          //@ts-expect-error metadata.name always exists till i find better way to get the name this stays like this 
          console.log(chalk.green(`pod/"${manifest?.metadata?.name ?? ""}" created`))
        }
        return
      }

      case "delete": {
        res = await api.delete(url)
        console.log(res)
        if (res.status == 200) {
          console.log(chalk.green(`pod "${name}" deleted from ${namespace} namespace`))
        } else {
          console.log(chalk.red(`${res.data.message}`))
        }

        return
      }

      case "get": {
        res = await api.get(url)
        if (res?.data?.item ? res?.data?.items?.length === 0 : false) {
          console.log(`no resource are found in the ${namespace ?? "default"} namespace`)
        }
        return renderToTerminal(res.data.items ?? [res.data], "Pod")
      }

      default: {
        console.log("default")
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
