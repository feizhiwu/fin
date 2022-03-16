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

func recovery(c *Context) {
	defer func() {
		if err := recover(); err != nil {
			for _, v := range dbs {
				RollbackDB(v)
			}
			message := fmt.Sprintf("%s", err)
			logFile := logFile()
			log.New(logFile, "", log.LstdFlags).Printf("%s\n\n", trace(message))
			log.New(logFile, "", log.LstdFlags).Printf("%s", "----------------------------------------------------------------------")
			if Mode() == DebugMode {
				log.Printf("%s\n\n", trace(message))
				log.Printf("%s", "----------------------------------------------------------------------")
			}
			c.NewDisplay().Show(err)
		}
	}()
	c.Next()
}
