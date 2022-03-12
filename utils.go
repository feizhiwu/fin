package fin

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
	"strings"
)

func Assert(guard bool, text string) {
	if !guard {
		log.Printf("\x1b[31;20m[ERROR] %s\x1b[0m\n", text)
		os.Exit(500)
	}
}

func GetFuncName(f interface{}) string {
	name := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	fmt.Println(name)
	arr := strings.Split(name, ".")
	return strings.Split(arr[len(arr)-1], "-")[0]
}

func InArray(n int, f func(int) bool) bool {
	for i := 0; i < n; i++ {
		if f(i) {
			return true
		}
	}
	return false
}
