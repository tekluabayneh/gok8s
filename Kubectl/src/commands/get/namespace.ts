import { Args, Command } from '@oclif/core'
import creaetApi from '../../client/create-api.js'
import createFlags from '../../flags/statis.js'
import { nameSpaceService } from '../../services/namespace.js'
import { table } from 'table'

export default class Namespace extends Command {
  static aliases = ['get:ns']
  static args = {
    nsName: Args.string(),
  }
  static flags = createFlags({})
  static strict = false

  async run(): Promise<void> {
    const { argv } = await this.parse(Namespace)
    const api = await creaetApi()

    if (argv.length > 0) {
      await Promise.all(argv.map(async (name) => {
        const url = `/api/v1/namespaces/${name}`
        const tableData = await nameSpaceService(api, url, null, "get")
        if (!tableData) return
        this.log(table(tableData))

      }))

      return
    }

    const url = "/api/v1/namespaces/"
    const tableData = await nameSpaceService(api, url, null, "get")

    if (!tableData) return
    this.log(table(tableData))




  }
}


