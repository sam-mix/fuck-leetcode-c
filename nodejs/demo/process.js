var child_process = require('child_process');
var util = require('util');

function copy(source, target, callback) {
    child_process.exec(
        util.format('cp -r %s/* %s', source, target), callback);
}

copy('/Users/m/fuck-leetcode-c/nodejs/resources/a', '/Users/m/fuck-leetcode-c/nodejs/resources/b', function (err) {
    console.log(err)
});