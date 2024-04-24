package main

import (
	"errors"
	"fmt"
)

type myFirstError struct {
	status  int
	message string
}

func (e myFirstError) Error() string {
	return e.message
}

func main() {
	err1 := errors.New("ошибка1")
	err2 := fmt.Errorf("ошибка2:%w", err1)
	err3 := fmt.Errorf("ошибка3:%w", err2)
	err4 := myFirstError{
		status:  1,
		message: "ошибка моего типа myFirstError",
	}
	//fmt.Println(err3)
	//fmt.Println(err4)
	fmt.Printf("ошибка типа %T была в %v: %v", err4, err3, errors.As(err3, new(myFirstError)))
}
