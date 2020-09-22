// 代码内容引用来源：http://www.topgoer.com/go%E5%9F%BA%E7%A1%80/%E6%95%B0%E7%BB%84Array.html
package main

import "fmt"

// 全局定义数组
var arr1 [5]int = [5]int{1, 2, 3}
var arr2 = [5]int{1, 2, 3, 4, 5}
var arr3 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "hello-world", 4: "tom"}

// 主函数 go run array.go 会执行 main函数
func main()  {

	// 一纬数组
	oneArray()	// 知识点 这种没有 return 返回值的方式叫 "裸返回"

	// 多维数组
	muArray()
}

// 一纬数组 散装英语
func oneArray(){

	fmt.Println("一维数组------------------------START")

	// 如果 数组定义了长度在位数不足的时候会自动用 0 补位 就是所谓的 【未初始化元素值为 0。】
	fmt.Println(arr1)	// 输出 [1 2 3 0 0]
	fmt.Println(arr2)	// 输出 [1 2 3 4 5]

	// 在不确定数据长度时可以通过 [...] 来初始化值确定数组长度。
	fmt.Println(arr3)	// 输出 [1 2 3 4 5 6]

	// 返回值中 "[   hel" 这里的空位 是因为字符串类型通过 空格占位 ["索引号0占位" "索引号1占位" "索引号2站位" hello-world tom]
	fmt.Println(str)	// 输出 [   hello-world tom]


	// 局部定义数组
	a := [3]int{1, 2}           // 未初始化元素值为 0。
	b := [...]int{1, 2, 3, 4}   // 通过初始化值确定数组长度。
	c := [5]int{2: 100, 4: 200} // 使用索引号初始化元素。

	// 数组定义：var a [长度]类型，比如：var a [5]int， 这里的 struct数据结构体可以理解成PHP中的无序数组
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10}, // 可省略元素类型。
		{"user2", 20}, // 别忘了最后一行的逗号。
	}

	fmt.Println(a, b, c, d)	// 输出 [1 2 0] [1 2 3 4] [0 0 100 0 200] [{user1 10} {user2 20}]
	fmt.Println("一维数组------------------------END")

}

// 多维数组 全局变量
var muarr1 [5][3]int
var muarr2 [2][3]int = [...][3]int{{1, 2, 3}, {7, 8, 9}}

// 多维数组
func muArray()  {

	// 多维数组 局部变量
	muarr3 := [5][3]int{{1, 2, 3}, {4, 5, 6}}
	muarr4 := [...][2]int{{1, 1}, {2, 2}, {3, 3}}	// 第 2 纬度不能用 "..."
	muarr5 := [1][2][3]int{
		0:{
			0:{3, 3, 3},
			1:{3, 3, 3},
		},
	}	// 通过例子你应该可以扩展出 其他组合的数组了

	fmt.Println("多维数组------------------------START")
	fmt.Println(muarr1)		// 输出 [[0 0 0] [0 0 0] [0 0 0] [0 0 0] [0 0 0]]
	fmt.Println(muarr2)		// 输出 [[1 2 3] [7 8 9]]
	fmt.Println(muarr3)		// 输出 [[1 2 3] [4 5 6] [0 0 0] [0 0 0] [0 0 0]]
	fmt.Println(muarr4)		// 输出 [[1 1] [2 2] [3 3]]
	fmt.Println(muarr5)		// 输出 [[[3 3 3] [3 3 3]]]
	fmt.Println("多维数组------------------------END")

}