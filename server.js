
const fs = require('fs')
const path = require('path')
const express = require('express')

const app = express()
const server = require('http').createServer(app)

const ReplaceTarget = /\$(\s|\S)*\$/

app.get('/', (req, res) => {
  res.sendFile(path.join(__dirname, 'index.html'))
})

// 毎回ファイルを読み込ませないためにオンメモリで情報は保持しておく
const template = fs.readFileSync('template.js', { encoding: 'utf8' })
const templateInjection = fs.readFileSync('templateInjection.js', { encoding: 'utf8' })

// scriptタグが取得するjsを作成し送り返す
app.get('/di.js', (req, res) => {
  // ここでtemplateにDBとかの値を入れる
  const sendJsFile = template.replace(ReplaceTarget, 234)
  res.send(sendJsFile)
})

// ユーザの入力内容により任意のJavaScriptが実行される
app.get('/injection.js', (req, res) => {
  const sendJsFile = templateInjection.replace(ReplaceTarget, "'abc'; alert('konnitiha konnitiha'); //")
  res.send(sendJsFile)
})

server.listen(3000, () => {
  console.log('run server')
})