package log

import (
	"fmt"
	"path/filepath"
	"runtime"
)

const (
	LOGINFO  = "log_info"
	LOGERROR = "log_err"
)

func Info(act interface{}) {

	pc, file, line, _ := runtime.Caller(1)
	fmt.Printf("%s:%s:%d:--method:%v--got: %#v-\n\n", LOGINFO, filepath.Base(file), line, runtime.FuncForPC(pc).Name(), act)

}

func Error(act interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	fmt.Printf("%s:%s:%d:--method:%v--got: %#v-\n\n", LOGERROR, filepath.Base(file), line, runtime.FuncForPC(pc).Name(), act)
}
