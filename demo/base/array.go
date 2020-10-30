// 代码内容引用来源：http://www.topgoer.com/go%E5%9F%BA%E7%A1%80/%E6%95%B0%E7%BB%84Array.html
package main

import "fmt"

// 主函数 go run array.go 会执行 main函数
func main()  {

	// 1. 数组：是同一种数据类型的固定长度的序列。
	// 2. 数组定义：var a [len]int，比如：var a [5]int，数组长度必须是常量，且是类型的组成部分。一旦定义，长度不能变。
	// 3. 长度是数组类型的一部分，因此，var a[5] int和var a[10]int是不同的类型。
	// 4. 数组可以通过下标进行访问，下标是从0开始，最后一个元素下标是：len-1
	//  for i := 0; i < len(a); i++ {
	//  }
	//  for index, v := range a {
	//  }
	// 5. 访问越界，如果下标在数组合法范围之外，则触发访问越界，会panic
	// 6. 数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
	// 7. 支持 "=="、"!=" 操作符，因为内存总是被初始化过的。
	// 8. 指针数组 [n]*T，数组指针 *[n]T。

	// 一维数组
	oneArray()	// 知识点 这种没有 return 返回值的方式叫 "裸返回"

	// 多维数组
	muArray()

	// 拷贝和传参
	copyAndTransfer()

	// 指针数组 和 数组指针
	extension()
}

// 一维数组 全局变量
var arr1 [5]int = [5]int{1, 2, 3}
var arr2 = [5]int{1, 2, 3, 4, 5}
var arr3 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "hello-world", 4: "tom"}

// 一维数组 散装英语
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

// 拷贝和传参
func copyAndTransfer()  {

	// 知识点：*代表指针 &代表取址（取内存地址）
	var arr5 [3]int
	fmt.Printf("arr5 内存地址: %p\n", &arr5)

	// 输出数组中的每一个值
	for key, value := range arr5  {
		fmt.Println("key:", key, " value:", value, "内存地址:", &arr5[key])
	}

	// 知识点：数组是值类型，赋值和传参会复制整个数组，而不是指针。因此改变副本的值，不会改变本身的值。
	printArr(arr5)

	// 输出 arr5 查看内容变化
	fmt.Println("copyAndTransfer 方法中的数组：", arr5)

	// go语言的赋值是值拷贝
	arr6 := arr5
	fmt.Printf("arr6 内存地址: %p\n", &arr6)

	// 通过上面一系列的操作我们发现在传参的过程中系统会申请一个新的内存地址 证明了赋值和传参会复制整个数组
}

// 打印数组
func printArr(arr [3]int){

	fmt.Println("\n打印数组------------------------START")
	fmt.Printf("arr  传参内存地址: %p\n", &arr)

	for key, value := range arr  {
		fmt.Println("key:", key, " value:", value, "内存地址:", &arr[key])

		// 修改数组内容
		arr[key] = key
	}

	fmt.Println("printArr 方法中修改后的数组：", arr)
	fmt.Println("打印数组------------------------END\n ")
}

// 指针数组 和 数组指针
func extension()  {

	// 指针数组 [n]*T，数组指针 *[n]T。

	// 定义两个变量
	x,y := 1, 2

	// 定义数组
	var arr =  [...]int{5:2}

	//数组指针
	var pf *[6]int = &arr

	//指针数组
	pfArr := [...]*int{&x,&y}
	fmt.Println(pf)			// 输出  &[0 0 0 0 0 2]
	fmt.Println(pfArr)		// 输出	[0xc000134008 0xc000134010] 两个内容地址

	// 简单理解
	// 数组指针 pf 的值得到的是一个指向 arr 地址的一个指针
	// 指针数组 pfArr 则是一个数组内的元素全是指针类型
	// 关于 指针 相关知识可以在后续文章中学习 这里大致了解一下即可
}