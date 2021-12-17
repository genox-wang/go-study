package _518_water_bottles

/*
1518. 换酒问题

小区便利店正在促销，用 numExchange 个空酒瓶可以兑换一瓶新酒。你购入了 numBottles 瓶酒。

如果喝掉了酒瓶中的酒，那么酒瓶就会变成空的。

请你计算 最多 能喝到多少瓶酒。

*/

func numWaterBottles(numBottles int, numExchange int) (ans int) {
	ans = numBottles
	for numBottles >= numExchange {
		ans += numBottles / numExchange
		numBottles = numBottles/numExchange + numBottles%numExchange
	}
	return ans
}
