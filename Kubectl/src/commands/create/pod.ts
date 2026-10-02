import { Command } from '@oclif/core'
import { yamlToJson } from '../../utils/yaml-to-json.js'
import chalk from 'chalk'
import creaetApi from '../../client/create_api.js'
import createFlags from '../../flags/statis.js'
import { PodService } from '../../services/pod.js'


export default class Pod extends Command {
  static flags = createFlags({ filename: true })

  // FIRE: 
  // identify or learn what apply and create command do do they differ or same and if they are the same can i just put one as alias 
  //
  // HOT: 
  // get files and maker sure files are not empty 
  // convert to json 
  // send to 
  // show the relevent message to usr
  // identify what are the falgs reqruired beside -f in createing or applying for pods 
  async run(): Promise<void> {
    const { flags } = await this.parse(Pod)
    const { filename, namespace } = flags
    const api = await creaetApi()

    if (!filename) {
      this.warn(chalk.yellow("you must path valid file name"))
      return
    }

    const RootPath = process.cwd() + filename
    const jsonfile = await yamlToJson(RootPath)

    const url = `/api/v1/namespaces/${namespace ?? "default"}/pods`
    PodService(api, url, namespace, jsonfile as JSON, "create")
    // HOT: this reponse types need to be fixed not console log but as the real kubect does it reponse with table type response 
    // and also should hadle fake pod create only with command line like (mykubectl run nginx )
  }

}






