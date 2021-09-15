package _69_pancake_sorting

/**
给你一个整数数组 arr ，请使用 煎饼翻转 完成对数组的排序。

一次煎饼翻转的执行过程如下：

选择一个整数 k ，1 <= k <= arr.length
反转子数组 arr[0...k-1]（下标从 0 开始）
例如，arr = [3,2,1,4] ，选择 k = 3 进行一次煎饼翻转，反转子数组 [3,2,1] ，得到 arr = [1,2,3,4] 。

以数组形式返回能使 arr 有序的煎饼翻转操作所对应的 k 值序列。任何将数组排序且翻转次数在10 * arr.length 范围内的有效答案都将被判断为正确。
*/

func pancakeSort(arr []int) []int {
	var res []int
	reverse := func(arr []int, k int) {
		i, j := 0, k
		for i < j {
			arr[i], arr[j] = arr[j], arr[i]
			i, j = i+1, j-1
		}
	}
	type sortFunc func(arr []int, k int)
	var sort sortFunc
	sort = func(arr []int, k int) {
		if k == 1 {
			return
		}
		maxV, maxI := 0, 0
		for i := 0; i < k; i++ {
			if arr[i] > maxV {
				maxV = arr[i]
				maxI = i
			}
		}
		if maxI < k-1 {
			reverse(arr, maxI)
			res = append(res, maxI+1)
			reverse(arr, k-1)
			res = append(res, k)
		}
		sort(arr, k-1)
	}
	sort(arr, len(arr))
	return res
}
