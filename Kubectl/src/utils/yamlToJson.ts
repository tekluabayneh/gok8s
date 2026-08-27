import { parse } from "yaml"
import fs from "fs"

export const yamlToJson = async (Path: string): Promise<"" | JSON> => {
  try {
    const RootDir = process.cwd()
    fs.readFile(RootDir + Path, (error, file) => {
      if (error) {
        console.error(`Path does not exist: ${Path}`);
        return
      }

      if (file.length == 0) {
        console.error(`File is empty: ${Path}`);
        return;
      }

      const yaml = fs.readFileSync(RootDir + Path, "utf-8")
      const jsonFile = parse(yaml)
      return jsonFile
    })
    return ""
  } catch (error) {
    console.log(error)
    return ""
  }
}





