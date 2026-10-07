import { Args, Command } from '@oclif/core'
import { yamlToJson } from '../../utils/yaml-to-json.js'
import creaetApi from '../../client/create-api.js'
import { nameSpaceService } from '../../services/namespace.js'
import createFlags from '../../flags/statis.js'
import { table } from 'table'


export default class Namespace extends Command {
  static aliases = ['create:ns']
  static args = {
    nameOfNs: Args.string()
  }
  static flags = createFlags({})
  static strict = false

  async run(): Promise<void> {
    const { args, flags } = await this.parse(Namespace)
    const { filename } = flags
    const api = await creaetApi()

    if (filename) {
      const RootPath = process.cwd() + filename
      const jsonfile = await yamlToJson(RootPath)
      const tableData = await nameSpaceService(api, `/api/v1/namespaces`, jsonfile as JSON, "create", args, filename)
      if (!tableData) return
      this.log(table(tableData))


      return

    }

    const jsonfile = { metadata: { name: args.nameOfNs } }
    const tableData = await nameSpaceService(api, `/api/v1/namespaces`, jsonfile as object, "create", args)
    if (!tableData) return
    this.log(table(tableData))


  }

}






