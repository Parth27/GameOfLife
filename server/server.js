var http = require('http');

var express = require('express');
var server = express();
const bodyParser = require('body-parser');
var fs = require('fs');
var solutionGrid = '';
const startTime = Date.now();
var tokenStack = [];

fs.readFile('grid.txt','utf8', function(err,data) {
    if (err) throw err;
    solutionGrid = data;
})

fs.readFile('tokens.txt','utf8',function(err,data) {
    if (err) throw err;
    tokenStack = data.split('\n');
})

server.use(bodyParser.urlencoded({
    extended: false
}));
  
server.use(bodyParser.json());

server.post('/CheckGrid',function(req,res) {
    console.log('Received request');
    console.log(req.headers["user-agent"]);
    const device = req.headers["user-agent"];

    if (solutionGrid == req.body.grid) {
        console.log('Debug Successful');
        endTime = Date.now();
        var DebugTime = (endTime - startTime)/1000;
        console.log(String(DebugTime)+' seconds');
        const token = tokenStack.pop();
        fs.appendFile('DebugData.txt','\n'+token+', '+String(DebugTime)+' seconds, '+device+', '+req.body.language,function(err) {
            if (err) throw err;
        });
        res.send('Debug Successful');
        fs.writeFile('tokens.txt',tokenStack.join('\n'),function(err) {
            if (err) throw err;
        })
    }
    else {
        console.log('Grid does not match correct output');
        res.send('Grid does not match correct output');
    }
})

server.listen(5000, function () {
    console.log('Node server is running..');
});