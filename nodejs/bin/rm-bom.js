var fs = require('fs')

function readText(pathname) {
    var bin = fs.readFileSync(pathname);
    console.log(bin.toString('utf-8'))
    if (bin[0] === 0xEF && bin[1] === 0xBB && bin[2] === 0xBF) {
        bin = bin.subarray(3);
    }

    console.log(bin.toString('utf-8'))

    return bin.toString('utf-8');
}

readText('/Users/m/fuck-leetcode-c/nodejs/bin/bom-utf8-file.txt')