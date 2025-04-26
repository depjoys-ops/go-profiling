package factorial

const (
	maxSliceSize = 50000
)

//Recursive
func Recursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * Recursive(n-1)
}

//Dynamic
func Dynamic(n int) int {
	memo := make([]int, n+1)
	memo[0] = 1
	for i := 1; i <= n; i++ {
		memo[i] = i * memo[i-1]
	}
	return memo[n]
}

