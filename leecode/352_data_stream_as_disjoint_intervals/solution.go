package _52_data_stream_as_disjoint_intervals

/**
352. 将数据流变为多个不相交区间

给你一个由非负整数 a1, a2, ..., an 组成的数据流输入，
请你将到目前为止看到的数字总结为不相交的区间列表。

实现 SummaryRanges 类：

SummaryRanges() 使用一个空数据流初始化对象。
void addNum(int val) 向数据流中加入整数 val 。
int[][] getIntervals() 以不相交区间 [starti, endi] 的列表形式返回对数据流中整数的总结。
*/

type SummaryRanges struct {
	nums map[int]bool
}

func Constructor() SummaryRanges {
	return SummaryRanges{make(map[int]bool, 0)}
}

func (s *SummaryRanges) AddNum(val int) {
	s.nums[val] = true
}

func (s *SummaryRanges) GetIntervals() [][]int {
	start, end := -1, -1
	ans := make([][]int, 0)
	for i := 0; i < 10001; i++ {
		if _, has := s.nums[i]; has {
			if start == -1 {
				start = i
				end = i
			} else {
				end = i
			}
		} else {
			if start != -1 {
				ans = append(ans, []int{start, end})
				start = -1
				end = -1
			}
		}
	}
	if start != -1 {
		ans = append(ans, []int{start, end})
	}
	return ans
}
