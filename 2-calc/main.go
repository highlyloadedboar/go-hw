package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
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
		fmt.Printf("Введите одну из операций %v!\nДля выополнения выбранной операции введите любую букву\n", actions)
		fmt.Scan(&operation)
	}

	nums := []int{}
	for {
		var input string
		fmt.Scan(&input)
		num, err := strconv.Atoi(strings.TrimSpace(input))

		if err != nil {
			break
		}
		nums = append(nums, num)
	}

	if len(nums) > 0 {

		switch operation {
		case AVG:
			fmt.Println("Average = ", getAvg(nums))
		case SUM:
			fmt.Println("Sum = ", getSum(nums))
		case MED:
			fmt.Println("Med = ", getMed(nums))
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
