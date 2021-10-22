package _29_majority_element_2

/**
229. 求众数 II

给定一个大小为 n 的整数数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
*/

func majorityElement(nums []int) (ans []int) {
	num1, vote1, num2, vote2 := 0, 0, 0, 0
	for _, num := range nums {
		if vote1 == 0 && num != num2 {
			num1 = num
			vote1++
		} else if vote2 == 0 && num != num1 {
			num2 = num
			vote2++
		} else if num1 == num {
			vote1++
		} else if num2 == num {
			vote2++
		} else {
			vote1--
			vote2--
		}
	}
	cnt1, cnt2 := 0, 0
	for _, num := range nums {
		if vote1 > 0 && num == num1 {
			cnt1++
		}
		if vote2 > 0 && num == num2 {
			cnt2++
		}
	}
	if cnt1 > len(nums)/3 {
		ans = append(ans, num1)
	}
	if cnt2 > len(nums)/3 {
		ans = append(ans, num2)
	}
	return
}
