package main

import "fmt"

const N = 1010

func main() {
	n, m := 0, 0
	fmt.Scanln(&n, &m)
	textList := make([]string, 0)
	for i := 0; i < n; i++ {
		text := ""
		fmt.Scanln(&text)
		textList = append(textList, text)
	}
	for i := 0; i < m; i++ {
		ques := ""
		limit := 0
		fmt.Scanln(&ques, &limit)
		count := 0
		for i := 0; i < n; i++ {
			res := minDistance(textList[i], ques, len(textList[i]), len(ques))
			if res <= limit {
				count++
			}
		}
		fmt.Println(count)
	}

}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func minDistance(text1 string, text2 string, n int, m int) int {
	var dp [N][N]int //dp represent the set of all operation of modify the i character of text1  and let text1 equals text2
	text1 = fmt.Sprint(" ", text1)
	text2 = fmt.Sprint(" ", text2)
	for i := 0; i <= n; i++ {
		dp[i][0] = i //delete all the charaters of text1,  null == null
	}
	for i := 0; i <= m; i++ {
		dp[0][i] = i //add the character,  modify text1 , and let it equals text2
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1) // delete the i character of text1, or delete the j character of text2.
			if text1[i] == text2[j] {
				dp[i][j] = min(dp[i-1][j-1], dp[i][j]) // modif
			} else {
				dp[i][j] = min(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}
	return dp[n][m]
}
