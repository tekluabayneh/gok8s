import { Args, Command } from '@oclif/core'
import createFlags from '../../flags/statis.js'
import { nodeService } from '../../services/node.js'
import creaetApi from '../../client/create_api.js'


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
      for (const element of argv) {
        const url = `/api/v1/nodes/${element}`
        nodeService(api, url, null, "get")
      }

      return
    }

    nodeService(api, "/api/v1/nodes", null, "get")
  }

}


