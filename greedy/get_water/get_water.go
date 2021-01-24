package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*

913. 排队打水
   题目
   提交记录
   讨论
   题解
   视频讲解

有 n 个人排队到 1 个水龙头处打水，第 i 个人装满水桶所需的时间是 ti，请问如何安排他们的打水顺序才能使所有人的等待时间之和最小？

输入格式
第一行包含整数 n。

第二行包含 n 个整数，其中第 i 个整数表示第 i 个人装满水桶所花费的时间 ti。

输出格式
输出一个整数，表示最小的等待时间之和。

数据范围
1≤n≤105,
1≤ti≤104
输入样例：
7
3 6 1 4 2 5 7
输出样例：
56
难度：简单
时/空限制：1s / 64MB
总通过数：2187
总尝试数：3902
来源：模板题
算法标签
*/

/*
	f(x) = (n-1) *t1 + (n-2) *t2 + (n-3) * t3 + ... + 0 * tn
	t1, .. tn  按升序排列， f(x) 最小
*/

func main() {
	n := 0
	fmt.Scanln(&n)
	rd := bufio.NewReader(os.Stdin)
	intArr := readline(rd)
	sort.SliceStable(intArr, func(i, j int) bool {
		return intArr[i] < intArr[j]
	})

	res := 0
	for i := 0; i < n; i++ {
		res += intArr[i] * (n - i - 1)
	}
	fmt.Println(res)
}

func readline(rd *bufio.Reader) []int {
	line, _ := rd.ReadString('\n')
	params := strings.Fields(line)
	intArr := make([]int, 0)
	for _, elem := range params {
		intElem, _ := strconv.Atoi(elem)
		intArr = append(intArr, intElem)
	}
	return intArr
}
