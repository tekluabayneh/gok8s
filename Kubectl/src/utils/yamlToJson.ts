import { parse } from "yaml"
import fs from "fs"
import chalk from "chalk";

export const yamlToJson = async (Path: string): Promise<void | JSON> => {
  try {
    const FullPath = process.cwd() + Path
    fs.stat(FullPath, (error, stats) => {

      if (error) {
        console.error(chalk.red(`Path does not exist: ${Path}`));
        return
      }


      if (stats.isFile()) {
        if (stats.size == 0) {
          console.error(chalk.yellow("empty files"));
          return
        }

      }
    })

    const yaml = fs.readFileSync(FullPath, "utf-8")
    const jsonFile = parse(yaml)
    return jsonFile
  } catch (error) {
    console.error(chalk.red(`Path does not exist: ${Path}`));
  }
}




