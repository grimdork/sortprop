package sortprop

// Map-based implementations that avoid sorting and preserve original order.

// UniqueKeysMap returns a slice with only one instance of each unique key.
// If keeplast is true, the last occurrence is kept; otherwise the first occurrence is kept.
// The input is not mutated and the original order is preserved (first-seen order or last-seen order).
func UniqueKeysMap(kp KeyProperties, keeplast bool) KeyProperties {
	if len(kp) == 0 {
		return KeyProperties{}
	}
	if keeplast {
		// keep last: iterate and record index of last occurrence
		last := make(map[string]Property, len(kp))
		for _, p := range kp {
			last[p.Key] = p
		}
		res := make(KeyProperties, 0, len(last))
		// preserve original order of keys by iterating original and adding when seeing key first time in result
		seen := make(map[string]bool, len(last))
		for _, p := range kp {
			if seen[p.Key] {
				continue
			}
			res = append(res, last[p.Key])
			seen[p.Key] = true
		}
		return res
	}
	// keep first: iterate and add first occurrence
	seen := make(map[string]bool, len(kp))
	res := make(KeyProperties, 0, len(kp))
	for _, p := range kp {
		if seen[p.Key] {
			continue
		}
		seen[p.Key] = true
		res = append(res, p)
	}
	return res
}

// UniqueValuesMap returns a slice with only one instance of each unique value.
// If keeplast is true, the last occurrence is kept; otherwise the first occurrence is kept.
// The input is not mutated and original order is preserved.
func UniqueValuesMap(vp ValueProperties, keeplast bool) ValueProperties {
	if len(vp) == 0 {
		return ValueProperties{}
	}
	if keeplast {
		last := make(map[string]Property, len(vp))
		for _, p := range vp {
			last[p.Value] = p
		}
		res := make(ValueProperties, 0, len(last))
		seen := make(map[string]bool, len(last))
		for _, p := range vp {
			if seen[p.Value] {
				continue
			}
			res = append(res, last[p.Value])
			seen[p.Value] = true
		}
		return res
	}
	seen := make(map[string]bool, len(vp))
	res := make(ValueProperties, 0, len(vp))
	for _, p := range vp {
		if seen[p.Value] {
			continue
		}
		seen[p.Value] = true
		res = append(res, p)
	}
	return res
}
