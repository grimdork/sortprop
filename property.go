package sortprop

import "sort"

// KeyProperties is a slice of Property elements sorted by Key.
type KeyProperties []Property

// ValueProperties is a slice of Property elements sorted by Value.
type ValueProperties []Property

// Property is a key-value pair of strings.
type Property struct {
	Key   string
	Value string
}

// UniqueKeys returns a slice with only one instance of each unique-key property.
// If keeplast is true, the last element of the same key will be kept rather than the first.
// The returned list will be sorted by Key. The input slice is not mutated.
func UniqueKeys(kp KeyProperties, keeplast bool) KeyProperties {
	// work on a copy so the caller's slice is not mutated
	copyKP := append(KeyProperties(nil), kp...)
	sort.Sort(copyKP)

	list := make(KeyProperties, 0, len(copyKP))
	for i := 0; i < len(copyKP); i++ {
		if i == 0 {
			list = append(list, copyKP[i])
			continue
		}

		if copyKP[i].Key == copyKP[i-1].Key {
			if keeplast {
				list[len(list)-1] = copyKP[i]
			}
		} else {
			list = append(list, copyKP[i])
		}
	}
	return list
}

// UniqueValues returns a slice with only one instance of each unique value-property.
// If keeplast is true, the last element of the same value will be kept rather than the first.
// The returned list will be sorted by Value. The input slice is not mutated.
func UniqueValues(vp ValueProperties, keeplast bool) ValueProperties {
	// work on a copy so the caller's slice is not mutated
	copyVP := append(ValueProperties(nil), vp...)
	sort.Sort(copyVP)

	list := make(ValueProperties, 0, len(copyVP))
	for i := 0; i < len(copyVP); i++ {
		if i == 0 {
			list = append(list, copyVP[i])
			continue
		}

		if copyVP[i].Value == copyVP[i-1].Value {
			if keeplast {
				list[len(list)-1] = copyVP[i]
			}
		} else {
			list = append(list, copyVP[i])
		}
	}
	return list
}
