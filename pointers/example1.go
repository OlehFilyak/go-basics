package main

import (
	"fmt"
)

// Не могу не написать свое объяснение четырем строкам (операторов больше) в теле функции,
// так как тема очень простая, но требует аккуратности и последовательности. Итак, есть тело функции:
//         a := 200
//         b := &a
//         *b++
//         fmt.Println(a)
// В этом коде и а и b - это переменные, но имеющие разные типы:
// a имеет тип int и этот тип ей задает компилятор поскольку в правой части оператора присваивания (первая строка) стоит целочисленное значение  (200);
// b имеет тип указателя и этот тип ей задает компилятор, поскольку в правой части оператора присваивания
// (вторая строка) стоит операция взятия адреса (&) переменной, перед которой эта операция стоит, то есть a.
// В результате выполнения в переменной a, расположенной в памяти по адресу, сохраненному в переменной b,
// имеющей тип указателя, будет храниться значение 200.
// Третья строка. Поскольку унарная операция * имеет более высокий приоритет, чем операция ++ ,
// то "statement *p++ is the same as (*p)++" (https://go.dev/ref/spec#Operators),
// то есть: значение, хранящееся по адресу, хранящемуся в переменной b (то есть 200) будет увеличено на 1 (операция ++).
// В результате по адресу, хранящемуся в переменной b окажется новое значение, то есть 201.
// А поскольку адрес, хранящийся в переменной b есть адрес переменной a, то 201 это и есть новое значение переменной a.
// Новое значение переменной a (то есть 201) будет выведено оператором в четвертой строке.
// По сути в этом примере продемонстрировано два способа работы с одной и той же ячейкой памяти: посредством использования имени переменной,
// a и адреса этой ячейки, сохраненной в переменной b. Скриншот ниже содержит слегка, модифицированный код примера, иллюстрирующий мой комментарий.
// В зеленом прямоугольнике результат выполнения:

func main() {
	a := 200
	b := &a
	*b++
	// a := 200
	// var b *int = &a
	fmt.Println(a)
}
