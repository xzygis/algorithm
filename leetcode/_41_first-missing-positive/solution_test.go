package _41_first_missing_positive

import (
	"testing"

	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_searchMatrix(t *testing.T) {
	PatchConvey("Test fist missing positive", t, func() {
		nums := []int{1, 2, 0}
		expected := 3
		result := firstMissingPositive(nums)
		So(result, ShouldResemble, expected)
	})

	PatchConvey("Test fist missing positive", t, func() {
		nums := []int{3, 4, -1, 1}
		expected := 2
		result := firstMissingPositive(nums)
		So(result, ShouldResemble, expected)
	})

	PatchConvey("Test fist missing positive", t, func() {
		nums := []int{7, 8, 9, 11, 12}
		expected := 1
		result := firstMissingPositive(nums)
		So(result, ShouldResemble, expected)
	})

	PatchConvey("Test fist missing positive", t, func() {
		nums := []int{1}
		expected := 2
		result := firstMissingPositive(nums)
		So(result, ShouldResemble, expected)
	})
}
