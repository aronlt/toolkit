package ds

import (
	"fmt"

	"github.com/aronlt/toolkit/ttypes"
)

type OrderSet[K ttypes.Ordered] struct {
	set  BuiltinSet[K]
	list *SList[K]
}

func NewOrderSet[K ttypes.Ordered]() *OrderSet[K] {
	return &OrderSet[K]{
		set:  NewSet[K](),
		list: NewSList[K](),
	}
}

func OrderSetFromUnpack[K ttypes.Ordered](ks ...K) *OrderSet[K] {
	o := NewOrderSet[K]()
	for _, k := range ks {
		o.Insert(k)
	}
	return o
}

func OrderSetFromSet[K ttypes.Ordered](set BuiltinSet[K]) *OrderSet[K] {
	os := NewOrderSet[K]()
	set.ForEach(func(k K) {
		os.Insert(k)
	})
	return os
}

// IsEmpty implements the Container interface.
func (o *OrderSet[K]) IsEmpty() bool {
	return o.set.Len() == 0
}

func (o *OrderSet[K]) Len() int {
	return o.set.Len()
}

// Clear implements the Container interface.
func (o *OrderSet[K]) Clear() {
	o.set.Clear()
	o.list.Clear()
}

// Has implements the Set interface.
func (o *OrderSet[K]) Has(k K) bool {
	return o.set.Has(k)
}

func (o *OrderSet[K]) Insert(k K) bool {
	if o.set.Has(k) {
		return false
	}
	_ = o.set.Insert(k)
	o.list.InsertLessBound(k, func(a, b K) bool {
		return a < b
	})
	return true
}

func (o *OrderSet[K]) InsertN(ks ...K) int {
	count := 0
	for _, k := range ks {
		if ok := o.Insert(k); ok {
			count++
		}
	}
	return count
}

func (o *OrderSet[K]) Remove(k K) bool {
	if ok := o.set.Has(k); !ok {
		return false
	}
	o.set.Delete(k)
	_ = o.list.RemoveValue(k, func(a, b K) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	})

	return true
}

func (o *OrderSet[K]) RemoveN(ks ...K) int {
	count := 0
	for _, k := range ks {
		if ok := o.Remove(k); ok {
			count++
		}
	}
	return count
}

func (o *OrderSet[K]) DeleteN(ks ...K) {
	for _, k := range ks {
		_ = o.Remove(k)
	}
}

func (o *OrderSet[K]) Delete(k K) {
	_ = o.Remove(k)
}

// Keys return a copy of all keys as a slice.
func (o *OrderSet[K]) Keys() *SList[K] {
	newList := NewSList[K]()
	o.list.ForEach(func(k K) {
		newList.PushBack(k)
	})
	return newList
}

// ForEach implements the Set interface.
func (o *OrderSet[K]) ForEach(cb func(k K)) {
	o.list.ForEach(cb)
}

// ForEachIf implements the Container interface.
func (o *OrderSet[K]) ForEachIf(cb func(k K) bool) {
	o.list.ForEachIf(cb)
}

// String implements the fmt.Stringer interface.
func (o *OrderSet[K]) String() string {
	return fmt.Sprintf("BuiltinSet %v", o.Keys())
}

// Update adds all elements from other to set. set |= other.
func (o *OrderSet[K]) Update(other *OrderSet[K]) {
	other.ForEach(func(k K) {
		o.InsertN(k)
	})
}

// Union returns a new set with elements from the set and other.
func (o *OrderSet[K]) Union(other *OrderSet[K]) *OrderSet[K] {
	result := NewOrderSet[K]()
	result.Update(o)
	result.Update(other)
	return result
}

func orderSetSize[K ttypes.Ordered](a, b *OrderSet[K]) (small, large *OrderSet[K]) {
	if a.Len() < b.Len() {
		return a, b
	}
	return b, a
}

// Intersection returns a new set with elements common to the set and other.
func (o *OrderSet[K]) Intersection(other *OrderSet[K]) *OrderSet[K] {
	result := NewOrderSet[K]()
	small, large := orderSetSize(o, other)
	small.ForEach(func(k K) {
		if large.Has(k) {
			result.Insert(k)
		}
	})
	return result
}

// Difference returns a new set with elements in the set that are not in other.
func (o *OrderSet[K]) Difference(other *OrderSet[K]) *OrderSet[K] {
	result := NewOrderSet[K]()
	o.ForEach(func(k K) {
		if !other.Has(k) {
			result.Insert(k)
		}
	})
	return result
}

// IsDisjointOf return True if the set has no elements in common with other.
// Sets are disjoint if and only if their intersection is the empty set.
func (o *OrderSet[K]) IsDisjointOf(other *OrderSet[K]) bool {
	small, large := orderSetSize(o, other)

	ok := true
	small.ForEachIf(func(k K) bool {
		if large.Has(k) {
			ok = false
			return false
		}
		return true
	})
	return ok
}

// IsSubsetOf tests whether every element in the set is in other.
func (o *OrderSet[K]) IsSubsetOf(other *OrderSet[K]) bool {
	if o.Len() > other.Len() {
		return false
	}
	ok := true
	o.ForEachIf(func(k K) bool {
		if !other.Has(k) {
			ok = false
			return false
		}
		return true
	})
	return ok
}

// IsSupersetOf tests whether every element in other is in the set.
func (o *OrderSet[K]) IsSupersetOf(other *OrderSet[K]) bool {
	return other.IsSubsetOf(o)
}

func (o *OrderSet[K]) Equal(other *OrderSet[K]) bool {
	if o.Len() != other.Len() {
		return false
	}
	return o.IsSubsetOf(other) && other.IsSubsetOf(o)
}
