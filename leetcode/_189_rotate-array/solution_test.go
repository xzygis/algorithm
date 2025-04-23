package _189_rotate_array

import (
	"testing"

	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_Rotate(t *testing.T) {
	PatchConvey("Test ratate", t, func() {
		nums := []int{1, 2, 3, 4, 5, 6, 7}
		k := 3
		expected := []int{5, 6, 7, 1, 2, 3, 4}
		rotate(nums, k)
		So(nums, ShouldResemble, expected)
	})

	PatchConvey("Test ratate", t, func() {
		nums := []int{-1, -100, 3, 99}
		k := 2
		expected := []int{3, 99, -1, -100}
		rotate(nums, k)
		So(nums, ShouldResemble, expected)
	})

	PatchConvey("Test ratate", t, func() {
		nums := []int{-1, -100, 3, 99}
		k := 1
		expected := []int{99, -1, -100, 3}
		rotate(nums, k)
		So(nums, ShouldResemble, expected)
	})

	PatchConvey("Test ratate", t, func() {
		nums := []int{-1, -100, 3, 99}
		k := 0
		expected := []int{-1, -100, 3, 99}
		rotate(nums, k)
		So(nums, ShouldResemble, expected)
	})

	PatchConvey("Test ratate", t, func() {
		nums := []int{-1}
		k := 2
		expected := []int{-1}
		rotate(nums, k)
		So(nums, ShouldResemble, expected)
	})

}
