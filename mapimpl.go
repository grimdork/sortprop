package sortprop

// Map-based implementations that avoid sorting and preserve original order.
// Optimized to reduce allocations by pre-sizing maps and result slices.

// UniqueKeysMap returns a slice with only one instance of each unique key.
// If keeplast is true, the last occurrence is kept; otherwise the first occurrence is kept.
// The input is not mutated and the original order is preserved (first-seen order or last-seen order).
func UniqueKeysMap(kp KeyProperties, keeplast bool) KeyProperties {
	if len(kp) == 0 {
		return KeyProperties{}
	}
	if keeplast {
		// keep last: record last occurrence in a map
		last := make(map[string]Property, len(kp))
		for _, p := range kp {
			last[p.Key] = p
		}
		// preallocate result with exact size
		res := make(KeyProperties, 0, len(last))
		seen := make(map[string]struct{}, len(last))
		for _, p := range kp {
			if _, ok := seen[p.Key]; ok {
				continue
			}
			res = append(res, last[p.Key])
			seen[p.Key] = struct{}{}
		}
		return res
	}
	// keep first: iterate and add first occurrence
	seen := make(map[string]struct{}, len(kp))
	res := make(KeyProperties, 0, len(kp))
	for _, p := range kp {
		if _, ok := seen[p.Key]; ok {
			continue
		}
		seen[p.Key] = struct{}{}
		res = append(res, p)
	}
	// shrink to fit exact unique count
	if len(res) != cap(res) {
		out := make(KeyProperties, len(res))
		copy(out, res)
		return out
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
		seen := make(map[string]struct{}, len(last))
		for _, p := range vp {
			if _, ok := seen[p.Value]; ok {
				continue
			}
			res = append(res, last[p.Value])
			seen[p.Value] = struct{}{}
		}
		return res
	}
	seen := make(map[string]struct{}, len(vp))
	res := make(ValueProperties, 0, len(vp))
	for _, p := range vp {
		if _, ok := seen[p.Value]; ok {
			continue
		}
		seen[p.Value] = struct{}{}
		res = append(res, p)
	}
	if len(res) != cap(res) {
		out := make(ValueProperties, len(res))
		copy(out, res)
		return out
	}
	return res
}
