package _19_random_flip_matrix

import "math/rand"

/**
519. 随机翻转矩阵

给你一个 m x n 的二元矩阵 matrix ，且所有值被初始化为 0 。请你设计一个算法，随机选取一个满足 matrix[i][j] == 0 的下标 (i, j) ，并将它的值变为 1 。所有满足 matrix[i][j] == 0 的下标 (i, j) 被选取的概率应当均等。

尽量最少调用内置的随机函数，并且优化时间和空间复杂度。

实现 Solution 类：

Solution(int m, int n) 使用二元矩阵的大小 m 和 n 初始化该对象
int[] flip() 返回一个满足 matrix[i][j] == 0 的随机下标 [i, j] ，并将其对应格子中的值变为 1
void reset() 将矩阵中所有的值重置为 0
*/

type Solution struct {
	m, n, cnt int
	set       map[int]int
}

func Constructor(m int, n int) Solution {
	return Solution{m, n, m * n, make(map[int]int)}
}

func (s *Solution) Flip() []int {
	x := rand.Intn(s.cnt)
	s.cnt--
	idx := x
	if _, ok := s.set[x]; ok {
		idx = s.set[x]
	}
	s.set[x] = s.cnt
	if _, ok := s.set[s.cnt]; ok {
		s.set[x] = s.set[s.cnt]
	}
	return []int{idx / s.n, idx % s.n}
}

func (s *Solution) Reset() {
	s.cnt = s.m * s.n
	s.set = make(map[int]int)
}
