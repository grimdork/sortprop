package sortprop

import (
	"sort"
)

// Hybrid implementations: dedupe with map (fast), then return a sorted result.

// UniqueKeysHybrid returns unique properties by Key, using a map to dedupe then sorting the result by Key.
// If keeplast is true, the last occurrence is kept; otherwise the first occurrence is kept.
func UniqueKeysHybrid(kp KeyProperties, keeplast bool) KeyProperties {
	if len(kp) == 0 {
		return KeyProperties{}
	}
	m := make(map[string]Property, len(kp))
	if keeplast {
		for _, p := range kp {
			m[p.Key] = p
		}
	} else {
		for _, p := range kp {
			if _, ok := m[p.Key]; !ok {
				m[p.Key] = p
			}
		}
	}
	res := make(KeyProperties, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	sort.Sort(res)
	return res
}

// UniqueValuesHybrid returns unique properties by Value, using a map to dedupe then sorting the result by Value.
func UniqueValuesHybrid(vp ValueProperties, keeplast bool) ValueProperties {
	if len(vp) == 0 {
		return ValueProperties{}
	}
	m := make(map[string]Property, len(vp))
	if keeplast {
		for _, p := range vp {
			m[p.Value] = p
		}
	} else {
		for _, p := range vp {
			if _, ok := m[p.Value]; !ok {
				m[p.Value] = p
			}
		}
	}
	res := make(ValueProperties, 0, len(m))
	for _, v := range m {
		res = append(res, v)
	}
	sort.Sort(res)
	return res
}
