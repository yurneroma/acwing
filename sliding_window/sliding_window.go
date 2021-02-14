package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, k := 0, 0
	fmt.Scanln(&n, &k)

	//sanner read n int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	nums := make([]int, 0)
	for scanner.Scan() {
		if scanner.Text() != "" {
			val, _ := strconv.Atoi(scanner.Text())
			nums = append(nums, val)
		}
	}

	res := maxSlidingWindow(nums, k)

	for _, val := range res {
		fmt.Printf("%d ", val)
	}
	fmt.Println(" ")
}

func maxSlidingWindow(nums []int, k int) []int {
	mq := make([]int, 0)
	pushBack := func(val int) {
		for len(mq) > 0 && mq[len(mq)-1] < val {
			mq = mq[:len(mq)-1]
		}
		mq = append(mq, val)
	}

	for i := 0; i < k-1; i++ {
		pushBack(nums[i])
	}

	pop := func(val int) {
		if val == mq[0] {
			mq = mq[1:]
		}
	}

	res := make([]int, 0)
	for i := k - 1; i < len(nums); i++ {
		pushBack(nums[i])
		res = append(res, mq[0])
		pop(nums[i-k+1])
	}

	return res
}
