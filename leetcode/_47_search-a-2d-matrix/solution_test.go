package _47_search_a_2d_matrix

import (
	"testing"

	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/xzygis/algorithm/leetcode/utils"
)

func Test_searchMatrix(t *testing.T) {
	PatchConvey("Test search matrix", t, func() {
		matrix := utils.ParseMatrix("[[1,3,5,7],[10,11,16,20],[23,30,34,60]]")
		target := 3
		expected := true
		result := searchMatrix(matrix, target)
		So(result, ShouldResemble, expected)
	})

	PatchConvey("Test search matrix", t, func() {
		matrix := utils.ParseMatrix("[[1,3,5,7],[10,11,16,20],[23,30,34,60]]")
		target := 13
		expected := false
		result := searchMatrix(matrix, target)
		So(result, ShouldResemble, expected)
	})

	PatchConvey("Test search matrix", t, func() {
		matrix := utils.ParseMatrix("[[1]]")
		target := 2
		expected := false
		result := searchMatrix(matrix, target)
		So(result, ShouldResemble, expected)
	})

}
