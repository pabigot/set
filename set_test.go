// Copyright 2022 Peter Bigot Consulting, LLC
// SPDX-License-Identifier: Apache-2.0

package set

import (
	"sort"
	"testing"
)

func TestStringSet(t *testing.T) {
	var nilSet Set[string]
	if nilSet != nil {
		t.Error("nilSet is not nil")
	}

	emptySet := make(Set[string])
	if len(emptySet) != 0 {
		t.Error("emptySet not empty")
	}
	if !emptySet.IsEqual(nilSet) || !nilSet.IsEqual(emptySet) {
		t.Error("equal on empty and nil set wrong")
	}
	if !emptySet.IsSubsetOf(nilSet) || !nilSet.IsSubsetOf(emptySet) {
		t.Error("subset on empty and nil set wrong")
	}

	if nilSet.Has("a") {
		t.Error("nilSet found value")
	}
	if emptySet.Has("a") {
		t.Error("emptySet found value")
	}

	var ss Set[string]
	var sl []string
	if ss = MakeSet[string](); ss != nil {
		t.Error("make from empty should be nil")
	}
	if ss = nilSet.Remove("b"); ss != nil {
		t.Error("remove from nil should be nil")
	}
	if ss = emptySet.Remove("b"); ss == nil {
		t.Error("remove from empty should not be nil")
	}

	// staticcheck says the compares are never true
	//if emptySet == nil {
	//	t.Error("emptySet is nil")
	//}
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
	if ss.IsEqual(nilSet) || ss.IsEqual(emptySet) {
		t.Error("non-empty equals nil or empty")
	}
	if !nilSet.IsSubsetOf(ss) || !emptySet.IsSubsetOf(ss) {
		t.Error("nil or empty is not subset of non-empty")
	}

	ss = ss.Add("b")
	if v := len(ss); v != 2 {
		t.Errorf("add produced wrong len: %d", v)
	}
	if !ss.Has("b") {
		t.Error("set missing element")
	}
	if !ss.IsEqual(MakeSet[string]("a", "b")) {
		t.Error("set not equal expected")
	}

	sa := Set[string]{}.Add("a")
	if v := len(sa); v != 1 {
		t.Errorf("make from empty wrong len: %d", v)
	}
	if !sa.Has("a") {
		t.Error("set missing element")
	}
	if !sa.IsSubsetOf(ss) {
		t.Errorf("subSet false but %v ⊆ %v", sa, ss)
	}
	if ss.IsSubsetOf(sa) {
		t.Errorf("subSet true but %v ⊈ %v", sa, ss)
	}
	if st := MakeSet[string]("c"); st.IsSubsetOf(ss) {
		t.Errorf("subSet true but %v ⊈ %v", st, ss)
	} else if st.IsEqual(sa) {
		t.Errorf("equal true but %v ≠ %v", st, ss)
	}

	sc := ss.Copy()
	if v := len(sc); v != len(ss) {
		t.Errorf("copy wrong size: %d", v)
	}
	if !ss.IsSubsetOf(sc) {
		t.Error("subset of equal wrong")
	}
	if !ss.IsEqual(sc) {
		t.Error("equal wrong")
	}
	sc.Remove("a", "b")
	if v := len(sc); v != len(ss)-2 {
		t.Errorf("copy affects original")
	}
	if sc.Copy() != nil {
		t.Errorf("copy empty not nil")
	}
	sc = nil
	if sc.Copy() != nil {
		t.Errorf("copy nil not nil")
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
