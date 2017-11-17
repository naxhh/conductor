var http = require('http');

var handleRequest = function(request, response) {
	  console.log('Received request for URL: ' + request.url);
    response.writeHead(200);
    response.end('Hello World again!');
};

console.log('Listening in 8080');
var www = http.createServer(handleRequest);
www.listen(8080);
