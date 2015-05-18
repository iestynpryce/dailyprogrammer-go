package main

// Find sad cycles

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func pow(a, b int) int {
	p := 1
	for b > 0 {
		if b&1 != 0 {
			p *= a
		}
		b >>= 1
		a *= a
	}
	return p
}

func power_each_digit(power int, number string) int {
	var sum int = 0

	for _, c := range number {
		num, err := strconv.Atoi(string(c))
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
		sum += pow(num, power)
	}
	return sum
}

func find_cycle(numbers []int) []int {
	// use floyd's algolithm
	mu := 0
	lam := 1
	if len(numbers) > 1 {
		i := 1
		j := 2
		tortoise := numbers[i]
		hare := numbers[j]
		for tortoise != hare {
			i += 1
			j += 2
			tortoise = numbers[i]
			hare = numbers[j]
		}

		i = 0
		tortoise = numbers[i]
		for tortoise != hare {
			i += 1
			j += 1
			tortoise = numbers[i]
			hare = numbers[j]
			mu += 1
		}

		j = i + 1
		hare = numbers[j]
		for tortoise != hare {
			j += 1
			hare = numbers[j]
			lam += 1
		}
	}
	return numbers[mu : mu+lam]
}

func calculate_sad_number(power int, number string) {
	// TODO: there should be a better way than creating a fix sized array
	numberList := make([]int, 100*power)
	for i := 0; i < 100*power; i++ {
		n := power_each_digit(power, number)
		number = strconv.Itoa(n)
		numberList[i] = n
	}
	cycle := find_cycle(numberList)
	ncycle := len(cycle)
	for i, n := range cycle {
		fmt.Printf("%d", n)
		if i < ncycle-1 {
			fmt.Print(", ")
		} else {
			fmt.Println()
		}
	}
}

func main() {
	var power int
	var readPower bool = false

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Text()

		num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}

		if readPower {
			calculate_sad_number(power, s)
		} else {
			power = num
			readPower = true
		}
	}
}
