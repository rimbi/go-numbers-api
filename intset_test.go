package main

import (
	"reflect"
	"testing"
)

func TestNewlyAllocatedIntSetIsEmpty(t *testing.T) {
	set := NewIntSet()
	if len(set) != 0 {
		t.Errorf("Newly allocated set is not empty")
	}
}

func TestIntSetDoesNotContainDuplicates(t *testing.T) {
	set := NewIntSet()
	set.Add([]int{1, 2, 2})
	if len(set) != 2 {
		t.Errorf("Intset contains duplicate entries")
	}
}

func TestIntSetCanExportItsContentAsSortedIntSlice(t *testing.T) {
	set := NewIntSet()
	set.Add([]int{1, 7, 2, 6, 11})
	if !reflect.DeepEqual(set.AsSortedSlice(), []int{1, 2, 6, 7, 11}) {
		t.Errorf("Intset does not return its content as sorted int slice")
	}
}
