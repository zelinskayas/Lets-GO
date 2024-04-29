package main

import (
	"errors"
	"fmt"
)

type Bird interface {
	Sing() string
}

type Duck struct {
	voice string
}

func (d *Duck) Sing() string {
	if d == nil {
		return "пустой указатель, инициализируйте переменную"
	}
	return d.voice
}

func main() {
	//var d *Duck
	d := &Duck{"кря кря"}

	song, err := Sing(d)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)
}

func Sing(b Bird) (string, error) {
	if b == nil {
		return "", errors.New("Ошибка пения!")
	}

	if val, ok := b.(*Duck); ok && val == nil {
		return "", errors.New("Утка сломалась и не может петь!")
	}
	return b.Sing(), nil
}

/*
Паника возникает из-за того что в функцию func Sing(b Bird) (string, error)
передается переменная d *Duck, объявленная как указатель на Duck, но ей не присвоено
значение, то есть она равно nil.
В функции func Sing(b Bird) (string, error)  есть проверка:
if b != nil {
		return b.Sing(), nil
	}
в данном случае b не равно nil, так как интерфейс, содержащий nil указатель, сам по себе не равен nil,
из-за того, что интерфейс содержит информацию и о типе и о значении.
В нашем примере тип это *Duck, а значение nil, интерфейс не будет равен nil.
В конечном итоге метод Sing() вызывается на nil указателе, из-за чего возникает паника.

варианты решения проблемы:
1. инициализируем переменную Duck
этот код:
var d *Duck
изменим на:
d := &Duck{"кря кря"}
указатель будет равен конкретному значению, получим:
"кря кря".

2. добавим проверку в метод func (d *Duck) Sing() string на nil
func (d *Duck) Sing() string {
	if d == nil {
		return "пустой указатель, инициализируйте переменную"
	}
	return d.voice
}
в ошибке можно вывести какую-то информацию и избежать паники,
в данной случае выведется ошибка:
"пустой указатель, инициализируйте переменную".

3. изменим функцию func Sing(b Bird) (string, error)
добавим проверку на nil конкретного типа *Duck:
func Sing(b Bird) (string, error) {
	if b == nil {
		return "", errors.New("Ошибка пения!")
	}

	if val, ok := b.(*Duck); ok && val == nil {
		return "", errors.New("Утка сломалась и не может петь!")
	}
	return b.Sing(), nil
}
вместо паники получим сообщение: "Утка сломалась и не может петь!".
 */
