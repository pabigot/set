// Copyright 2022 Peter Bigot Consulting, LLC
// SPDX-License-Identifier: Apache-2.0

package set

import (
	"sort"
	"testing"
)

func TestStringSet(t *testing.T) {
	var ss Set[string]
	var sl []string
	if ss.Has("a") {
		t.Error("nil Set found value")
	}
	if ss = MakeSet[string](); ss != nil {
		t.Error("make from nil not nil")
	}
	if ss = ss.Remove("b"); ss != nil {
		t.Error("remove created non-nil set")
	}
	// staticcheck says the compares are never true
	// ss = Set[string]{}
	// if ss == nil {
	// 	t.Error("not expected")
	// }
	// sl := make([]string, 0)
	// if sl == nil {
	// 	t.Error("empty slice nil")
	// }
	if ss = MakeSet[string](sl...); ss != nil {
		t.Error("make from empty not nil")
	}

	sl = append(sl, "a")
	ss = MakeSet[string](sl...)
	if v := len(ss); v != 1 {
		t.Errorf("make from singleton wrong len: %d", v)
	}
	if !ss.Has("a") {
		t.Error("set missing element")
	}

	ss = ss.Add("b")
	if v := len(ss); v != 2 {
		t.Errorf("add produced wrong len: %d", v)
	}
	if !ss.Has("b") {
		t.Error("set missing element")
	}

	sa := Set[string]{}.Add("a")
	if v := len(sa); v != 1 {
		t.Errorf("make from empty wrong len: %d", v)
	}
	if !sa.Has("a") {
		t.Error("set missing element")
	}

	sm := ss.Minus(nil)
	if len(sm) != len(ss) {
		t.Errorf("set minus nil wrong len: %d", len(sm))
	}
	sm.Remove("a")
	if len(sm) == len(ss) {
		t.Error("set minus nil tied sets")
	}

	sb := ss.Minus(sa)
	if v := len(sb); v != 1 {
		t.Errorf("set minus sa wrong len: %d", v)
	}
	if v := len(ss); v != 2 {
		t.Error("set minus changed source")
	}

	if sb.Has("a") || !sb.Has("b") {
		t.Errorf("set minus wrong result")
	}
	sl = ss.Elements()
	sort.Strings(sl)

	if len(sl) != 2 || sl[0] != "a" || sl[1] != "b" {
		t.Error("elements wrong result")
	}

	ss = ss.Remove("c")
	if v := len(ss); v != 2 {
		t.Errorf("remove absent changed len: %d", v)
	}
	ss.Remove("b")
	if v := len(ss); v != 1 {
		t.Errorf("remove present wrong len: %d", v)
	}
	ss = nil
	ss = ss.Remove("a", "b", "c")
	if ss != nil {
		t.Error("remove created non-nil set")
	}
	ss = Set[string]{}.Add("a").Remove("a")
	if ss == nil {
		t.Error("add/remove created nil set")
	}
	if v := len(ss); v != 0 {
		t.Errorf("add/remove not empty: %d", v)
	}
	sl = ss.Elements()
	if sl != nil {
		t.Error("elements of empty set not nil")
	}
}
