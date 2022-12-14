var net = require('net')

var options = {
    port: 80,
    host: 'www.example.com'
}

var client = net.connect(options, function () {
    client.write([
        'GET / HTTP/1.1',
        'Host: www.baidu.com',
        'User-Agent: curl/7.84.0',
        'Accept: */*',
    ].join('\n'))
})

client.on('data', function (data) {
console.log(data.toString())
client.end()
})