// Copyright 2021-2022 Peter Bigot Consulting, LLC
// SPDX-License-Identifier: Apache-2.0

package set

// Set captures a set of strings.
type Set[E comparable] map[E]struct{}

// MakeSet constructs and returns a Set containing the passed elements.  The
// returned set is nil if it has no members.
func MakeSet[E comparable](s ...E) Set[E] {
	if len(s) == 0 {
		return nil
	}
	rv := make(Set[E], len(s))
	for _, e := range s {
		rv[e] = struct{}{}
	}
	return rv
}

// Add adds the passed elements to the set and returns it.  The returned set
// is not nil even if it is empty.
func (ss Set[E]) Add(as ...E) Set[E] {
	if ss == nil {
		ss = make(Set[E], len(as))
	}
	for _, s := range as {
		ss[s] = struct{}{}
	}
	return ss
}

// Remove removes the passed elements from the set.  There is no error if a
// string to be removed is not present.  The returned set is not nil even if
// it is empty.
func (ss Set[E]) Remove(rs ...E) Set[E] {
	for _, s := range rs {
		delete(ss, s)
	}
	return ss
}

// Has returns true if and only if s is in the set.
func (ss Set[E]) Has(s E) (ok bool) {
	if ss != nil {
		_, ok = ss[s]
	}
	return ok
}

// Minus returns a new set containing the members of the receiver that are not
// present in the parameter.  The content of the receiving Set is not
// changed.  If the resulting set has no elements a nil set is returned.
//
// The idiom "ss.Minus(nil)" creates a copy of ss.
func (sl Set[E]) Minus(sr Set[E]) Set[E] {
	var rv Set[E]
	for s := range sl {
		if !sr.Has(s) {
			rv = rv.Add(s)
		}
	}
	return rv
}

// Elements returns a slice of the members of the set, or nil if the set is
// empty.  The order of elements in the slice is not specified.
func (sl Set[E]) Elements() []E {
	if len(sl) == 0 {
		return nil
	}
	rv := make([]E, 0, len(sl))
	for s := range sl {
		rv = append(rv, s)
	}
	return rv
}
