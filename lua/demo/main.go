package main

import (
	lua "github.com/yuin/gopher-lua"
	luajson "layeh.com/gopher-json"
)

func main() {
	L := lua.NewState()
	luajson.Preload(L)
}
