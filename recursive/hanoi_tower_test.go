package recursive

import (
	"testing"
)


/*
the correct output
A---->C
A---->B
C---->B
A---->C
B---->A
B---->C
A---->C
 */
func TestHanoi_tower(t *testing.T) {
	startTower := "A"
	middleTower := "B"
	endTower := "C"
	n := 3
	hanoi_tower(n, startTower, middleTower, endTower)
}
