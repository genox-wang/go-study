package _109_corporate_flight_bookings

/**
1109. 航班预订统计

这里有 n 个航班，它们分别从 1 到 n 进行编号。

有一份航班预订表 bookings ，表中第 i 条预订记录 bookings[i] = [firsti, lasti, seatsi] 意味着在从 firsti 到 lasti （包含 firsti 和 lasti ）的 每个航班 上预订了 seatsi 个座位。

请你返回一个长度为 n 的数组 answer，里面的元素是每个航班预定的座位总数。
*/

func corpFlightBookings(bookings [][]int, n int) []int {
	seats := make([]int, n)
	for _, b := range bookings {
		seat := b[2]
		seats[b[0]-1] += seat
		if b[1] < n {
			seats[b[1]] -= seat
		}
	}
	for i := 1; i < n; i++ {
		seats[i] += seats[i-1]
	}
	return seats
}
