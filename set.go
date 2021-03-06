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

// IsSubsetOf returns true if and only if all the elements in the receiver are
// in ss.
func (sub Set[E]) IsSubsetOf(ss Set[E]) bool {
	if len(sub) > len(ss) {
		return false
	}
	for e := range sub {
		if !ss.Has(e) {
			return false
		}
	}
	return true
}

// IsEqual returns true if and only if all the elements in the receiver are in
// the argument and vice versa.
func (sl Set[E]) IsEqual(sr Set[E]) bool {
	if len(sl) != len(sr) {
		return false
	}
	for e := range sl {
		if !sr.Has(e) {
			return false
		}
	}
	return true
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

// Copy creates a copy of the set.  If the resulting set is empty a nil set is
// returned.
func (se Set[E]) Copy() (rv Set[E]) {
	if n := len(se); n != 0 {
		rv = make(Set[E], n)
		for e := range se {
			rv[e] = struct{}{}
		}
	}
	return rv
}

// Union returns the set of all elements that are in either the receiver and
// sr.  If the resulting set is empty a nil set is returned.
func (sl Set[E]) Union(sr Set[E]) (rv Set[E]) {
	if len(sl) != 0 {
		rv = sl.Copy()
		for e := range sr {
			rv = rv.Add(e)
		}
	} else if len(sr) != 0 {
		rv = sr
	}
	return rv
}

// Intersect returns the set of all elements that are in both the receiver and
// sr.  If the resulting set is empty a nil set is returned.
func (sl Set[E]) Intersect(sr Set[E]) (rv Set[E]) {
	if sl != nil && sr != nil {
		for e := range sl {
			if sr.Has(e) {
				rv = rv.Add(e)
			}
		}
	}
	return rv
}
