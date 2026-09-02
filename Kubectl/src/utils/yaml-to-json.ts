import { parse } from "yaml"
import fs from "node:fs"
import chalk from "chalk";

export const yamlToJson = async <T>(Path: string): Promise<T | void> => {
  try {
    fs.stat(Path, (error, stats) => {

      if (error) {
        console.error(chalk.red(`Path does not exist: ${Path}`));
        return
      }


      if (stats.isFile() && stats.size === 0) {
        console.error(chalk.yellow("empty files"));
      }
    })

    const yaml = fs.readFileSync(Path, "utf8")
    const jsonFile = parse(yaml)

    return jsonFile
  } catch (error) {
    console.error(chalk.red(`Path does not exist: ${Path}`, error));
  }
}




