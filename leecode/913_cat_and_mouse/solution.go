package _13_cat_and_mouse

/*
913. 猫和老鼠

两位玩家分别扮演猫和老鼠，在一张 无向 图上进行游戏，两人轮流行动。

图的形式是：graph[a] 是一个列表，由满足 ab 是图中的一条边的所有节点 b 组成。

老鼠从节点 1 开始，第一个出发；猫从节点 2 开始，第二个出发。在节点 0 处有一个洞。

在每个玩家的行动中，他们 必须 沿着图中与所在当前位置连通的一条边移动。例如，如果老鼠在节点 1 ，那么它必须移动到 graph[1] 中的任一节点。

此外，猫无法移动到洞中（节点 0）。

然后，游戏在出现以下三种情形之一时结束：

如果猫和老鼠出现在同一个节点，猫获胜。
如果老鼠到达洞中，老鼠获胜。
如果某一位置重复出现（即，玩家的位置和移动顺序都与上一次行动相同），游戏平局。
给你一张图 graph ，并假设两位玩家都都以最佳状态参与游戏：

如果老鼠获胜，则返回 1；
如果猫获胜，则返回 2；
如果平局，则返回 0 。
*/

const (
	draw     = 0
	mouseWin = 1
	catWin   = 2
)

func catMouseGame(graph [][]int) int {
	n := len(graph)
	// dp[mouse][cat][turns]
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2*n)
			for k := range dp[i][j] {
				dp[i][j][k] = -1
			}
		}
	}
	var dfs func(mouse, cat, turns int) int
	dfs = func(mouse, cat, turns int) int {
		if turns == 2*n {
			return draw
		}
		if mouse == 0 {
			return mouseWin
		}
		if cat == mouse {
			return catWin
		}
		if dp[mouse][cat][turns] != -1 {
			return dp[mouse][cat][turns]
		}
		ans := -1
		if turns%2 == 0 { // mouse move
			ans = catWin
			for _, node := range graph[mouse] {
				r := dfs(node, cat, turns+1)
				if r != catWin {
					ans = r
					if ans != draw {
						break
					}
				}
			}
		} else { // cat move
			ans = mouseWin
			for _, node := range graph[cat] {
				if node == 0 {
					continue
				}
				r := dfs(mouse, node, turns+1)
				if r != mouseWin {
					ans = r
					if ans != draw {
						break
					}
				}
			}
		}
		dp[mouse][cat][turns] = ans
		return ans
	}
	return dfs(1, 2, 0)
}
