import { Args, Command } from '@oclif/core'
import chalk from 'chalk'
import creaetApi from '../../client/create_api.js'
import createFlags from '../../flags/statis.js'
import { PodService } from '../../services/pod.js'


export default class DelPods extends Command {
  static aliases: string[] = ["delete:pods"]
  static args = {
    podName: Args.string(),
  }
  static flags = createFlags({})
  static strict = false


  async run(): Promise<void> {
    const { argv, flags } = await this.parse(DelPods)
    const { namespace } = flags

    const api = await creaetApi()

    if (argv.length === 0) {
      console.log(chalk.red("at list one resouce name is requied"))
      return
    }


    for (const name of argv) {
      const url = `/api/v1/namespaces/${namespace ?? "default"}/pods/${name}`
      PodService(api, url, namespace, null, "delete", name as string)


    }

  }

}


