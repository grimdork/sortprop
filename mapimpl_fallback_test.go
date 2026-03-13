package sortprop

import "testing"

// Test that lowering FallbackThreshold forces the map-based path to fall back to the sort-based Unique* implementations.
func TestUniqueKeysMap_FallbackToSort(t *testing.T) {
	orig := FallbackThreshold
	defer func() { FallbackThreshold = orig }()
	// force fallback by setting threshold to 0
	FallbackThreshold = 0

	kp := KeyProperties{
		{Key: "a", Value: "1"},
		{Key: "b", Value: "2"},
		{Key: "a", Value: "3"},
	}

	outMap := UniqueKeysMap(kp, false)
	outSort := UniqueKeys(kp, false)

	if len(outMap) != len(outSort) {
		t.Fatalf("fallback mismatch: len(map)=%d len(sort)=%d", len(outMap), len(outSort))
	}
	for i := range outMap {
		if outMap[i].Key != outSort[i].Key || outMap[i].Value != outSort[i].Value {
			t.Fatalf("fallback mismatch at %d: map=%v sort=%v", i, outMap[i], outSort[i])
		}
	}
}

func TestUniqueValuesMap_FallbackToSort(t *testing.T) {
	orig := FallbackThreshold
	defer func() { FallbackThreshold = orig }()
	FallbackThreshold = 0

	vp := ValueProperties{
		{Key: "x", Value: "v1"},
		{Key: "y", Value: "v2"},
		{Key: "z", Value: "v1"},
	}

	outMap := UniqueValuesMap(vp, false)
	outSort := UniqueValues(vp, false)

	if len(outMap) != len(outSort) {
		t.Fatalf("fallback mismatch: len(map)=%d len(sort)=%d", len(outMap), len(outSort))
	}
	for i := range outMap {
		if outMap[i].Key != outSort[i].Key || outMap[i].Value != outSort[i].Value {
			t.Fatalf("fallback mismatch at %d: map=%v sort=%v", i, outMap[i], outSort[i])
		}
	}
}
