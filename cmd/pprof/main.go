package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

// 执行 go run main -help 查看帮助信息
// 执行 go run main -cpuprofile cpu.prof 生成cpu性能分析文件
func main() {
	var cpuprofile = flag.String("cpuprofile", "", "请输入 -cpuprofile 指定cpu性能分析文件名称")
	//在所有flag都注册之后，调用：flag.Parse()
	flag.Parse()
	f, err := os.Create(*cpuprofile)
	if err != nil {
		log.Fatal("could not create CPU profile: ", err)
	}
	// StartCPUProfile为当前进程开启CPU profile。
	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("could not start CPU profile: ", err)
	}
	// StopCPUProfile会停止当前的CPU profile（如果有）
	defer pprof.StopCPUProfile()

	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	fmt.Printf("sum=%d\n", sum)
}
