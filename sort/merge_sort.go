package sort

func MergeSort(a []int) []int {
	inputLen := len(a)
	if inputLen == 1 {
		return a
	}
	m := inputLen / 2
	result := make([]int, inputLen)

	part1 := MergeSort(a[0:m])
	part2 := MergeSort(a[m:])

	i := 0
	j := 0

	for k := 0; k < inputLen; k++ {
		if part1[i] <= part2[j] {
			result[k] = part1[i]
			i++
			if i >= len(part1) {
				copy(result[k+1:], part2[j:])
				break
			}
		} else {
			result[k] = part2[j]
			j++
			if j >= len(part2) {
				copy(result[k+1:], part1[i:])
				break
			}
		}
	}

	return result
}
