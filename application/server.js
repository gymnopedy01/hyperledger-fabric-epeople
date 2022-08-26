const express = require('express')
const app = express()

const path = require('path')

const PORT = 8080
const HOST = "0.0.0.0"

app.use(express.static(path.join(__dirname, "views")))

app.get('/', (req,res) => {
    res.sendFile(__dirname+"index.html")
})

app.listen(PORT, HOST)
console.log(`EPeople Server launched successfully http://${HOST}:${PORT}`)