package sort

/*
1.取基准数
从数列中取出一个数作为基准数
2.分区
把大于基准数的放在右边，小于基准数的放在这个数左边
1)	i = L, j = R, 取a[i]为基准数，把这个数挖出来，想象挖出来之后形成一个坑
2)	从j开始从右向左找，找到第一个比基准数a[i]小的数，找到后把这个数填到a[i]那个坑中，a[j]挖出来之后会有一个坑
3)	从i开始从左往右找，找到后把这个数挖出来填到步骤2)中的那个坑中
4)	重复执行2)和3)步，直到i == j，这时应该还有一个坑a[i]，然后将基准数填入a[i]中
3.	递归
对左右区间重复“取基准数”和“分区两个步骤”，递归结束条件：每个区间只有一个数

*/
func quick_sort(array []int, start int, end int) {
	if start < end {
		i := start
		j := end
		base := array[i]

		for i < j {
			for i < j && array[j] >= base {
				j--
			}

			if i < j {
				array[i] = array[j]
				i++
			}

			for i < j && array[i] < base {
				i++
			}

			if i < j {
				array[j] = array[i]
				j--
			}
		}

		array[i] = base
		quick_sort(array, start, i-1)
		quick_sort(array, i+1, end)
	}
}
