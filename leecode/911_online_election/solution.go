package _11_online_election

import "sort"

/*
给你两个整数数组 persons 和 times 。在选举中，第\ i\ 张票是在时刻为\ times[i]\ 时投给候选人 persons[i]\ 的。

对于发生在时刻 t 的每个查询，需要找出在\ t 时刻在选举中领先的候选人的编号。

在\ t 时刻投出的选票也将被计入我们的查询之中。在平局的情况下，最近获得投票的候选人将会获胜。

实现 TopVotedCandidate 类：

TopVotedCandidate(int[] persons, int[] times) 使用\ persons 和 times 数组初始化对象。
int q(int t) 根据前面描述的规则，返回在时刻 t 在选举中领先的候选人的编号。
*/

type TopVotedCandidate struct {
	tops, times []int
}

func Constructor(persons []int, times []int) TopVotedCandidate {
	var tops []int
	top := -1
	voteCount := make(map[int]int)
	for _, p := range persons {
		voteCount[p]++
		if voteCount[p] >= voteCount[top] {
			top = p
		}
		tops = append(tops, top)
	}
	return TopVotedCandidate{tops, times}
}

func (c *TopVotedCandidate) Q(t int) int {
	return c.tops[sort.SearchInts(c.times, t+1)-1]
}
