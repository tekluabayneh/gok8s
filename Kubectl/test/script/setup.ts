import fs from "node:fs"
import os from "node:os"
import path from "node:path";


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
    console.log("it reached hre")

    const defaultPath = path.join(os.homedir() + ".kube", "config")
    if (!fs.existsSync(defaultPath)) {
      fs.mkdirSync(path.dirname(defaultPath), { recursive: true })
      fs.writeFileSync(defaultPath, testKubeConfig)
      console.log("files is writeen")
    }
  } catch (error) {
    console.log(error)
  }
}

createConfigFile()
