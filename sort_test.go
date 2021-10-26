package sorter

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestSelection_Sort(t *testing.T) {
	type args struct {
		arr []int
	}
	rand.Seed(time.Now().UnixNano())
	tests := []struct {
		name string
		args args
	}{
		{
			name: "sSort-1",
			args: args{
				arr: rand.Perm(10),
			},
		},
		{
			name: "sSort-2",
			args: args{
				arr: rand.Perm(10),
			},
		},
		{
			name: "sSort-3",
			args: args{
				arr: rand.Perm(10),
			},
		},
		{
			name: "sSort-4",
			args: args{
				arr: rand.Perm(10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ss := &Selection{}
			ss.Sort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
				t.Errorf("cannot sort: %v", tt.args.arr)
			}
		})
	}
}

func TestInsertion_Sort(t *testing.T) {
	type args struct {
		arr []int
	}
	rand.Seed(time.Now().UnixNano())
	tests := []struct {
		name string
		args args
	}{
		{
			name: "iSort-1",
			args: args{
				arr: rand.Perm(10),
			},
		},
		{
			name: "iSort-2",
			args: args{
				arr: rand.Perm(10),
			},
		},
		{
			name: "iSort-3",
			args: args{
				arr: rand.Perm(10),
			},
		},
		{
			name: "iSort-4",
			args: args{
				arr: rand.Perm(10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := &Insertion{}
			is.Sort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
				t.Errorf("cannot sort: %v", tt.args.arr)
			}
		})
	}
}

func TestBubble_Sort(t *testing.T) {
	type args struct {
		arr []int
	}
	rand.Seed(time.Now().UnixNano())
	tests := []struct {
		name string
		args args
	}{
		{
			name: "bSort-1",
			args: args{
				arr: rand.Perm(10),
			},
		},
		{
			name: "bSort-2",
			args: args{
				arr: rand.Perm(10),
			},
		},
		{
			name: "bSort-3",
			args: args{
				arr: rand.Perm(10),
			},
		},
		{
			name: "bSort-4",
			args: args{
				arr: rand.Perm(10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := &Bubble{}
			bs.Sort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
				t.Errorf("cannot sort: %v", tt.args.arr)
			}
		})
	}
}

func TestQuick_Sort(t *testing.T) {
	type args struct {
		arr []int
	}
	rand.Seed(time.Now().UnixNano())
	tests := []struct {
		name string
		args args
	}{
		{
			name: "qSort-1",
			args: args{
				arr: rand.Perm(10),
			},
		},
		{
			name: "qSort-2",
			args: args{
				arr: rand.Perm(10),
			},
		},
		{
			name: "qSort-3",
			args: args{
				arr: rand.Perm(10),
			},
		},
		{
			name: "qSort-4",
			args: args{
				arr: rand.Perm(10),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qs := &Quick{}
			qs.Sort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
				t.Errorf("cannot sort: %v", tt.args.arr)
			}
		})
	}
}

func makeRandIntArr(length int) []int {
	var arr []int
	for i := 0; i < length; i++ {
		arr = append(arr, rand.Intn(10000))
	}
	return arr
}

type testCase struct {
	job  func(b *testing.B)
	name string
}

func BenchmarkSorter(b *testing.B) {
	genArr := makeRandIntArr(1000000)

	var testCases []testCase

	sortAlgo := &testCase{}

	sortAlgo.name = "quickSort"
	sortAlgo.job = func(b *testing.B) {
		qs := &Quick{}
		raw := genArr
		for i := 0; i < b.N; i++ {
			qs.Sort(raw)
		}
	}
	testCases = append(testCases, *sortAlgo)

	sortAlgo.name = "insertionSort"
	sortAlgo.job = func(b *testing.B) {
		bs := &Insertion{}
		raw := genArr
		for i := 0; i < b.N; i++ {
			bs.Sort(raw)
		}
	}
	testCases = append(testCases, *sortAlgo)

	sortAlgo.name = "selectionSort"
	sortAlgo.job = func(b *testing.B) {
		bs := &Insertion{}
		raw := genArr
		for i := 0; i < b.N; i++ {
			bs.Sort(raw)
		}
	}
	testCases = append(testCases, *sortAlgo)

	sortAlgo.name = "bubbleSort"
	sortAlgo.job = func(b *testing.B) {
		bs := &Bubble{}
		raw := genArr
		for i := 0; i < b.N; i++ {
			bs.Sort(raw)
		}
	}
	testCases = append(testCases, *sortAlgo)

	for _, testCase := range testCases {
		b.Run(fmt.Sprintf("%s bench", testCase.name), testCase.job)
	}
}
