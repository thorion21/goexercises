package list

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateEmpty(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Create a new list and check if is not null", t, func() {
		l := New([]int{1, 2, 3})
		So(l, ShouldNotBeNil)

		Convey("Assert list contains three elements", func() {
			So(l.Size(), ShouldEqual, 3)
		})
	})

}

func TestPush(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Create a new list", t, func() {
		l := New([]int{1, 2, 3})
		So(l, ShouldNotBeNil)

		Convey("Push 2 times", func() {
			l.Push(4)
			l.Push(999)

			Convey("Check if list has the correct length", func() {
				So(l.Size(), ShouldEqual, 5)

				Convey("Check if the last item is correct", func() {
					x := l.Array()[4]
					So(x, ShouldEqual, 999)
				})
			})
		})
	})

}

func TestPop(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Create a new list with 2 elements", t, func() {
		l := New([]int{1, 2})

		Convey("Pop 1 time and check for success", func() {
			_, err := l.Pop()
			So(err, ShouldBeNil)

			Convey("Check if list has the correct length", func() {
				So(l.Size(), ShouldEqual, 1)

				Convey("Pop until empty list", func() {
					l.Pop()

					Convey("Pop on empty list", func() {
						_, err1 := l.Pop()
						So(err1, ShouldNotBeNil)
						So(l.Size(), ShouldBeZeroValue)
					})
				})
			})
		})
	})

}

func TestArray(t *testing.T) {

	// Only pass t into top-level Convey calls
	Convey("Create a new list and check if is not null", t, func() {
		l := New([]int{1, 2, 3})
		So(l, ShouldNotBeNil)

		Convey("Return the list as an array", func() {
			So(l.Reverse().Array(), ShouldResemble, []int{3, 2, 1})
		})
	})

}
