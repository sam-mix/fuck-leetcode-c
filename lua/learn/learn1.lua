f = print
t = type
ft = function(x)
    f(t(x))
end


ft(nil)
ft(true)
ft(10.4 * 3)
ft("Hello world!")
ft(io.stdin)
ft(print)
ft(type)
ft({})
ft(type(nil))


f('\n\n\n')


ft(a)
a = 10
ft(a)
a = ""
ft(a)
a = {}
ft(a)
a = print
ft(a)
