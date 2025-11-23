package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

const (
	AVG = "AVG"
	SUM = "SUM"
	MED = "MED"
)

func main() {

	actions := []string{AVG, SUM, MED}
	var operation string

	for !slices.Contains(actions, strings.ToUpper(operation)) {
		fmt.Printf("Введите одну из операций %v!\n", actions)
		fmt.Scan(&operation)
	}

	operation = strings.ToUpper(operation)

	fmt.Println("Введите числа через запятую, затем нажмите enter, чтобы выполнить", operation)

	var nums string
	fmt.Scan(&nums)
	f := func(c rune) bool {
		return !unicode.IsNumber(c)
	}

	fields := strings.FieldsFunc(nums, f)
	fmt.Println(fields)

	numFields := make([]int, 0, len(fields))

	for _, v := range fields {
		num, _ := strconv.Atoi(v)
		numFields = append(numFields, num)
	}
	fmt.Println(numFields)
	if len(nums) > 0 {

		switch operation {
		case AVG:
			fmt.Println("Average = ", getAvg(numFields))
		case SUM:
			fmt.Println("Sum = ", getSum(numFields))
		case MED:
			fmt.Println("Med = ", getMed(numFields))
		default:
			fmt.Println("Wrong operation!")
		}
	} else {
		fmt.Println("Empty nums!")
	}
}

func getAvg(nums []int) float64 {
	size := len(nums)
	sum := getSum(nums)

	return float64(sum) / float64(size)
}

func getSum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	return sum
}

func getMed(nums []int) int {
	slices.Sort(nums)

	return nums[len(nums)/2]
}
