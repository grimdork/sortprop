package sortprop

import "sync"

// Map-based implementations that avoid sorting and preserve original order.
// Optimized to reduce allocations by pre-sizing maps and result slices and reusing map objects via sync.Pool.

var (
	propMapPool = sync.Pool{New: func() interface{} { return make(map[string]Property) }}
	seenMapPool = sync.Pool{New: func() interface{} { return make(map[string]struct{}) }}
	// result slice pools to reuse allocations for common sizes
	kpPool = sync.Pool{New: func() interface{} { return make(KeyProperties, 0, 0) }}
	vpPool = sync.Pool{New: func() interface{} { return make(ValueProperties, 0, 0) }}
	// avoid returning huge backing arrays from pools; cap threshold
	poolCapThreshold = 16384 // elements
)

func getPropMap(capacity int) map[string]Property {
	if capacity > 0 {
		return make(map[string]Property, capacity)
	}
	m := propMapPool.Get().(map[string]Property)
	for k := range m {
		delete(m, k)
	}
	return m
}

func putPropMap(m map[string]Property) {
	for k := range m {
		delete(m, k)
	}
	propMapPool.Put(m)
}

func getSeenMap(capacity int) map[string]struct{} {
	if capacity > 0 {
		return make(map[string]struct{}, capacity)
	}
	m := seenMapPool.Get().(map[string]struct{})
	for k := range m {
		delete(m, k)
	}
	return m
}

func putSeenMap(m map[string]struct{}) {
	for k := range m {
		delete(m, k)
	}
	seenMapPool.Put(m)
}

// UniqueKeysMap returns a slice with only one instance of each unique key.
// If keeplast is true, the last occurrence is kept; otherwise the first occurrence is kept.
// The input is not mutated and the original order is preserved (first-seen order or last-seen order).
func UniqueKeysMap(kp KeyProperties, keeplast bool) KeyProperties {
	if len(kp) == 0 {
		return KeyProperties{}
	}
	if keeplast {
		// keep last: record last occurrence in a map
		last := getPropMap(len(kp))
		for _, p := range kp {
			last[p.Key] = p
		}
			// get a result slice from pool and ensure capacity
		res := kpPool.Get().(KeyProperties)
		res = res[:0]
		if cap(res) < len(last) {
			res = make(KeyProperties, 0, len(last))
		}
		seen := getSeenMap(len(last))
		for _, p := range kp {
			if _, ok := seen[p.Key]; ok {
				continue
			}
			res = append(res, last[p.Key])
			seen[p.Key] = struct{}{}
		}
		putPropMap(last)
		putSeenMap(seen)
		// if result backing cap is reasonable, return it directly; otherwise copy to shrink-to-fit
		if cap(res) <= poolCapThreshold {
			out := res[:len(res)]
			kpPool.Put(res[:0])
			return out
		}
		out := make(KeyProperties, len(res))
		copy(out, res)
		kpPool.Put(res[:0])
		return out
	}
	// keep first: iterate and add first occurrence
	seen := getSeenMap(len(kp))
	res := kpPool.Get().(KeyProperties)
	res = res[:0]
	if cap(res) < len(kp) {
		res = make(KeyProperties, 0, len(kp))
	}
	for _, p := range kp {
		if _, ok := seen[p.Key]; ok {
			continue
		}
		seen[p.Key] = struct{}{}
		res = append(res, p)
	}
	putSeenMap(seen)
	out := make(KeyProperties, len(res))
	copy(out, res)
	kpPool.Put(res)
	return out
}

// UniqueValuesMap returns a slice with only one instance of each unique value.
// If keeplast is true, the last occurrence is kept; otherwise the first occurrence is kept.
// The input is not mutated and original order is preserved.
func UniqueValuesMap(vp ValueProperties, keeplast bool) ValueProperties {
	if len(vp) == 0 {
		return ValueProperties{}
	}
	if keeplast {
		last := getPropMap(len(vp))
		for _, p := range vp {
			last[p.Value] = p
		}
			res := vpPool.Get().(ValueProperties)
		res = res[:0]
		if cap(res) < len(last) {
			res = make(ValueProperties, 0, len(last))
		}
		seen := getSeenMap(len(last))
		for _, p := range vp {
			if _, ok := seen[p.Value]; ok {
				continue
			}
			res = append(res, last[p.Value])
			seen[p.Value] = struct{}{}
		}
		putPropMap(last)
		putSeenMap(seen)
		if cap(res) <= poolCapThreshold {
			out := res[:len(res)]
			vpPool.Put(res[:0])
			return out
		}
		out := make(ValueProperties, len(res))
		copy(out, res)
		vpPool.Put(res[:0])
		return out
	}
	seen := getSeenMap(len(vp))
	res := vpPool.Get().(ValueProperties)
	res = res[:0]
	if cap(res) < len(vp) {
		res = make(ValueProperties, 0, len(vp))
	}
	for _, p := range vp {
		if _, ok := seen[p.Value]; ok {
			continue
		}
		seen[p.Value] = struct{}{}
		res = append(res, p)
	}
	putSeenMap(seen)
	out := make(ValueProperties, len(res))
	copy(out, res)
	vpPool.Put(res)
	return out
}
