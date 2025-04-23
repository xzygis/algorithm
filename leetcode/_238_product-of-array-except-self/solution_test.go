package _238_product_of_array_except_self

import (
	"testing"

	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
)

func Test_productExceptSelf(t *testing.T) {
	PatchConvey("Test product of array except self", t, func() {
		nums := []int{1, 2, 3, 4}
		expected := []int{24, 12, 8, 6}
		result := productExceptSelf(nums)
		So(result, ShouldResemble, expected)
	})

	PatchConvey("Test product of array except self", t, func() {
		nums := []int{-1, 1, 0, -3, 3}
		expected := []int{0, 0, 9, 0, 0}
		result := productExceptSelf(nums)
		So(result, ShouldResemble, expected)
	})

	PatchConvey("Test product of array except self", t, func() {
		nums := []int{-1, 1}
		expected := []int{1, -1}
		result := productExceptSelf(nums)
		So(result, ShouldResemble, expected)
	})
}
