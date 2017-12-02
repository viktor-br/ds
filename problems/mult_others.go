package problems

// MultOthers calculate multiplication of numbers except in indx without division.
func MultOthers(a []int, indx int) (mult int) {
	var tree [][]int
	mult = 1
	tree = append(tree, a)
	tree, _, _ = createMultTree(tree, a)
	l := len(tree) - 1
	for i := 0; i < l; i++ {
		mult = mult * tree[i][indx^1]
		indx = indx >> 1
	}

	return mult
}

func createMultTree(tree [][]int, a []int) (t [][]int, mult int, level int) {
	l := len(a)
	if l == 2 {
		level = 1
		mult = a[0] * a[1]
		if len(tree) <= level {
			tree = append(tree, []int{})
		}
		tree[level] = append(tree[level], mult)
		return tree, mult, level
	}

	if l == 1 {
		level = 1
		mult = a[0]
		if len(tree) <= level {
			tree = append(tree, []int{})
		}
		tree[level] = append(tree[level], mult)
		return tree, mult, level
	}
	m := l / 2
	tree, mult1, _ := createMultTree(tree, a[0:m])
	tree, mult2, level := createMultTree(tree, a[m:])
	mult = mult1 * mult2

	level = level + 1
	if len(tree) <= level {
		tree = append(tree, []int{})
	}

	tree[level] = append(tree[level], mult)

	return tree, mult, level
}
