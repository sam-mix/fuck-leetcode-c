function foo(x)
    print(x)
    error("do not do this")
end

function bar(x)
    print(x)
    foo(x * x)
end

bar(2)
