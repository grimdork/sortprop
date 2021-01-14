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
