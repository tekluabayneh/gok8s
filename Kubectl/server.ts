import e from "express"
import express from "express"
const app = express()

app.get("/", (req, res) => {
  console.log("test")
})

app.listen(8000, (error) => {
  if (error) {
    console.error(error)
    return
  }

  console.log("Kubectl Server started successfully")
})



