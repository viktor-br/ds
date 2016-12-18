package problems

func BestStocks(a []int) (i, j int) {
	if len(a) == 1 {
		return 0, 0
	}
	if len(a) == 2 {
		return 0, 1
	}
	b := len(a) / 2
	i1, j1 := BestStocks(a[0:b])
	i2, j2 := BestStocks(a[b:])

	// TODO check, maybe we need only 3 of next variants.
	max := a[j1] - a[i1] // max1 - min1
	i, j = i1, j1

	v2 := a[b+j2] - a[b+i2] // max2 - min2
	if v2 > max {
		max = v2
		i, j = b+i2, b+j2
	}
	v3 := a[b+j2] - a[j1] // min2 - max1
	if v3 > max {
		max = v3
		i, j = j1, b+j2
	}
	v4 := a[b+i2] - a[i1] // min2 - min1
	if v4 > max {
		max = v4
		i, j = i1, b+i2
	}
	v5 := a[b+j2] - a[i1] // max2 - min1
	if v5 > max {
		i, j = i1, b+j2
	}
	v6 := a[b+j2] - a[j1] // max2 - max1
	if v6 > max {
		max = v6
		i, j = j1, b+j2
	}

	return i, j
}

func BestStocksBruteForce(a []int) (first, second int) {
	aLen := len(a)
	// Wrong input
	if aLen <= 1 {
		return 0, 0
	}
	max := a[1] - a[0]
	first, second = 0, 1
	for i := 0; i < aLen-1; i++ {
		for j := i + 1; j < aLen; j++ {
			v := a[j] - a[i]
			if v > max {
				first, second = i, j
				max = v
			}
		}
	}

	return
}
