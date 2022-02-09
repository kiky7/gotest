/*
 * @author:kiky
 * @date: 2022/1/3 10:20 上午
**/

package series

import "fmt"

func init()  {
	fmt.Println("init1")
}

func init()  {
	fmt.Println("init2")
}



func GetFibonacciSeries(n int) ([]int)  {

	fibList := []int{1,1}

	for i := 2;i<n;i++ {
		fibList = append(fibList,fibList[i-2]+fibList[i-1])
	}
	return fibList
}

func Sequare(n int) int{
	return n*n
}

