package _75_koko_eating_bananas

import (
	"math"
)

/**
珂珂喜欢吃香蕉。这里有 N 堆香蕉，第 i 堆中有 piles[i] 根香蕉。警卫已经离开了，将在 H 小时后回来。

珂珂可以决定她吃香蕉的速度 K （单位：根/小时）。每个小时，她将会选择一堆香蕉，从中吃掉 K 根。如果这堆香蕉少于 K 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。

珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。

返回她可以在 H 小时内吃掉所有香蕉的最小速度 K（K 为整数）。
*/

func minEatingSpeed(piles []int, h int) int {
	var possible func(k int) bool
	possible = func(k int) bool {
		t := 0
		for _, p := range piles {
			t += (p-1)/k + 1
		}
		return t <= h
	}
	lo, hi := 1, math.MaxInt32
	for lo < hi {
		mid := lo + (hi-lo)/2
		if !possible(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
