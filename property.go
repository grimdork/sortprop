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
// The list will be sorted by Key.
func UniqueKeys(kp KeyProperties, keeplast bool) KeyProperties {
	var list KeyProperties
	sort.Sort(kp)
	for i := 0; i < len(kp); i++ {
		if i == 0 {
			list = append(list, kp[i])
			continue
		}

		if kp[i].Key == kp[i-1].Key {
			if keeplast {
				list[len(list)-1] = kp[i]
			}
		} else {
			list = append(list, kp[i])
		}
	}
	return list
}

// UniqueValues returns a slice with only one instance of each unique value-property.
// The list will be sorted by Value.
func UniqueValues(vp ValueProperties, keeplast bool) ValueProperties {
	var list ValueProperties
	sort.Sort(vp)
	for i := 0; i < len(vp); i++ {
		if i == 0 {
			list = append(list, vp[i])
			continue
		}

		if vp[i].Value == vp[i-1].Value {
			if keeplast {
				list[len(list)-1] = vp[i]
			}
		} else {
			list = append(list, vp[i])
		}
	}
	return list
}
