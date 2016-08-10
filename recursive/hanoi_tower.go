package recursive

import "fmt"


/*
有A、B、C三个塔
思路:
如果想把n个盘子从A移动到C
(1)先把n-1个盘子从A移动到B(借助C)
(2)把第n个盘子从A移动到C
(3)把n-1个盘子从B移动到C(借助A)

 */
func hanoi_tower(n int, startTower string, middleTower string, endTower string) {
	if n == 1 {
		fmt.Println(startTower + "---->" + endTower)
	} else {
		hanoi_tower(n - 1, startTower, endTower, middleTower)
		fmt.Println(startTower + "---->" + endTower)
		hanoi_tower(n - 1, middleTower, startTower, endTower)
	}

}
