package algos

// Bsearch finds n inside the hs and returns its index if found, otherwise -1
func Bsearch(n int, hs []int) int {
	lo := 0
	hi := len(hs)
	idx := -1

	for lo < hi {
		mid := lo + (hi-lo)/2
		val := hs[mid]
		if val == n {
			idx = mid
			break
		} else if n < val {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return idx
}
