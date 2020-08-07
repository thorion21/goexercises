package myset

import (
	"strconv"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNew(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Create set and check", t, func() {
		s := New()
		So(s, ShouldNotBeNil)
		So(s.IsEmpty(), ShouldBeTrue)
	})

	Convey("Create set from slice and check", t, func() {
		s := NewFromSlice([]string{"1", "2", "3"})
		So(s, ShouldNotBeNil)
		So(s.IsEmpty(), ShouldBeFalse)
	})
}

func TestAdd(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Create set and check", t, func() {
		s := New()
		So(s, ShouldNotBeNil)

		Convey("Add 3 elements (1, 11, 111)", func() {
			s.Add("1")
			s.Add("11")
			s.Add("111")

			So(s.IsEmpty(), ShouldBeFalse)

			Convey("Check for existence of 11", func() {
				s.Add("1")
				s.Add("11")
				s.Add("111")

				So(s.Has("11"), ShouldBeTrue)
			})
		})
	})
}

func TestEqual(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Create identical sets and check for equality", t, func() {
		s1 := NewFromSlice([]string{"1", "2", "3"})
		s2 := NewFromSlice([]string{"1", "2", "3"})
		So(Equal(s1, s2), ShouldBeTrue)
	})

	Convey("Create different sets and check for equality", t, func() {
		s1 := NewFromSlice([]string{"1", "2", "3"})
		s2 := NewFromSlice([]string{"1", "2", "4"})
		So(Equal(s1, s2), ShouldBeFalse)
	})
}

func TestIntersect(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Create two sets with 2 elements in common", t, func() {
		s1 := NewFromSlice([]string{"1", "2", "3"})
		s2 := NewFromSlice([]string{"7", "2", "3"})
		So(len(Intersection(s1, s2).items), ShouldEqual, 2)
		So(Intersection(s1, s2).String(), ShouldBeIn, []string{"{\"2\", \"3\"}", "{\"3\", \"2\"}"})
	})
}

func TestUnion(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Create two sets and union them together", t, func() {
		s1 := NewFromSlice([]string{"1", "2", "3"})
		s2 := NewFromSlice([]string{"7", "2", "3"})

		res := Union(s1, s2)
		So(len(res.items), ShouldEqual, 4)

		hasAllItems := true
		for _, e := range []int{1, 2, 3, 7} {
			if !res.Has(strconv.Itoa(e)) {
				hasAllItems = false
				break
			}
		}

		So(hasAllItems, ShouldBeTrue)
	})
}

func TestDifference(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Create two sets and apply difference", t, func() {
		s1 := NewFromSlice([]string{"1", "2", "3"})
		s2 := NewFromSlice([]string{"7", "2", "3"})

		res := Difference(s1, s2)
		So(len(res.items), ShouldEqual, 1)
		So(res.Has("1"), ShouldBeTrue)
	})
}

func TestDisjoint(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Create two sets with some common items", t, func() {
		s1 := NewFromSlice([]string{"1", "2", "3"})
		s2 := NewFromSlice([]string{"7", "2", "3"})
		So(Disjoint(s1, s2), ShouldBeFalse)
	})

	Convey("Create two sets with no common items", t, func() {
		s1 := NewFromSlice([]string{"1", "2", "3"})
		s2 := NewFromSlice([]string{"7", "9", "12"})
		So(Disjoint(s1, s2), ShouldBeTrue)
	})
}

func TestSubset(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Create two sets with some common items", t, func() {
		s1 := NewFromSlice([]string{"2", "3"})
		s2 := NewFromSlice([]string{"7", "2", "3"})
		So(Subset(s1, s2), ShouldBeTrue)
	})

	Convey("Create two sets with no common items", t, func() {
		s1 := NewFromSlice([]string{"1", "4", "5"})
		s2 := NewFromSlice([]string{"7", "9", "12"})
		So(Subset(s1, s2), ShouldBeFalse)
	})
}
