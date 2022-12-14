var head = require('./head');
var body = require('./body');

create = function (name) {
    return {
        name: name,
        head: head.create(),
        body: body.create()
    };
};

cat = create('咪咪')
console.log(cat)