package _48_rotate_image

import (
	"testing"

	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/xzygis/algorithm/leetcode/utils"
)

func Test_rotate(t *testing.T) {
	PatchConvey("Test rotate image", t, func() {
		matrix := utils.ParseMatrix("[[1,2,3],[4,5,6],[7,8,9]]")
		expected := utils.ParseMatrix("[[7,4,1],[8,5,2],[9,6,3]]")
		rotate(matrix)
		So(matrix, ShouldResemble, expected)
	})

	PatchConvey("Test rotate image", t, func() {
		matrix := utils.ParseMatrix("[[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]")
		expected := utils.ParseMatrix("[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]")
		rotate(matrix)
		So(matrix, ShouldResemble, expected)
	})

}
