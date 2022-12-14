package main

import (
	"errors"
	"os"
	"path/filepath"
	"runtime/pprof"
)

var (
	profileName = "cpuprofile.out"
)

// CreateFile 用于在当前目录下创建一个指定名称的文件。
// 若同名文件已存在，则清空并复用。
func CreateFile(dir, name string) (*os.File, error) {
	if dir == "" {
		var err error
		dir, err = os.Getwd()
		if err != nil {
			return nil, err
		}
	}
	path := filepath.Join(dir, name)
	return os.Create(path)
}

func joinSlice() []string {
	var arr []string
	// 为了进行性能分析,这里在己知元素大小的情况下,还是使用 append()
	// 函数不断地添加切片。性能较低,在实际中应该避免,这里为了性能分析,故意这样写。
	for i := 0; i < 1000000; i++ {
		arr = append(arr, "arr")
	}

	return arr
}

func main() {
	f, _ := CreateFile("", profileName)

	defer f.Close()
	startCPUProfile(f)

	joinSlice()

	stopCPUProfile()
}

func startCPUProfile(f *os.File) error {
	if f == nil {
		return errors.New("nil file")
	}
	return pprof.StartCPUProfile(f)
}

func stopCPUProfile() {
	pprof.StopCPUProfile()
}
