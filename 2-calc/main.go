package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	operation := getUserOperation()
	numbers := getUserNumbers()

	switch operation {
	case "AVG":
		fmt.Printf("%.2f", getAVG(numbers))
	case "SUM":
		fmt.Println(getSum(numbers))
	case "MED":
		fmt.Println(getMedian(numbers))
	}
}

func getUserNumbers() []int {
	var input string

	for {
		fmt.Println("Введите числа через запятую")
		fmt.Scanln(&input)

		var numbersStrings []string = strings.Split(input, ",")
		numbers := make([]int, len(numbersStrings))
		for indx, val := range numbersStrings {
			i, err := strconv.Atoi(strings.TrimSpace(val))
			numbers[indx] = i

			if err != nil {
				numbers = nil
				fmt.Println("Введены некорректные значения")
				break
			}
		}

		if numbers == nil {
			continue
		} else {
			return numbers
		}
	}
}

func getUserOperation() string {
	var input string
	for {
		fmt.Println("Введите требуемую операцию: AVG - среднее, SUM - сумму, MED - медиану")
		fmt.Scanln(&input)
		if input == "AVG" || input == "SUM" || input == "MED" {
			return input
		}
	}
}

func getAVG(nums []int) float64 {
	return float64(getSum(nums)) / float64(len(nums))
}

func getSum(nums []int) int {
	total := 0
	for _, val := range nums {
		total += val
	}
	return total
}

func getMedian(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var median, center int

	slices.SortFunc(nums, func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		} else {
			return 0
		}
	})

	if len(nums)%2 == 0 {
		center = len(nums) / 2
		median = (nums[center] + nums[center-1]) / 2
	} else {
		median = nums[len(nums)/2]
	}

	return median
}
