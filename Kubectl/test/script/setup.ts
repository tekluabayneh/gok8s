import fs from "fs"
import os from "os"


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
