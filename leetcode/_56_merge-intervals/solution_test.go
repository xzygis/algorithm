package _56_merge_intervals

import (
	"testing"

	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/xzygis/algorithm/leetcode/utils"
)

func Test_merge(t *testing.T) {
	PatchConvey("Test overlapping intervals", t, func() {
		intervals := utils.ParseMatrix("[[1,3],[2,6],[8,10],[15,18]]")
		expected := utils.ParseMatrix("[[1,6],[8,10],[15,18]]")
		result := merge(intervals)
		So(result, ShouldResemble, expected)
	})

	PatchConvey("Test overlapping intervals", t, func() {
		intervals := utils.ParseMatrix("[[1,4],[4,5]]")
		expected := utils.ParseMatrix("[[1,5]]")
		result := merge(intervals)
		So(result, ShouldResemble, expected)
	})

	PatchConvey("Test overlapping intervals", t, func() {
		intervals := utils.ParseMatrix("[[1,4],[2,3]]")
		expected := utils.ParseMatrix("[[1,4]]")
		result := merge(intervals)
		So(result, ShouldResemble, expected)
	})

}
