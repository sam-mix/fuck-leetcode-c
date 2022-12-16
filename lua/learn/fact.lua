-- 定义一个计算阶乘的函数
function fact(n)
    if n == 0 then
        return 1
    else
        return n * fact(n - 1)
    end
end

-- n = io.read("*n")
-- print(fact(n))

for i = 0, 5 do
    print(i, fact(i))
end

-- lua "/Users/m/fuck-leetcode-c/lua/learn/fact.lua"
