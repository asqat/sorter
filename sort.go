package sorter

type Sorter interface {
	QuickSorter
	BubbleSorter
	InsertionSorter
}

type QuickSorter interface {
	Sort(arr []int)
}

type QuickSort struct{}

func (qs *QuickSort) Sort(arr []int) {
	n := len(arr)
	if n < 2 {
		return
	}

	left, right := 0, n-1

	pivot := n / 2

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}
	arr[right], arr[left] = arr[left], arr[right]

	qs.Sort(arr[:left])
	qs.Sort(arr[left+1:])
}

type BubbleSorter interface{}

type InsertionSorter interface{}
