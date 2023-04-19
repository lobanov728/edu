func myRand(N int) []int {
	list := make([]int, N)
	added := make(map[int]struct{}, N)
	i := 0
	for {
		x := rand.Int()
		_, ok := added[x]
		if !ok {
			list[i] = x
			i++
			added[x] = struct{}{}
		}

		if i == N {
			break
		}
	}

	return list
}