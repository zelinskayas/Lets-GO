package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("ошибка1")
	err = fmt.Errorf("ошибка2:%w", err)
	err = fmt.Errorf("ошибка3:%w", err)
	err = errors.Unwrap(err)

	fmt.Println(err)
}
