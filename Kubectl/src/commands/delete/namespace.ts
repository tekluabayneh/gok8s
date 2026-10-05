import { Args, Command } from '@oclif/core'
import creaetApi from '../../client/create_api.js'
import createFlags from '../../flags/statis.js'
import { NamespaceService } from '../../services/namespace.js'
import chalk from 'chalk'
import { table } from 'table'

export default class DelNamespace extends Command {
  static aliases = ['delete:ns']
  static args = {
    podName: Args.string(),
  }
  static flags = createFlags({})
  static strict = false


  async run(): Promise<void> {
    const { argv } = await this.parse(DelNamespace)

    const api = await creaetApi()
    if (argv.length === 0) {
      console.log(chalk.red("at list one resouce name is requied"))
      return
    }

    for (const name of argv) {
      const url = `/api/v1/namespaces/${name}`
      const tableData = await NamespaceService(api, url, null, "delete", undefined, name as string)
      if (!tableData) return
      this.log(table(tableData))



    }
  }
}


