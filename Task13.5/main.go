package main

import (
	"fmt"
	v10 "v10"
	v11 "v11"
	v20 "v20"
	v21 "v21"
)

func main() {
	err := v10.Do("patients", "res10")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Файл успешно записан. v1.0.0")
	}

	err = v11.Do("patients", "res11")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Файл успешно записан. v1.1.0")
	}

	err = v20.Do("patients", "res20")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Файл успешно записан. v2.0.0")
	}

	err = v21.Do("patients", "res21")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Файл успешно записан. v2.1.0")
	}
}
