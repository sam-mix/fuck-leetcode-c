package main

import (
	"io/ioutil"

	lua "github.com/yuin/gopher-lua"
)

func main() {
	bytes, _ := ioutil.ReadFile("/Users/m/fuck-leetcode-c/lua/fibr.lua")
	L := lua.NewState()
	defer L.Close()
	if err := L.DoString(string(bytes)); err != nil {
		panic(err)
	}
}
