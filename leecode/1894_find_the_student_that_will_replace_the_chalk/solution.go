package _1894_find_the_student_that_will_replace_the_chalk

/**
1894. 找到需要补充粉笔的学生编号

一个班级里有n个学生，编号为 0到 n - 1。每个学生会依次回答问题，编号为 0的学生先回答，然后是编号为 1的学生，以此类推，直到编号为 n - 1的学生，然后老师会重复这个过程，重新从编号为 0的学生开始回答问题。

给你一个长度为 n且下标从 0开始的整数数组chalk和一个整数k。一开始粉笔盒里总共有k支粉笔。当编号为i的学生回答问题时，他会消耗 chalk[i]支粉笔。如果剩余粉笔数量 严格小于chalk[i]，那么学生 i需要 补充粉笔。

请你返回需要 补充粉笔的学生 编号。
*/

func chalkReplacer(chalk []int, k int) int {
	n := len(chalk)
	for i := 1; i < n; i++ {
		chalk[i] = chalk[i] + chalk[i-1]
	}

	k %= chalk[n-1]
	i := 0
	for i = range chalk {
		if chalk[i] > k {
			break
		}
	}
	return i
}
