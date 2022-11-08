package bfs

// 力扣第 752 题
// 打开转盘锁

// https://leetcode.cn/problems/open-the-lock/description/

// https://labuladong.github.io/algo/1/10/#三解开密码锁的最少次数

func OpenLock(deadends []string, target string) int {
	// 记录需要跳过的死亡密码
	deadendsSet := make(map[string]struct{}, len(deadends))
	for _, dead := range deadends {
		deadendsSet[dead] = struct{}{}
	}
	// 记录已经穷举过的密码，防止走回头路
	visited := make(map[string]struct{})
	// 从起点开始启动广度优先搜索
	times := 0
	queue := make([]string, 0)
	origin := "0000"
	queue = append(queue, origin)
	visited[origin] = struct{}{}
	for len(queue) != 0 {
		// 将当前队列中的所有节点向周围扩散
		for _, item := range queue {
			queue = queue[1:]
			// 判断是否到达终点
			if item == target {
				return times
			}
			if _, ok := deadendsSet[item]; ok {
				continue
			}
			// 将一个节点的未遍历相邻节点加入队列
			for i := 0; i < len(target); i++ {
				puls := plusOne(item, i)
				minus := minusOne(item, i)
				if _, ok := visited[puls]; !ok {
					visited[puls] = struct{}{}
					queue = append(queue, puls)
				}
				if _, ok := visited[minus]; !ok {
					visited[minus] = struct{}{}
					queue = append(queue, minus)
				}
			}
		}
		// 在这里增加步数
		times++
	}
	// 如果穷举完都没找到目标密码，那就是找不到了
	return -1
}

// 将 taget[index] 向上拨动一次
func plusOne(taget string, index int) string {
	value := []byte(taget)
	if value[index] == '9' {
		value[index] = '0'
	} else {
		value[index] += 1
	}
	return string(value)
}

// 将 taget[index] 向下拨动一次
func minusOne(taget string, index int) string {
	value := []byte(taget)
	if value[index] == '0' {
		value[index] = '9'
	} else {
		value[index] -= 1
	}
	return string(value)
}
