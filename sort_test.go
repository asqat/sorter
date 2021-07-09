package sorter

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestQuickSort_Sort(t *testing.T) {
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
			qs := &QuickSort{}
			qs.Sort(tt.args.arr)
			if !reflect.DeepEqual(tt.args.arr, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) {
				t.Errorf("cannot sort: %v", tt.args.arr)
			}
		})
	}
}
