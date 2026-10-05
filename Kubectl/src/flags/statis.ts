import { Flags } from '@oclif/core'

type FlagOptions = {
  namespaces?: boolean
  allNamespaces?: boolean
  output?: boolean
  selector?: boolean
  fieldSelector?: boolean
  watch?: boolean
  watchOnly?: boolean
  showLabels?: boolean
  showKind?: boolean
  sortBy?: boolean
  noHeaders?: boolean
  ignoreNotFound?: boolean
  filename?: boolean
  kustomize?: boolean
  recursive?: boolean
  chunkSize?: boolean
  outputWatchEvents?: boolean
  raw?: boolean
  subresource?: boolean
  context?: boolean
  kubeconfig?: boolean
}

const createFlags = (options: FlagOptions = {}) => {
  return {
    namespace: Flags.string({
      char: 'n',
      description: 'namespace scope for this request',
      required: options.namespaces ?? false,
    }),

    allNamespaces: Flags.boolean({
      char: 'A',
      description: 'list the requested object(s) across all namespaces',
      required: options.allNamespaces ?? false,
    }),

    output: Flags.string({
      char: 'o',
      description: 'output format (json|yaml|wide|name)',
      required: options.output ?? false,
    }),

    selector: Flags.string({
      char: 'l',
      description: 'label selector to filter results',
      required: options.selector ?? false,
    }),

    fieldSelector: Flags.string({
      description: 'field selector to filter results',
      required: options.fieldSelector ?? false,
    }),

    watch: Flags.boolean({
      char: 'w',
      description: 'watch for changes after listing/getting',
      required: options.watch ?? false,
    }),

    watchOnly: Flags.boolean({
      description: 'watch for changes without doing an initial list',
      required: options.watchOnly ?? false,
    }),

    showLabels: Flags.boolean({
      description: 'show all labels as the last column',
      required: options.showLabels ?? false,
    }),

    showKind: Flags.boolean({
      description: 'show the kind name for each resource',
      required: options.showKind ?? false,
    }),

    sortBy: Flags.string({
      description: 'sort list using a jsonpath expression',
      required: options.sortBy ?? false,
    }),

    noHeaders: Flags.boolean({
      description: 'omit headers from the output',
      required: options.noHeaders ?? false,
    }),

    ignoreNotFound: Flags.boolean({
      description: 'treat "resource not found" as a successful exit',
      required: options.ignoreNotFound ?? false,
    }),

    filename: Flags.string({
      char: 'f',
      description: 'file, directory, or URL to identify resources',
      required: options.filename ?? false,
      multiple: false,
    }),

    kustomize: Flags.string({
      char: 'k',
      description: 'process a kustomization directory',
      required: options.kustomize ?? false,
    }),

    recursive: Flags.boolean({
      char: 'R',
      description: 'process the directory used in -f recursively',
      required: options.recursive ?? false,
    }),

    chunkSize: Flags.integer({
      description: 'batch size for large list requests',
      required: options.chunkSize ?? false,
    }),

    outputWatchEvents: Flags.boolean({
      description: 'output watch event objects with type and object',
      required: options.outputWatchEvents ?? false,
    }),

    raw: Flags.string({
      description: 'raw URI to request from the server',
      required: options.raw ?? false,
    }),

    subresource: Flags.string({
      description: 'fetch a named subresource (status|scale) instead of the object',
      required: options.subresource ?? false,
    }),

    context: Flags.string({
      description: 'name of the kubeconfig context to use',
      required: options.context ?? false,
    }),

    kubeconfig: Flags.string({
      description: 'path to the kubeconfig file to use',
      required: options.kubeconfig ?? false,
    }),
  }
}

export default createFlags
