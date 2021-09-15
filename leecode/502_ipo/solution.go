package _02_ipo

/**
假设 力扣（LeetCode）即将开始 IPO 。为了以更高的价格将股票卖给风险投资公司，力扣 希望在 IPO 之前开展一些项目以增加其资本。 由于资源有限，它只能在 IPO 之前完成最多 k 个不同的项目。帮助 力扣 设计完成最多 k 个不同项目后得到最大总资本的方式。

给你 n 个项目。对于每个项目 i ，它都有一个纯利润 profits[i] ，和启动该项目需要的最小资本 capital[i] 。

最初，你的资本为 w 。当你完成一个项目时，你将获得纯利润，且利润将被添加到你的总资本中。

总而言之，从给定项目中选择 最多 k 个不同项目的列表，以 最大化最终资本 ，并输出最终可获得的最多资本。

答案保证在 32 位有符号整数范围内。
*/

import (
	"container/heap"
	"sort"
)

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	type Pair struct {
		p int
		c int
	}
	n := len(profits)
	arr := make([]*Pair, n)
	for i := 0; i < n; i++ {
		arr[i] = &Pair{profits[i], capital[i]}
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].c < arr[j].c
	})
	h := &hp{}
	i := 0
	for k > 0 {
		for i < n && arr[i].c <= w {
			heap.Push(h, arr[i].p)
			i++
		}
		if h.Len() == 0 {
			break
		}
		w = w + heap.Pop(h).(int)
		k--
	}
	return w
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
