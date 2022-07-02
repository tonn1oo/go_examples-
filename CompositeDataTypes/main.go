package main

import "fmt"

/*составные типы данных:
Array -Массивы
Slice - Слайсы
Map - Отоброжение
Struct - Структура
*/

func main() {
	var x1 [5]int //массив из 5 целых числел
	var x2 [0]int
	var x3 [0]string

	var arr [3]int = [3]int{1, 2, 3} /// [1,2,3]
	var arr2 = [3]int{1, 2, 3}       ///более короткая форма
	arr3 := [3]int{1, 2, 3}          /// [1,2,3] еще короче

	/////arr3 = [2]int{1,2} /// error нельзя присвоить переменной типа [3]int значение типа [2]int

	arr5 := [...]int32{1, 2, 3, 4, 5} // [1,2,3,4,5]

	s := len(arr5)

	fmt.Println(arr, arr2, arr3, arr5, x1, x2, x3, s)
}

func slice() {

	var list []int64

	l := len(list) //0
	c := cap(list) //0

	list = []int64{1, 2, 3, 4, 5} ///[1,2,3,4,5]
	l = len(list)                 //5
	c = cap(list)                 ///5

	list = make([]int64, 0, 5) //[]
	l = len(list)              // 0
	c = cap(list)              //5

	list = make([]int64, 5, 5) ///[0,0,0,0,0]
	//list = make([]string, 5, 5) //["","","","",""]
	l = len(list) /// 5
	c = cap(list) /// 5

	list = []int64{1, 2, 3, 4} /// [1,2,3,4,5]
	list = append(list, 5)
	l = len(list) ///
	c = cap(list) ///

	list = make([]int64, 0, 3) // [] len:0, cap;3
	list = append(list, 1)     // [1] len:1, cap;3
	list = append(list, 2)     // [1,2] len:2, cap;3
	list = append(list, 3)     // [1,2,3] len:3, cap;3
	list = append(list, 4)     // [1,2,3,4] len:4, cap;6

	fmt.Println(l, c)
}
