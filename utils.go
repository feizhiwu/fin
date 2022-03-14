package fin

import (
	"crypto/md5"
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

func CopyParams(val []string, data map[string]interface{}) map[string]interface{} {
	params := make(map[string]interface{})
	for _, v := range val {
		if data[v] != nil {
			params[v] = data[v]
		}
	}
	return params
}

func EncryptPass(pass string) string {
	salt := Config("salt").(string)
	sum := md5.Sum([]byte(pass + salt))
	return fmt.Sprintf("%x", sum)
}
