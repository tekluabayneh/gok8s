import { Args, Command } from '@oclif/core'
import creaetApi from '../../client/create-api.js'
import chalk from 'chalk'
import loadConfig from '../../client/config.js'
import { podService } from '../../services/pod.js'
import createFlags from '../../flags/statis.js'
import { table } from 'table'


export default class Pods extends Command {
  static aliases = ['get:pod']
  static args = {
    podName: Args.string(),
  }
  static flags = createFlags({})


  // TODO: 
  // this function must also handle single post request with muttiple names 

  async run(): Promise<void> {
    const { args, flags } = await this.parse(Pods)
    const { namespace } = flags

    // TODO: 
    // change the iterating with Promise.ALL instade of just using this one which also raise lint error
    const resOfConfig = await loadConfig()
    if (!resOfConfig) {
      console.log(chalk.red("internal problem loading context configuration, Opps!"))
      return
    }

    const namePart = args.podName ? "/" + args.podName : ""
    const url = `/api/v1/namespaces/${namespace ?? "default"}/pods${namePart}`
    const api = await creaetApi()
    const tableData = await podService(api, url, namespace, null, "get")
    if (!tableData) return
    this.log(table(tableData))

  }
}


