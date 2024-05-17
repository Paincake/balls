package main

import "testing"

func TestMain_CanSort(t *testing.T) {
	failed := 0
	passed := 0

	expected := true
	data := [][]int{{1, 2}, {2, 1}}
	helper(data, expected, &failed, &passed, t)

	expected = false
	data = [][]int{{30, 10, 20}, {1, 1, 1}, {0, 0, 10}}
	helper(data, expected, &failed, &passed, t)

	expected = false
	data = [][]int{{30, 10, 20}, {1, 1, 1}, {0, 0, 10}}
	helper(data, expected, &failed, &passed, t)

	expected = false
	data = [][]int{{100, 100}, {200, 200}}
	helper(data, expected, &failed, &passed, t)

	expected = false
	data = [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	helper(data, expected, &failed, &passed, t)

	expected = true
	data = [][]int{{0}, {0}}
	helper(data, expected, &failed, &passed, t)

	expected = true
	data = [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	helper(data, expected, &failed, &passed, t)

	if failed > 0 {
		t.Fatalf("\nPassed: %d Failed: %d", passed, failed)
	}

	t.Logf("\nPassed: %d Failed: %d", passed, failed)
}

func helper(data [][]int, expected bool, failed *int, passed *int, t *testing.T) {
	actual := canSortContainers(len(data), data)
	if actual != expected {
		*failed++
		t.Logf("\nTestcase: %v\nExpected: %t Got: %t\n", data, expected, !expected)
	} else {
		*passed++
	}
}
