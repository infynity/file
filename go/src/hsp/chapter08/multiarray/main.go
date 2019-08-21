package main
import (
	"fmt"
)

func main() {
	//二维数组的演示案例
	/*
	0 0 0 0 0 0
	0 0 1 0 0 0
	0 2 0 3 0 0
	0 0 0 0 0 0
	*/

	//定义/声明二维数组
	var arr332 [4][4]int
	fmt.Println(arr332)
	var arr [4][6]int
	//赋初值
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	//遍历二维数组，按照要求输出图形
	for i := 0; i < 4; i++ {
		for j := 0; j < 6; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}
	
	fmt.Println()
	
	var arr2 [2][10]int8 //以这个为例来分析arr2在内存的布局!!
	arr2[1][1] = 10
	fmt.Println(arr2)
	fmt.Printf("arr2[111111111]的地址%p\n %p\n %p\n",&arr2,&arr2[0],&arr2[1])

	var test int =323
	fmt.Println(&test)

	fmt.Printf("arr2[0]的地址%p\n", &arr2[0])
	fmt.Printf("arr2[1]的地址%p\n", &arr2[1])

	fmt.Printf("arr2[0][0]的地址%p\n", &arr2[0][0])
	fmt.Printf("arr2[0][0]的地址%p\n", &arr2[0][1])
	fmt.Printf("arr2[0][0]的地址%p\n", &arr2[0][2])
	fmt.Printf("arr2[0][0]的地址%p\n", &arr2[0][3])
	fmt.Printf("arr2[0][0]的地址%p\n", &arr2[0][4])
	fmt.Printf("arr2[0][0]的地址%p\n", &arr2[0][5])
	fmt.Printf("arr2[1][0]的地址%p\n", &arr2[1][0])

	fmt.Println()
	arr3  := [2][3]int8{{1,2,3}, {4,5,127}}  //128以上就超出了    -128~127
	fmt.Println("arr3=", arr3)

}