package sorter

type Sorter interface {
	QuickSorter
	BubbleSorter
	InsertionSorter
}

type QuickSorter interface{}

type BubbleSorter interface{}

type InsertionSorter interface{}
