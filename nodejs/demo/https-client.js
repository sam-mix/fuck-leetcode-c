var https = require('https')

var options = {
    hostname: 'www.example.com',
    port: 443,
    path: '/',
    method: 'GET'
};

var request = https.request(options, function (response) {
    var body = []
    response.on('data', function (chunk) {
        body.push(chunk)
    })
    console.log(body.toString())
})

request.sendDate()

request.end()