package problem

import "math"

/*
素数的定义:在大约1的自然数中,除了1和该数本身,无法被其他自然数整除。
最小的质数是:2
 */


/*
试除法判断一个数是否是素数:从2到math.Sqrt2(num)依次整除num,如果能整除,则num不是素数;
如果都不能整除,则num是素数
 */
func IsPrime(num int64) bool {
	// 素数必须是大于1的自然数
	if num < 2 {
		return false
	}

	boundary := int64(math.Floor(math.Sqrt(float64(num))))

	for i := int64(2); i < boundary; i++ {
		if num % boundary == 0 {
			return false
		}
	}
	return true
}
