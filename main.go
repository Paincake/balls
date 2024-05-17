package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func canSortContainers(n int, containers [][]int) bool {
	totalBalls := make([]int, n)
	for _, container := range containers {
		for color, count := range container {
			totalBalls[color] += count
		}
	}

	containerSums := make([]int, n)
	for i, container := range containers {
		for _, count := range container {
			containerSums[i] += count
		}
	}

	sort.Ints(totalBalls)
	sort.Ints(containerSums)

	for i := 0; i < n; i++ {
		if totalBalls[i] != containerSums[i] {
			return false
		}
	}
	return true
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	containers := make([][]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		parts := strings.Fields(line)
		containers[i] = make([]int, n)
		for j := 0; j < n; j++ {
			containers[i][j], _ = strconv.Atoi(parts[j])
		}
	}

	if canSortContainers(n, containers) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
