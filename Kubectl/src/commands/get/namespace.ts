import { Args, Command } from '@oclif/core'
import creaetApi from '../../client/create_api.js'
import createFlags from '../../flags/statis.js'
import { NamespaceService } from '../../services/namespace.js'

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
      for (const name of argv) {
        const url = `/api/v1/namespaces/${name}`
        NamespaceService(api, url, null, "get")
      }

      return
    }

    const url = "/api/v1/namespaces/"
    NamespaceService(api, url, null, "get")
  }
}


