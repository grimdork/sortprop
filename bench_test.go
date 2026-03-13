package sortprop

import (
	"fmt"
	"testing"
)

func genKP(n int) KeyProperties {
	kp := make(KeyProperties, 0, n)
	for i := 0; i < n; i++ {
		kp = append(kp, Property{Key: fmt.Sprintf("k%06d", i%100), Value: fmt.Sprintf("v%06d", i)})
	}
	return kp
}

func genVP(n int) ValueProperties {
	vp := make(ValueProperties, 0, n)
	for i := 0; i < n; i++ {
		vp = append(vp, Property{Key: fmt.Sprintf("k%06d", i), Value: fmt.Sprintf("v%06d", i%100)})
	}
	return vp
}

func BenchmarkUniqueKeys_Copy(b *testing.B) {
	for _, n := range []int{1000, 10000} {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			kp := genKP(n)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = UniqueKeys(kp, false)
			}
		})
	}
}

func BenchmarkUniqueValues_Copy(b *testing.B) {
	for _, n := range []int{1000, 10000} {
		b.Run(fmt.Sprintf("n=%d", n), func(b *testing.B) {
			vp := genVP(n)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = UniqueValues(vp, false)
			}
		})
	}
}
