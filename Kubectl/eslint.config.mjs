import { includeIgnoreFile } from '@eslint/compat'
import oclif from 'eslint-config-oclif'
import prettier from 'eslint-config-prettier'
import path from 'node:path'
import { fileURLToPath } from 'node:url'

const gitignorePath = path.resolve(
  path.dirname(fileURLToPath(import.meta.url)),
  '.gitignore',
)

export default [
  includeIgnoreFile(gitignorePath),
  ...oclif,

  {
    rules: {
      // Don't force alphabetical ordering
      'perfectionist/sort-imports': 'off',
      'perfectionist/sort-objects': 'off',
      'perfectionist/sort-object-types': 'off',

      // Don't enforce this stylistic preference
      'arrow-body-style': 'off',

      // Axios warning isn't particularly useful here
      'import/no-named-as-default-member': 'off',
    },
  },

  prettier,
]
