package main

import (
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	tests := []struct {
		head *ListNode
		want *ListNode
	}{
		{
			head: newList(1, 1, 2, 3, 3, 3),
			want: newList(1, 2, 3),
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := deleteDuplicates(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
