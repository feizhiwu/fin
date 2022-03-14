package fin

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

// print stack trace for debug
func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:])
	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				//预设错误不输出stack trace
				if _, ok := err.(int); !ok {
					message := fmt.Sprintf("%s", err)
					log.Printf("%s\n\n", trace(message))
				}
				c.NewDisplay().Show(err)
			}
		}()
		c.Next()
	}
}
