
const fs = require('fs')
const path = require('path')
const express = require('express')

const app = express()
const server = require('http').createServer(app)


app.get('/', (req, res) => {
  res.sendFile(path.join(__dirname, 'index.html'))
})

// scriptタグが取得するjs
app.get('/di.js', (req, res) => {
  // ここでtemplateにDBとかの値を入れる
  const template = fs.readFileSync('template.js', { encoding: 'utf8'})
  const sendJsFile = template.replace(/{{(\s|\S)*}}/, 234)
  res.send(sendJsFile)
})

server.listen(3000, () => {
  console.log('run server')
})