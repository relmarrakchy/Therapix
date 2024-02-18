const express = require('express');
const server = express();
const cors = require('cors');
const axios = require('axios')

server.use(express.json());
server.use(cors());

server.use((req, res, next) => {
    console.log(req.path, req.method);
    next();
});

server.post("/signup", async (req, res) => {
    let response
    try {
        response = await axios.post("http://localhost:3010/signup", req.body)
        console.log(response.data)
    } catch (err) {
        console.log(err)
    }
    res.json(response.data);
});

server.post("/login", async (req, res) => {
    let response
    try {
        response = await axios.post("http://localhost:3010/login", req.body)
        console.log(response.data)
    } catch (err) {
        console.log(err)
    }
    res.json(response.data);
});

server.post("/response", async (req, res) => {
    let response
    try {
        response = await axios.post("http://127.0.0.1:3011/response", req.body)
        console.log(response.data)
    } catch (err) {
        console.log(err)
    }
    res.json(response.data);
});

const port = 3001;
server.listen(port, () => {
  console.log(`Express server running on port ${port}`);
});