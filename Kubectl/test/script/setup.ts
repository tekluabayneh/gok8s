import { error } from "@oclif/core/errors";
import fs, { stat } from "node:fs"
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

    // much check first if files doe snot exist create the fils and add to there 
    //
    const defaultPath = os.homedir() + "/.kube/config.yaml"
    if (fs.existsSync(defaultPath)) {
      fs.mkdirSync(path.dirname(defaultPath), { recursive: true })
      fs.writeFileSync(defaultPath, testKubeConfig)
      console.log(error)
    }
  } catch (error) {
    console.log(error)
  }
}

createConfigFile()
