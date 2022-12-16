a = "a"

runtime.callcontext({ kill = { memory = 1000 } }, function() while true do a = a .. a end end)

print(#a)
