import chalk from 'chalk'
import creaetApi from '../../client/create-api.js'
import createFlags from '../../flags/statis.js'
import { podService } from '../../services/pod.js'
import { Args, Command } from '@oclif/core'
import { table } from 'table'


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


    await Promise.all(argv.map(async (name) => {
      const url = `/api/v1/namespaces/${namespace ?? "default"}/pods/${name}`
      const tableData = await podService(api, url, namespace, null, "delete", name as string)
      if (!tableData) return
      this.log(table(tableData))

    }))
  }

}


