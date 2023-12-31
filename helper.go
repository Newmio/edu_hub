package ed

import (
	"errors"
	"fmt"
	"runtime"
)

const TIMEFORMAT = "02-01-2006 15:04:05.000"

func ErrTrace(err error, str string) error {
	fmt.Println(str)
	return errors.New(err.Error() + "\n" + str)
}

func ErrDbTrace(err error, sql, str string) error {
	fmt.Println(str)
	return errors.New(err.Error() + "\n" + sql + "\n" + str)
}

func Trace() string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	return fmt.Sprintf("<<<<< %s:%d %s >>>>>", file, line, f.Name())
}
