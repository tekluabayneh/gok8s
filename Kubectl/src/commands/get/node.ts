import { Args, Command } from '@oclif/core'
import createFlags from '../../flags/statis.js'
import { nodeService } from '../../services/node.js'
import creaetApi from '../../client/create-api.js'
import { table } from 'table'


export default class Nodes extends Command {
  static aliases = ['get:nodes']
  static args = {
    podName: Args.string(),
  }
  static flags = createFlags({})
  static strict: boolean = false

  async run(): Promise<void> {
    const { argv } = await this.parse(Nodes)
    const api = await creaetApi()


    if (argv.length > 0) {
      await Promise.all(argv.map(async (name) => {
        const url = `/api/v1/nodes/${name}`
        const tableData = await nodeService(api, url, null, "get")
        if (!tableData) return
        this.log(table(tableData))
      }))
      return
    }

    const tableData = await nodeService(api, "/api/v1/nodes", null, "get")
    if (!tableData) return
    this.log(table(tableData))


  }

}


