Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
Вывод 

<nil>
false

При первом выводе, мы выводим значение err, которое равно nil
Но при сравнении nil и err происходит сравнение интерфейса error с nil,
а интерфейс считается равным nil, когда его значение равно nil и его тип равен nil
То есть мы сравниваем [nil, *os.PathError] и [nil, nil]. 
И на выходе получаем false

Пустой интерфейс в Go это отдельная структура содержащая тип хранимых данных
и ссылку на значение этих данных

А интерфейс, который реализует методы содержит в себе 
структуру(в ней хранится описание интерфейса, ссылку на список методов тип хеш) 
и ссылку на данные

...

```
