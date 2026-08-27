import { Command, Flags } from '@oclif/core'
import { yamlToJson } from '../../utils/yamlToJson.js'


export default class Pods extends Command {
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
    filename: Flags.string({ char: 'f', description: 'file, directory, or URL to identify resources', required: false, multiple: true }),
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
    const { flags } = await this.parse(Pods)
    const { filename, namespace, output, context, kubeconfig } = flags
    // TODO 
    // implement create pod method 
    // if (namespace == "default") {
    // api will use 
    // }
    // "body"{ 
    //  "namespace": namespace
    //  }
    // instade of using if statment just use waht ever you got not if statment this is soemthing you can pass like in the json body of api



    // console.log("") // get the the apiServer Base url from conf or someware not from here
    // const res = await fetch("")
    // const data = await res.json()
    // console.log(data)
    console.log("filename", filename)
    if (!filename) {
      return
    }

    if (filename[0] == "/") {
      console.log("you must path valid file name")
      return
    }

    const jsonfile = await yamlToJson(filename[0])
    console.log("json file", jsonfile)
    console.log("namespace", namespace)
    // console.log(filename, namespace, output, context, kubeconfig)
  }

}


// Arguments / flags — what the user explicitly provides.
// Resource body/data — the object or configuration being created/modified.
// User identity / privileges — who is making the request and what they are allowed to do.
// Namespace — the Kubernetes namespace in which the resource operates.
// Cluster/context — which cluster the command is targeting.
// Authentication credentials — how the client proves its identity.
// API server address — where the Kubernetes API server is.
// Configuration — usually the kubeconfig/context that ties cluster + user + credentials together.
// Resource type/name — e.g. Pod and its name.
// Request metadata — things such as HTTP headers/content type where relevant.
// Output preferences — e.g. normal output, YAML, JSON, wide output, etc.
// Validation/defaulting — values the CLI needs to validate or fill in before making the request.
//
// But don't treat all of these as things every command receives directly.
//
// // i have to taken core of auth too 

