package _705_maximun_number_of_eaten_apples

import (
	"container/heap"
	"sort"
)

/*
1705. 吃苹果的最大数目

有一棵特殊的苹果树，一连 n 天，每天都可以长出若干个苹果。在第 i 天，树上会长出 apples[i] 个苹果，这些苹果将会在 days[i] 天后（也就是说，第 i + days[i] 天时）腐烂，变得无法食用。也可能有那么几天，树上不会长出新的苹果，此时用 apples[i] == 0 且 days[i] == 0 表示。

你打算每天 最多 吃一个苹果来保证营养均衡。注意，你可以在这 n 天之后继续吃苹果。

给你两个长度为 n 的整数数组 days 和 apples ，返回你可以吃掉的苹果的最大数目。
*/

func eatenApples(apples []int, days []int) (ans int) {
	n := len(apples)
	// 过期日期对应的腐烂苹果数
	m := make(map[int]int)
	// 过期日期优先队列
	h := &Heap{}
	// 遍历所有日期直到没有过期苹果
	for i := 0; i < n || h.Len() > 0; i++ {
		// 当天过期苹果从队列中移除
		if h.Len() > 0 && h.IntSlice[0] == i {
			delete(m, i)
			heap.Remove(h, 0)
		}
		// 当天有长出苹果，则过期时间入优先队列，并更新待过期苹果数
		if i < n && apples[i] > 0 {
			cutDay := i + days[i]
			if m[cutDay] == 0 {
				heap.Push(h, cutDay)
			}
			m[cutDay] += apples[i]
		}
		// 有待过期苹果说明当天有苹果吃
		if h.Len() > 0 {
			ans++
			// 食用最早过期的苹果
			m[h.IntSlice[0]]--
			if m[h.IntSlice[0]] == 0 {
				delete(m, i)
				heap.Remove(h, 0)
			}
		}
	}
	return
}

type Heap struct {
	sort.IntSlice
}

func (h *Heap) Less(i, j int) bool {
	return h.IntSlice[i] < h.IntSlice[j]
}

func (h *Heap) Push(x interface{}) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *Heap) Pop() interface{} {
	a := h.IntSlice
	x := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return x
}
