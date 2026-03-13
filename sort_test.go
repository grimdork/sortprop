package sortprop

import (
	"sort"
	"testing"
)

func TestSortByKey(t *testing.T) {
	kp := KeyProperties{
		{Key: "b", Value: "2"},
		{Key: "a", Value: "1"},
		{Key: "c", Value: "3"},
	}
	sort.Sort(kp)
	if kp[0].Key != "a" || kp[1].Key != "b" || kp[2].Key != "c" {
		t.Fatalf("unexpected key order: %v", kp)
	}
}

func TestSortByValue(t *testing.T) {
	vp := ValueProperties{
		{Key: "k1", Value: "z"},
		{Key: "k2", Value: "a"},
		{Key: "k3", Value: "m"},
	}
	sort.Sort(vp)
	if vp[0].Value != "a" || vp[1].Value != "m" || vp[2].Value != "z" {
		t.Fatalf("unexpected value order: %v", vp)
	}
}

func TestSortEmptyAndNil(t *testing.T) {
	var kp KeyProperties
	sort.Sort(kp) // should not panic
	if len(kp) != 0 {
		t.Fatalf("expected empty slice, got %v", kp)
	}
	var vp ValueProperties
	sort.Sort(vp)
	if len(vp) != 0 {
		t.Fatalf("expected empty slice, got %v", vp)
	}
}

func TestSortDuplicates(t *testing.T) {
	kp := KeyProperties{
		{Key: "a", Value: "1"},
		{Key: "a", Value: "2"},
		{Key: "b", Value: "3"},
	}
	sort.Sort(kp)
	// keys equal for first two; ensure they are adjacent and b is last
	if kp[0].Key != "a" || kp[1].Key != "a" || kp[2].Key != "b" {
		t.Fatalf("unexpected order with duplicates: %v", kp)
	}
}
