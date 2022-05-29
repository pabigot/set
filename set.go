// Copyright 2021-2022 Peter Bigot Consulting, LLC
// SPDX-License-Identifier: Apache-2.0

package set

// Set captures a set of elements.
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
func (se Set[E]) Add(es ...E) Set[E] {
	if se == nil {
		se = make(Set[E], len(es))
	}
	for _, e := range es {
		se[e] = struct{}{}
	}
	return se
}

// Remove removes the passed elements from the set.  There is no error if an
// element to be removed is not present.  The returned set is not nil even if
// it is empty.
func (se Set[E]) Remove(es ...E) Set[E] {
	for _, e := range es {
		delete(se, e)
	}
	return se
}

// Has returns true if and only if s is in the set.
func (se Set[E]) Has(e E) (ok bool) {
	if se != nil {
		_, ok = se[e]
	}
	return ok
}

// Minus returns a new set containing the members of the receiver that are not
// present in the parameter.  The content of the receiving Set is not
// changed.  If the resulting set has no elements a nil set is returned.
//
// The idiom "se.Minus(nil)" creates a copy of se.
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
func (se Set[E]) Elements() []E {
	if len(se) == 0 {
		return nil
	}
	rv := make([]E, 0, len(se))
	for e := range se {
		rv = append(rv, e)
	}
	return rv
}
