package main

import "fmt"

func main() {
	n := 0
	fmt.Scanln(&n)
	res := 0
	for i := 2; i <= n; i++ {
		//fmt.Println("i:", i)
		ok := checkPrime(i)
		//fmt.Println("ok:", ok)
		if ok {
			res++
		}
	}
	fmt.Println(res)
}

//试除法判定素数 O(sqrt(n))
func checkPrime(x int) bool {
	//for i := 2; i*i <= x; i++ {
	// x 为整数最大值时， i * i 容易溢出 ,推荐使用  i <= x/i
	for i := 2; i <= x/i; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

//分解质因数
func divide(n int) {
	for i := 2; i <= n/i; i++ {
		res := 0
		if n/i == 0 {
			for n/i == 0 {
				n = n / i
				res++
			}
		}
		fmt.Println(i, res)
	}

	// n 只有一个 大于 sqrt(n) 的质因数，除到最后， 如果n> 1, 这个数就是大于 sqrt(n) 的质因数
	if n > 1 {
		fmt.Println(n)
	}
}

const N = 100010

var st [N]bool
var primes [N]int
var cnt int

//筛质数方法：朴素筛法 O(nlogn), 调和级数
func filterPrimes(n int) {
	for i := 2; i <= n; i++ {
		if !st[i] {
			primes[cnt] = i
			cnt++
		}
		for j := i + i; j <= n; j += i {
			st[j] = true
		}
	}
}

//埃式筛法 O(nloglogn), 不用遍历所有数， 把遇到的质数的倍数都筛掉，最后剩下的就是质数。
func afilterPrimes(n int) {
	for i := 2; i <= n; i++ {
		if !st[i] {
			primes[cnt] = i
			cnt++
			for j := i + i; i <= n; j += i {
				st[j] = true
			}
		}
	}
}

//线性筛法
func linearFiterPrimes(n int) {
	for i := 2; i <= n; i++ {
		if !st[i] {
			primes[cnt] = i
			cnt++
		}
		//核心思想， 所有合数都会被自己最小的质数筛掉
		for j := 0; primes[j] <= n/i; j++ {
			st[primes[j]*i] = true
			if i%primes[j] == 0 {
				break
			}
		}
	}
}
