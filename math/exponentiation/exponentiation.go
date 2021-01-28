package main

import "fmt"

func main() {
	n := 0
	fmt.Scanln(&n)
	for ; n > 0; n-- {
		a, k, p := 1, 1, 1
		fmt.Scanln(&a, &k, &p)
		fmt.Println(qmi(a, k, p))
	}
}

//快速幂， 反复平方法
//预处理出一些结果

func qmi(a, k, p int) int64 {
	var res int64
	res = 1
	for k != 0 {
		if k&1 == 1 {
			res = res * int64(a) % int64(p)
		}
		k >>= 1
		a = a * a % p
	}

	return res
}
