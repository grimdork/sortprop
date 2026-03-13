package sortprop

import (
	"fmt"
	"testing"
)

func genKP(n int, uniqueRatio float64) KeyProperties {
	kp := make(KeyProperties, 0, n)
	unique := int(float64(n) * uniqueRatio)
	if unique < 1 {
		unique = 1
	}
	for i := 0; i < n; i++ {
		kp = append(kp, Property{Key: fmt.Sprintf("k%06d", i%unique), Value: fmt.Sprintf("v%06d", i)})
	}
	return kp
}

func genVP(n int, uniqueRatio float64) ValueProperties {
	vp := make(ValueProperties, 0, n)
	unique := int(float64(n) * uniqueRatio)
	if unique < 1 {
		unique = 1
	}
	for i := 0; i < n; i++ {
		vp = append(vp, Property{Key: fmt.Sprintf("k%06d", i), Value: fmt.Sprintf("v%06d", i%unique)})
	}
	return vp
}

func BenchmarkUniqueKeys_All(b *testing.B) {
	for _, n := range []int{1000, 10000, 100000} {
		for _, ur := range []float64{0.01, 0.1, 0.5, 1.0} { // unique ratios
			label := fmt.Sprintf("n=%d/ur=%.2f", n, ur)
			b.Run(label, func(b *testing.B) {
				kp := genKP(n, ur)
				b.Run("Sort_first", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = UniqueKeys(kp, false)
					}
				})
				b.Run("Map_first", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = UniqueKeysMap(kp, false)
					}
				})
				b.Run("Hybrid_first", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = UniqueKeysHybrid(kp, false)
					}
				})
				b.Run("Sort_last", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = UniqueKeys(kp, true)
					}
				})
				b.Run("Map_last", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = UniqueKeysMap(kp, true)
					}
				})
				b.Run("Hybrid_last", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = UniqueKeysHybrid(kp, true)
					}
				})
			})
		}
	}
}

func BenchmarkUniqueValues_All(b *testing.B) {
	for _, n := range []int{1000, 10000, 100000} {
		for _, ur := range []float64{0.01, 0.1, 0.5, 1.0} {
			label := fmt.Sprintf("n=%d/ur=%.2f", n, ur)
			b.Run(label, func(b *testing.B) {
				vp := genVP(n, ur)
				b.Run("Sort_first", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = UniqueValues(vp, false)
					}
				})
				b.Run("Map_first", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = UniqueValuesMap(vp, false)
					}
				})
				b.Run("Hybrid_first", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = UniqueValuesHybrid(vp, false)
					}
				})
				b.Run("Sort_last", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = UniqueValues(vp, true)
					}
				})
				b.Run("Map_last", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = UniqueValuesMap(vp, true)
					}
				})
				b.Run("Hybrid_last", func(b *testing.B) {
					for i := 0; i < b.N; i++ {
						_ = UniqueValuesHybrid(vp, true)
					}
				})
			})
		}
	}
}
