package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

var funcs = map[string]func([]int) float64{
	"AVG": getAVG,
	"SUM": getSum,
	"MED": getMedian,
}

func main() {
	operation := getUserOperation()
	numbers := getUserNumbers()
	result := funcs[operation](numbers)
	fmt.Printf("%.2f", result)
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

func getSum(nums []int) float64 {
	total := 0
	for _, val := range nums {
		total += val
	}
	return float64(total)
}

func getMedian(nums []int) float64 {
	if len(nums) == 0 {
		return 0
	}
	var center int
	var median float64

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
		median = float64((nums[center] + nums[center-1])) / 2
	} else {
		median = float64(nums[len(nums)/2])
	}

	return median
}
