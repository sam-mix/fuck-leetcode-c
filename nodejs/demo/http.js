var http = require('http');

http.createServer(function (request, response) {
    console.log(request.url)
    console.log(request.method)
    console.log(request.headers)
    response.writeHead(200, { 'Content-Type': 'text-plain' });
    response.end('Hello World\n');
}).listen(8124);