import { Command, Flags } from '@oclif/core'
import api from '../../client/client.js'
import axios from 'axios'
import chalk from 'chalk'
import renderToTerminal from '../../utils/Render-to-terminal.js'
import type { RsPodtype } from '../../../types/configtypes.js'
export default class Namespace extends Command {
  static aliases = ['get:ns']
  static flags = {
    namespace: Flags.string({ char: 'n', description: 'namespace scope for this request', required: false }),
    allNamespaces: Flags.boolean({ char: 'A', description: 'list the requested object(s) across all namespaces', required: false }),
    output: Flags.string({ char: 'o', description: 'output format (json|yaml|wide|name)', required: false }),
    selector: Flags.string({ char: 'l', description: 'label selector to filter results', required: false }),
    fieldSelector: Flags.string({ description: 'field selector to filter results', required: false }),
    watch: Flags.boolean({ char: 'w', description: 'watch for changes after listing/getting', required: false }),
    watchOnly: Flags.boolean({ description: 'watch for changes without doing an initial list', required: false }),
    showLabels: Flags.boolean({ description: 'show all labels as the last column', required: false }),
    showKind: Flags.boolean({ description: 'show the kind name for each resource', required: false }),
    sortBy: Flags.string({ description: 'sort list using a jsonpath expression', required: false }),
    noHeaders: Flags.boolean({ description: 'omit headers from the output', required: false }),
    ignoreNotFound: Flags.boolean({ description: 'treat "resource not found" as a successful exit', required: false }),
    filename: Flags.string({ char: 'f', description: 'file, directory, or URL to identify resources', required: false, multiple: false }),
    kustomize: Flags.string({ char: 'k', description: 'process a kustomization directory', required: false }),
    recursive: Flags.boolean({ char: 'R', description: 'process the directory used in -f recursively', required: false }),
    chunkSize: Flags.integer({ description: 'batch size for large list requests', required: false }),
    outputWatchEvents: Flags.boolean({ description: 'output watch event objects with type and object', required: false }),
    raw: Flags.string({ description: 'raw URI to request from the server', required: false }),
    subresource: Flags.string({ description: 'fetch a named subresource (status|scale) instead of the object', required: false }),
    context: Flags.string({ description: 'name of the kubeconfig context to use', required: false }),
    kubeconfig: Flags.string({ description: 'path to the kubeconfig file to use', required: false }),
  }


  async run(): Promise<void> {
    const { flags } = await this.parse(Namespace)
    const { namespace } = flags

    try {
      const res = await api.get(`/api/v1/namespaces`)
      //TODO: 
      //this renderToTerminal function should be flexable enough to for all of terminal render things 
      console.dir(res.data.items, { depth: null })

      renderToTerminal(res.data.items, "Namespace")
      if (res.data.items.length === 0) {
        console.log(`no resource are found in the ${namespace ?? "default"} namespace`)
      }

    } catch (error) {
      if (axios.isAxiosError(error)) {
        console.log(error.response?.data.message)
      } else {
        console.log(chalk.red("something went wrong", error))
      }
    }

  }

}


