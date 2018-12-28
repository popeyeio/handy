package handy

// Abs completes abs for int64.
// Please see https://wuyin.io/2018/02/07/optimized-abs-func-for-int64-in-Go/.
func Abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}
