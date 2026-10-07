import fs from "node:fs"
import os from "node:os"


export const testKubeConfig = `
apiVersion: v1
kind: Config
clusters:
  - name: test
    cluster:
      server: https://test-server
contexts:
  - name: test-context
    context:
      cluster: test
      user: test-user
current-context: test-context
users:
  - name: test-user
    user:
      token: test-token
`;


function createConfigFile() {
  try {
    const defaultPath = os.homedir() + "/.kube/configfake.yaml"
    fs.writeFileSync(defaultPath, testKubeConfig)
  } catch (error) {
    console.log(error)
  }
}

createConfigFile()
