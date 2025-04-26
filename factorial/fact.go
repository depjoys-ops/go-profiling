package factorial

const (
	maxSliceSize = 50000
)

func Recursive(n int) int {
	if n == 0 {
		return 1
	}
	return n * Recursive(n-1)
}

func Dynamic(n int) int {
	memo := make([]int, n+1)
	memo[0] = 1
	for i := 1; i <= n; i++ {
		memo[i] = i * memo[i-1]
	}
	return memo[n]
}

func Calculate() []int{
	var slice []int
	for i := 0; i < maxSliceSize; i++ {
		slice = append(slice, Dynamic(i))
	}
	return slice
}


