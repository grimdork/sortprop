# sortprop
Package for sortable key-value properties, like HTTP headers, with functions to remove duplicates.

## Usage

### Key-sortable slices

Note: using sort.Sort on KeyProperties/ValueProperties sorts in-place and mutates the slice. If you need a non-mutating sort, copy the slice before sorting.


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

- We removed the Hybrid variant from the top-level API to keep the surface small; hybrid logic has been removed from the README and is not part of the public surface.
- Benchmarks focus on the common small-N case (20–1,000 properties). Extremely large inputs (100k+) are supported but not recommended unless you tune FallbackThreshold.

Full benchmark summary (representative; run with `go test -bench . -benchmem` on darwin/arm64 Apple M1 Max)

n | uniqueRatio | Sort_first (ns/op, B/op, allocs) | Map_first (ns/op, B/op, allocs)
---|---:|---:|---:
20 | 1.00 | 451 ns/op, 1.4 KB/op, 3 allocs/op | 1.12 µs/op, 1.7 KB/op, 6 allocs/op
50 | 0.10 | 1.75 µs/op, 3.6 KB/op, 3 allocs/op | 1.30 µs/op, 2.1 KB/op, 6 allocs/op
100 | 0.50 | 4.8 µs/op, 6.9 KB/op, 3 allocs/op | 3.0 µs/op, 5.4 KB/op, 6 allocs/op
1000 | 1.00 (mostly unique) | 15.8 µs/op, 65.5 KB/op, 3 allocs/op | 42.6 µs/op, 87.9 KB/op, 8 allocs/op
1000 | 0.01 (many duplicates) | 41.3 µs/op, 65.5 KB/op, 3 allocs/op | 14.3 µs/op, 55.1 KB/op, 8 allocs/op

(Full raw bench output is available by running the benchmarks locally; these representative numbers are from the run used to validate the map-based implementation.)

When to use UniqueValues (concrete examples)

- Inverted-index / tag overview: If you have a list of records where multiple records reference the same tag string (Value), deduping by Value gives a single representative per tag for UI overviews or tag lists.

- Content canonicalization: When different source records (different Keys) point to identical content (Value is a content hash or normalized text), dedupe by Value to keep one canonical record per content blob.

- Sensor/metric grouping: Sensors may report named measurements as Value while keys are sensor IDs. To display one sample per measurement label (e.g., "temperature"), dedupe by Value.

- De-duplicating lookups: If you build a list of external links (Value) collected from many sources (Keys), keeping unique Values prevents repeated downloads or checks.

Example (pseudo-Go):

```go
// keep one representative per content hash
vp := sortprop.ValueProperties{
    {Key: "src1", Value: "sha1:abc"},
    {Key: "src2", Value: "sha1:abc"},
    {Key: "src3", Value: "sha1:def"},
}
unique := sortprop.UniqueValues(vp, true) // keep last occurrence of each value
```

These cases are common enough to justify keeping UniqueValues as a first-class API: it communicates intent clearly and avoids forcing callers to use more general (and more complex) selector-based APIs.
