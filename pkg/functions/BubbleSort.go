package functions

func BubbleSort(sli []int) {
	len := len(sli)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-1-i; j++ {
			if sli[j] > sli[j+1] {
				swap(sli, j)
			}
		}
	}
}

func swap(sli []int, pos int) {
	sli[pos] = sli[pos] ^ sli[pos+1]
	sli[pos+1] = sli[pos] ^ sli[pos+1]
	sli[pos] = sli[pos] ^ sli[pos+1]
}
