const express = require('express')
const httpProxy = require('http-proxy')
const cors = require("cors")

const app = express()

app.use(cors())

const proxy = httpProxy.createProxyServer({})

const servers = [
  'http://localhost:3001',
  // Add more backend servers if needed
]

// Load balancing logic
let currentServerIndex = 0;
const getNextServer = () => {
  currentServerIndex = (currentServerIndex + 1) % servers.length
  return servers[currentServerIndex]
}

app.use((req, res, next) => {
  console.log(req.path, req.method)
  next()
})

app.post("/signup", (req, res) => {
  const target = getNextServer()
  proxy.web(req, res, { target })
})

app.post("/login", (req, res) => {
  const target = getNextServer()
  proxy.web(req, res, { target })
})

app.post("/response", (req, res) => {
  const target = getNextServer()
  proxy.web(req, res, { target })
})

// Start the server
const port = 4000;
app.listen(port, () => {
  console.log(`Load balancer server running on port ${port}`)
});
