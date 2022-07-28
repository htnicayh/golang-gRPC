package main

func single(nums []int) int {
	n := 0
	data := make(map[int]int)

	for _, val := range nums {
		data[val] = data[val] + 1
	}

	for key, val := range data {
		if val == 1 {
			n = key
			break
		}
	}
	return n
}

func main() {
	single([]int{1, 2, 2, 3, 3})
}
