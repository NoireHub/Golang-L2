Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
Вывод 
2
1

В функции test возвращается переменная x, при вызове return выполняется 
defer и только после него возвращается x, поэтому получаем 2

В функции anotherTest возвращается значение int, поэтому при вызове defer
локальная переменная x уже не имеет отношения к возвращаемому значению

Defer является ключевым словом
При добавлении defer к функции, данная функция будет выполнена
по окончании функции, в которой она была объявлена. 
Порядок вызова работает как в стеке (чем раньше добавлен, тем позже выполнится)

```