package _54_spiral_matrix

import (
	"testing"

	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/xzygis/algorithm/leetcode/utils"
)

func Test_spiralOrder(t *testing.T) {
	PatchConvey("Test split order", t, func() {
		matrix := utils.ParseMatrix("[[1,2,3],[4,5,6],[7,8,9]]")
		expected := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
		ans := spiralOrder(matrix)
		So(ans, ShouldResemble, expected)
	})

	PatchConvey("Test split order", t, func() {
		matrix := utils.ParseMatrix("[[1,2,3,4],[5,6,7,8],[9,10,11,12]]")
		expected := []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}
		ans := spiralOrder(matrix)
		So(ans, ShouldResemble, expected)
	})

}
