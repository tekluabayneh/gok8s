import e from "express"
import express from "express"
import { load, loadAll } from 'js-yaml'
import { readFileSync } from 'node:fs'
const app = express()

app.get("/api/v1/kubelet", (req, res) => {
  console.log("test")
  try {
    let val = load(readFileSync("./pod.yaml", 'utf8'))
    console.log(val)
    res.send(val)
  } catch (error) {
    res.send(error)
    console.log(error)
  }

})

app.listen(8000, (error) => {
  if (error) {
    console.error(error)
    return
  }

  console.log("Kubectl Server started successfully")
})



