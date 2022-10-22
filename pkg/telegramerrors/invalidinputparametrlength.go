package telegramerrors

import "fmt"

type InvalidInputParametrLength struct {
	functionlen int32
	needlen     int32
}

func NewInvalidInputParametrLength(len int32, lenanswer int32) *InvalidInputParametrLength {
	a := &InvalidInputParametrLength{}
	a.functionlen = lenanswer
	a.needlen = len
	return a
}

func (i *InvalidInputParametrLength) Error() string {
	return fmt.Sprintf("Invalid input parametr length. Need: %d, Length input parametrs: %d", i.needlen, i.functionlen)
}
