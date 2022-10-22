package errors

import (
	"fmt"
	"runtime"
)

type InvalidInputParametrType struct {
}

func (i *InvalidInputParametrType) Error() string {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%s, %d has invalid input parametr type", file, line)
}
