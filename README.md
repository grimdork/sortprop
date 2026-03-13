# sortprop
Package for sortable key-value properties, like HTTP headers, with functions to remove duplicates.

## Usage

### Key-sortable slices

```go
	…
	kp := sortprop.KeyProperties{
		sortprop.Property{"3", "third"},
		sortprop.Property{"1", "first"},
		sortprop.Property{"2", "second"},
	}
	sort.Sort(kp)
	…
```

### Value-sortable slices

```go
	…
	vp := sortprop.ValueProperties{
		sortprop.Property{"3", "cccc"},
		sortprop.Property{"1", "aaaa"},
		sortprop.Property{"2", "bbbb"},
	}
	sort.Sort(vp)
	…
```

### Appending properties

KeyProperties and ValueProperties are just slices, so use append as usual:

```go
	list := sortprop.KeyValues{}
	p := sortprop.Property{"key", "value"}
	list = append(list, p)
```

### Removing duplicates

These two functions will remove duplicate keys and values:
```go
func UniqueKeys(kp KeyProperties, keeplast bool) KeyProperties
func UniqueValues(vp ValueProperties, keeplast bool) ValueProperties
```

They return a new slice with only one of each repeated key or value. If keeplast is false, the first of each duplicate
element is kept. If keeplast is true, the last one encountered is kept.

### Notes on behavior and performance

- UniqueKeys and UniqueValues do not mutate the input slice; they operate on a copy and return a newly allocated,
  sorted slice of unique properties.
- Complexity: O(n log n) due to sorting, with O(n) additional work for deduplication. The functions pre-allocate result
  slices to the input length to reduce allocations.
- If you need to preserve the original (insertion) order instead of returning a sorted result, consider using a
  map-based approach tailored to preserve first/last occurrences without sorting.

### Benchmarks

Benchmarks are provided in bench_test.go. Run them with:

    go test -bench . -benchmem

They measure UniqueKeys/UniqueValues performance for typical input sizes and variants.

Recommendations (short)

- Sorted result, low allocations / GC pressure: use UniqueKeys / UniqueValues (sort-based). They return a sorted slice and do not mutate the input.
- Fast, preserve insertion order: use UniqueKeysMap / UniqueValuesMap (map-based). Use the short aliases UniqueMap / UniqueValueMap if you prefer concise names. Faster but higher allocations; map-based variants fall back to the sort-based implementation when the estimated unique count exceeds FallbackThreshold.

Pruning suggestion

To keep the API simple, expose two functions per need:
- UniqueKeys / UniqueValues (sort-based) — the safe default
- UniqueKeysMap / UniqueValuesMap (map-based) — the fast insertion-order-preserving variant (short aliases: UniqueMap / UniqueValueMap)

Fallback behavior

- FallbackThreshold (package-level, default 8192) controls when the map-based impl will fall back to the sort-based implementation to avoid excessive memory/GC pressure on large unique sets.
- You can tune it at runtime:

```go
import "github.com/grimdork/sortprop"

func init() {
    sortprop.FallbackThreshold = 4096 // smaller → safer memory footprint
}
```

Notes

- We removed the Hybrid variant from the top-level API to keep the surface small; hybrid logic can be reintroduced if needed.
- Benchmarks focus on the common small-N case (20–1,000 properties). Extremely large inputs (100k+) are supported but not recommended unless you tune FallbackThreshold.
