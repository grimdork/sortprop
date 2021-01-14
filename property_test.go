package sortprop_test

import (
	"sort"
	"strconv"
	"testing"

	"github.com/grimdork/sortprop"
)

func TestAppendSort(t *testing.T) {
	list := sortprop.KeyProperties{}
	var i int64
	for i = 9; i >= 0; i-- {
		k := strconv.FormatInt(i, 10)
		v := "element " + k
		list = append(list, sortprop.Property{k, v})
	}

	sort.Sort(list)
	if list[0].Key != "0" || list[0].Value != "element 0" {
		t.Fail()
	}

	t.Logf("%+v", list)
}

func TestAppendSortValue(t *testing.T) {
	list := sortprop.ValueProperties{}
	var i int64
	for i = 9; i >= 0; i-- {
		k := strconv.FormatInt(i, 10)
		v := "element " + k
		list = append(list, sortprop.Property{k, v})
	}

	sort.Sort(list)
	if list[0].Key != "0" || list[0].Value != "element 0" {
		t.Fail()
	}

	t.Logf("%+v", list)
}

func TestUniqueKeys(t *testing.T) {
	kp := sortprop.KeyProperties{
		sortprop.Property{"a", "this is a1"},
		sortprop.Property{"a", "this is a2"},
		sortprop.Property{"a", "this is a3"},
		sortprop.Property{"b", "this is b1"},
		sortprop.Property{"b", "this is b2"},
		sortprop.Property{"c", "this is c1"},
		sortprop.Property{"c", "this is c2"},
		sortprop.Property{"c", "this is c3"},
		sortprop.Property{"c", "this is c4"},
	}

	t.Log("Before:")
	for _, p := range kp {
		t.Logf("%s = %s", p.Key, p.Value)
	}

	list := sortprop.UniqueKeys(kp, false)
	t.Log("After: (keep first)")
	for _, p := range list {
		t.Logf("%s = %s", p.Key, p.Value)
	}

	list = sortprop.UniqueKeys(kp, true)
	t.Log("After: (keep last)")
	for _, p := range list {
		t.Logf("%s = %s", p.Key, p.Value)
	}
}

func TestUniqueValues(t *testing.T) {
	vp := sortprop.ValueProperties{
		sortprop.Property{"a1", "this is a"},
		sortprop.Property{"a2", "this is a"},
		sortprop.Property{"a3", "this is a"},
		sortprop.Property{"b1", "this is b"},
		sortprop.Property{"b2", "this is b"},
		sortprop.Property{"c1", "this is c"},
		sortprop.Property{"c2", "this is c"},
		sortprop.Property{"c3", "this is c"},
		sortprop.Property{"c4", "this is c"},
	}

	t.Log("Before:")
	for _, p := range vp {
		t.Logf("%s = %s", p.Key, p.Value)
	}

	list := sortprop.UniqueValues(vp, false)
	t.Log("After: (keep first)")
	for _, p := range list {
		t.Logf("%s = %s", p.Key, p.Value)
	}

	list = sortprop.UniqueValues(vp, true)
	t.Log("After: (keep last)")
	for _, p := range list {
		t.Logf("%s = %s", p.Key, p.Value)
	}
}
