package _73_set_matrix_zeroes

import (
	"testing"

	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/xzygis/algorithm/leetcode/utils"
)

func Test_setZeroes(t *testing.T) {
	PatchConvey("Test overlapping intervals", t, func() {
		matrix := utils.ParseMatrix("[[1,1,1],[1,0,1],[1,1,1]]")
		expected := utils.ParseMatrix("[[1,0,1],[0,0,0],[1,0,1]]")
		setZeroes(matrix)
		So(matrix, ShouldResemble, expected)
	})

	PatchConvey("Test overlapping intervals", t, func() {
		matrix := utils.ParseMatrix("[[0,1,2,0],[3,4,5,2],[1,3,1,5]]")
		expected := utils.ParseMatrix("[[0,0,0,0],[0,4,5,0],[0,3,1,0]]")
		setZeroes(matrix)
		So(matrix, ShouldResemble, expected)
	})

}
