package main

import (
	"fmt"
)

/*
1.1. — 2.1. Рекурсивное разбиение задачи на меньшие происходит до тех пор,
пока размер массива не достигнет единицы (любой массив длины 1 можно считать упорядоченным).

3.1. Соединение двух упорядоченных массивов в один.
Основную идею слияния двух отсортированных массивов можно объяснить на следующем примере.
Пусть мы имеем два уже отсортированных по возрастанию подмассива. Тогда:
3.2. Слияние двух подмассивов в третий результирующий массив.
На каждом шаге мы берём меньший из двух первых элементов подмассивов и записываем его в результирующий массив.
Счётчики номеров элементов результирующего массива и подмассива, из которого был взят элемент, увеличиваем на 1.
3.3. «Прицепление» остатка.
Когда один из подмассивов закончился, мы добавляем все оставшиеся элементы второго подмассива в результирующий
массив.
*/
func main() {
	data := []int{
		60, 54, 38, 87, 65, 21, 98, 74, 16, 59, 84, 65, 49, 31, 46, 98, 12,
	}
	// data = []int{
	// 	1, 2, 3, 4, 5,
	// }
	fmt.Println(len(data))

	mergeSort(data)

	fmt.Println(data)
	fmt.Println(len(data))
}

func mergeSort(data []int) []int {
	if len(data) == 1 || len(data) == 0 {
		return data
	}
	mid := len(data) / 2
	leftSide := mergeSort(data[:mid])
	rightSide := mergeSort(data[mid:])
	subResult := make([]int, len(leftSide)+len(rightSide))
	i, j, k := 0, 0, 0
	for ; i < len(leftSide) && j < len(rightSide); k++ {
		if leftSide[i] <= rightSide[j] {
			subResult[k] = leftSide[i]
			i++
		} else {
			subResult[k] = rightSide[j]
			j++
		}
	}

	for i < len(leftSide) {
		subResult[k] = leftSide[i]
		i++
		k++
	}

	for j < len(rightSide) {
		subResult[k] = rightSide[j]
		j++
		k++
	}

	copy(data, subResult)

	return data
}
