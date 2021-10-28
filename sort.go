package sorter

type Sorter interface {
	Sort(arr []int)
}

type Selection struct{}

func (ss *Selection) Sort(arr []int) {
	for j := 0; j < len(arr); j++ {
		min := j
		for i := j + 1; i < len(arr); i++ {
			if arr[min] > arr[i] {
				min = i
			}
		}
		arr[j], arr[min] = arr[min], arr[j]
	}
}

type Insertion struct{}

func (is *Insertion) Sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		current := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > current {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = current
	}
}

type Bubble struct{}

func (bs *Bubble) Sort(arr []int) {
	for j := range arr {
		isSorted := true
		for i := 1; i < len(arr)-j; i++ {
			if arr[i] < arr[i-1] {
				isSorted = false
				arr[i], arr[i-1] = arr[i-1], arr[i]
			}
		}
		if isSorted {
			return
		}
	}
}

type Quick struct{}

func (qs *Quick) Sort(arr []int) {
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

type Counting struct {
}

func (cs *Counting) Sort(arr []int) {
	min := arr[0]
	max := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
		if arr[i] < min {
			min = arr[i]
		}
	}
	bucket := make([]int, max-min+1)

	for i := 0; i < len(arr); i++ {
		bucket[arr[i]-min]++
	}
	arr = arr[:0]

	for i := 0; i < len(bucket); i++ {
		cnt := bucket[i]
		for j := 0; j < cnt; j++ {
			arr = append(arr, i+min)
		}
	}
}

type Merge struct {
}

//todo: change me
func (ms *Merge) Sort(arr []int) []int {
	var num = len(arr)

	if num == 1 {
		return arr
	}

	middle := num / 2
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = arr[i]
		} else {
			right[i-middle] = arr[i]
		}
	}

	return merge(ms.Sort(left), ms.Sort(right))
}

func merge(left, right []int) (merged []int) {

	merged = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			merged[i] = left[0]
			left = left[1:]
		} else {
			merged[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		merged[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		merged[i] = right[j]
		i++
	}

	return
}
