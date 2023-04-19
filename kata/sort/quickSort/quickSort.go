package main

import (
	"fmt"
)

var iteration int

/*
1 Выбрать из массива элемент, называемый опорным.
  Это может быть любой из элементов массива.
  От выбора опорного элемента не зависит корректность алгоритма,
  но в отдельных случаях может сильно зависеть его эффективность (см. ниже).
2 Сравнить все остальные элементы с опорным и переставить их в массиве так,
  чтобы разбить массив на три непрерывных отрезка, следующих друг за другом:
  «элементы меньшие опорного», «равные» и «большие».
3 Для отрезков «меньших» и «больших» значений выполнить рекурсивно
  ту же последовательность операций, если длина отрезка больше единицы.
*/

func quicksort(data []int) {
	if len(data) < 2 {
		return
	}
	pivot := data[0]
	l := 1
	r := len(data) - 1
	for l <= r {
		if data[l] <= pivot {
			l++
		} else {
			data[r], data[l] = data[l], data[r]
			r--
		}
	}
	data[l-1], data[0] = data[0], data[l-1]
	quicksort(data[:l-1])
	quicksort(data[l:])

}

// 60, 54, 38, 87, 65, 21, 98, 74, 16, 59, 84, 65, 49, 31, 46, 98, 12
// 60, 54, 38, 87, 65, 40, 66
// 54 38 65 87 40 66
// 54 38 40 60 87 66 65
func main() {
	data := []int{
		60, 54, 38, 87, 65, 21, 98, 74, 16, 59, 84, 65, 49, 31, 46, 98, 12,
	}
	fmt.Println(len(data))

	quicksort(data)

	fmt.Println(data)
}
