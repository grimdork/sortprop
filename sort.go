package sortprop

// Implements sort.Interface for KeyProperties and ValueProperties.
// Note: these implementations sort in-place (they mutate the slice) and use the
// standard library's sort.Sort which is not stable. If callers need a stable
// ordering for equal keys/values, use sort.SliceStable with an appropriate
// tie-breaker or use the provided SortedByKey/SortedByValue helpers (if added).

//
// Sort by Key
//

// Len is the number of properties.
func (p KeyProperties) Len() int {
	return len(p)
}

// Less reports whether the property with index a should sort before the property with index b.
// Comparison is lexicographic on the Key string.
func (p KeyProperties) Less(a, b int) bool {
	return p[a].Key < p[b].Key
}

// Swap the properties with indices a and b.
func (p KeyProperties) Swap(a, b int) {
	p[a], p[b] = p[b], p[a]
}

//
// Sort by Value
//

// Len is the number of properties.
func (p ValueProperties) Len() int {
	return len(p)
}

// Less reports whether the property with index a should sort before the property with index b.
// Comparison is lexicographic on the Value string.
func (p ValueProperties) Less(a, b int) bool {
	return p[a].Value < p[b].Value
}

// Swap the properties with indices a and b.
func (p ValueProperties) Swap(a, b int) {
	p[a], p[b] = p[b], p[a]
}
