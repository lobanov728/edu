package main

import "fmt"

func main() {
	data := []int{
		60, 54, 38, 87, 65, 21, 98, 74, 16, 59, 84, 65, 49, 31, 46, 98, 12,
	}
	data = []int{
		60, 54, 70, 38, 87, 65,
	}
	fmt.Println(len(data))

	insertSort(data)

	fmt.Println(data)
	fmt.Println(len(data))
}

/*
for j = 2 to A.length do
    key = A[j]
    i = j-1
    while (int i >= 0 and A[i] > key) do
        A[i + 1] = A[i]
        i = i - 1
    end while
    A[i+1] = key
end
*/

func insertSort(data []int) {
	for j := 1; j < len(data); j++ {
		key := data[j]
		i := j - 1
		for i >= 0 && data[i] > key {
			data[i+1] = data[i]
			i = i - 1
		}
		fmt.Println(data)
		data[i+1] = key
	}
}
