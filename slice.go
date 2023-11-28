package main

import (
	"fmt"
	"strconv"
)

func main() {
	slice := make([]int, 0, 3)

	var data string

	for data != "X" {
		fmt.Scan(&data)
		newData, err := strconv.Atoi(data)
		if err != nil {
			fmt.Println(slice)
			break
		}

		slice = append(slice, newData)
		a := SelectionS(slice, len(slice))
		fmt.Println(a)
	}

}
func SelectionS(Slice []int, len int) []int {
	var min, temp int
	_, _ = min, temp
	for step := 0; step < len; step++ {
		min = step

		for i := step + 1; i < len; i++ {
			if Slice[i] < Slice[min] {
				min = i
			}
		}
		temp = Slice[step]
		Slice[step] = Slice[min]
		Slice[min] = temp
	}
	return Slice
}
