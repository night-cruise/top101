package recursionbacktracking

// 回溯
//
// 在当前行选一个位置放置棋子，然后递归地在剩下的行中放置棋子，
// 同时使用四个访问数组记录已经放置棋子，当在一个位置放置棋子时首先通过访问数组判断这个位置是否跟已经放置的棋子冲突。
func Nqueen(n int) int {
	ans := 0
	NHelper(n, 0, make([]bool, n), make([]bool, n), make([]bool, 2*n-1), make([]bool, 2*n-1), &ans)
	return ans
}

func NHelper(n int, x int, rows, cols, loblis, roblis []bool, ans *int) {
	if x == n {
		*ans += 1
		return
	}
	for y := 0; y < n; y++ {
		if rows[x] || cols[y] || loblis[x+y] || roblis[n-1-y+x] {
			continue
		}
		rows[x], cols[y], loblis[x+y], roblis[n-1-y+x] = true, true, true, true
		NHelper(n, x+1, rows, cols, loblis, roblis, ans)
		rows[x], cols[y], loblis[x+y], roblis[n-1-y+x] = false, false, false, false
	}
}
