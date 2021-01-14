package sortprop

//
// Sort by Key
//

// Len is the number of properties.
func (p KeyProperties) Len() int {
	return len(p)
}

// Less reports whether the property with index a should sort before the property with index b.
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
func (p ValueProperties) Less(a, b int) bool {
	return p[a].Value < p[b].Value
}

// Swap the properties with indices a and b.
func (p ValueProperties) Swap(a, b int) {
	p[a], p[b] = p[b], p[a]
}
