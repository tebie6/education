// 代码内容引用来源：http://www.topgoer.com/go%E5%9F%BA%E7%A1%80/%E6%95%B0%E7%BB%84Array.html
package main

import (
	"fmt"
	"strings"
)

func main(){

	// 需要说明，slice 并不是数组或数组指针。它通过内部指针和相关属性引用数组片段，以实现变长方案。

	// 1. 切片：切片是数组的一个引用，因此切片是引用类型。但自身是结构体，值拷贝传递。
	// 2. 切片的长度可以改变，因此，切片是一个可变的数组。
	// 3. 切片遍历方式和数组一样，可以用len()求长度。表示可用元素数量，读写操作不能超过该限制。
	// 4. cap可以求出slice最大扩张容量，不能超出数组限制。0 <= len(slice) <= len(array)，其中array是slice引用的数组。
	// 5. 切片的定义：var 变量名 []类型，比如 举例1：var str []string  举例2：var arr []int。
	// 6. 如果 slice == nil，那么 len、cap 结果都等于 0。

	//  1.1.1. 创建切片的各种方式
	section1()

	// 1.1.2. 切片初始化
	section2()

	// 1.1.3. 通过make来创建切片
	section3()

	// 1.1.4. 用append内置函数操作切片（切片追加）
	section4()

	// 1.1.5. 超出原 slice.cap 限制，就会重新分配底层数组，即便原数组并未填满。
	section5()

	// 1.1.6. slice中cap重新分配规律：
	section6()

	// 1.1.7. 切片拷贝
	section7()

	// 1.1.8. slice遍历：
	section8()

	// 1.1.9. 切片resize（调整大小）
	section9()

	// 1.1.10. 字符串和切片（string and slice）
	section10()

}




// 1.1.1. 创建切片的各种方式
func section1(){

	// 1、声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("切片为空")
	} else {
		fmt.Println("不是空")
	}

	// 2、 :=
	s2 := []int{}

	// 3、make()
	var s3 = make([]int, 0)
	fmt.Println(s1, s2, s3)

	// 4、初始化赋值
	var s4 = make([]int, 0, 0)
	fmt.Println(s4)

	s5 := []int{1, 2 ,3}
	fmt.Println(s5)

	// 5、从数组切片
	arr := [5]int{1 ,2 ,3, 4, 5}
	var s6  = arr[1:4]	// // 前包后不包 1:4 只包含 下标 1-3 的的值

	fmt.Println(arr, s6)
}

// 1.1.2. 切片初始化

// 全局
var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
var slice0 []int = arr[2:8]
var slice1 []int = arr[0:6]          // 可以简写为 var slice []int = arr[:end] 例子: slice1 []int = arr[:6]
var slice2 []int = arr[5:10]         // 可以简写为 var slice[]int = arr[start:] 例子: slice2 []int = arr[5:]
var slice3 []int = arr[0:len(arr)]	 // var slice []int = arr[:]
var slice4 []int = arr[:len(arr)-1]	 // 去掉切片的最后一个元素

func section2(){

	// 图片资料 该文件同级 slice.jpg 文件

	// 初始化内容 全局 或 局部
	// 全局 就用 var 局部可以使用 :=
	fmt.Printf("全局变量 arr %v\n", arr)
	fmt.Printf("全局变量 slice0 %v\n", slice0)
	fmt.Printf("全局变量 slice1 %v\n", slice1)
	fmt.Printf("全局变量：slice2 %v\n", slice2)
	fmt.Printf("全局变量：slice3 %v\n", slice3)
	fmt.Printf("全局变量：slice4 %v\n", slice4)
	fmt.Printf("-----------------------------------\n")

	// 局部
	arr2 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	// 这里复习一下数组知识点 数组：是同一种数据类型的固定长度的序列。
	// 所以实际从 下标0 开始内容应该是 0,1,2,34,5,6,7,8,9 这样的顺序
	for key, value := range arr2  {
		fmt.Printf("key:%v  value:%v \n", key, value)
	}

	slice5 := arr[2:8]
	slice6 := arr[:6]
	slice7 := arr[5:]
	slice8 := arr[:]
	slice9 := arr[:len(arr)-1] //去掉切片的最后一个元素

	fmt.Printf("局部变量： arr2 %v\n", arr2)
	fmt.Printf("局部变量： slice5 %v\n", slice5)
	fmt.Printf("局部变量： slice6 %v\n", slice6)
	fmt.Printf("局部变量： slice7 %v\n", slice7)
	fmt.Printf("局部变量： slice8 %v\n", slice8)
	fmt.Printf("局部变量： slice9 %v\n", slice9)
}

// 1.1.3. 通过make来创建切片
func section3()  {

	var slice0 []int = make([]int, 10)
	var slice1 = make([]int, 10)
	var slice2 = make([]int, 10, 10)

	fmt.Printf("slice0: %v\n", slice0)
	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice2: %v\n", slice2)


	// 因为 切片只是数组的一个引用，如果原数组的数据发生变化，那么会连带着切片中的数据一起变化。
	// 例子一
	data := [...]int{0, 1, 2, 3, 4, 5}
	s := data[2:4]
	s[0] += 100
	s[1] += 200

	fmt.Println(s)
	fmt.Println(data)

	// 例子二
	data1 := [...]int{0, 1, 2, 3, 4, 5}
	s1 := data1[2:4]
	data1[2] = 100

	fmt.Println(s1)
	fmt.Println(data1)

	// 可直接创建 slice 对象，自动分配底层数组。
	s2 := make([]int, 6, 8) // 使用 make 创建，指定 len 和 cap 值。
	fmt.Println(s2, len(s2), cap(s2))

	s3 := make([]int, 6) // 省略 cap，相当于 cap = len。
	fmt.Println(s3, len(s3), cap(s3))

	s4 := []int{0, 1, 2, 3, 8: 100} // 通过初始化表达式构造，可使用索引号。
	fmt.Println(s4, len(s4), cap(s4))


	// 使用 make 动态创建slice，避免了数组必须用常量做长度的麻烦。还可用指针直接访问底层数组，退化成普通数组操作。
	s6 := []int{0, 1, 2, 3}
	p := &s6[2] // *int, 获取底层数组元素指针。
	*p += 100

	fmt.Println(s6)

	// 至于 [][]T，是指元素类型为 []T 。
	data2 := [][]int{
		[]int{1, 2, 3},
		[]int{100, 200},
		[]int{11, 22, 33, 44},
	}
	fmt.Println(data2)


	// 可直接修改 struct array/slice 成员。
	d := [5]struct{
		x int
	}{}

	s7 := d[:]

	d[1].x = 10
	s7[2].x = 20

	fmt.Println(d)
	fmt.Printf("%p, %p\n", &d, &d[0])
	// TODO 为什么 &d 和 &d[0] 输出的内存地址是一样的？
	// 因为 &d为指针类型，指的是数组首地址(即就是数组首元素的地址)
	// d 和 &d 不同类型，但数值是相同的原因是 d和 d[0]的地址相同。
}


// 1.1.4. 用append内置函数操作切片（切片追加）
func section4(){

	var a = []int{1, 2, 3}
	fmt.Printf("slice a : %v\n", a)

	var b = []int{4, 5, 6}
	fmt.Printf("slice b : %v\n", b)

	c := append(a, b...)	// 最后这三个点点点 ... 要加上，不然会报错，这是用于两个切片的合并的。
	fmt.Printf("slice c : %v\n", c)

	d := append(c, 7)
	fmt.Printf("slice d : %v\n", d)

	e := append(d, 8, 9, 10)
	fmt.Printf("slice e : %v\n", e)

	// append ：向 slice 尾部添加数据，返回新的 slice 对象。
	s1 := make([]int, 0 ,5)
	fmt.Printf("%p\n", &s1)

	s2 := append(s1, 1)
	fmt.Printf("%p\n", &s2)
	fmt.Println(s1, s2)
}

// 1.1.5. 超出原 slice.cap 限制，就会重新分配底层数组，即便原数组并未填满。
func section5()  {

	data := [...]int{0, 1, 2, 3, 4, 10:0}

	// 这个操作小白可能就懵逼了 我来解释一下
	// 其实就是 slice.jpg 图片里 s[low:high:max] 格式 等于data[:2] 取最大容量（cap）为3
	s := data[:2:3]

	// 可以通过注释该行代码 来看append前后的内存地址的不同
	s = append(s, 100 ,200)	  // 一次 append 两个值，超出 s.cap 限制。

	fmt.Println(s, data)
	fmt.Println(&s[0], &data[0])


	// 从输出结果可以看出，append 后的 s 重新分配了底层数组，并复制数据。
	// 如果只追加一个值，则不会超过 s.cap 限制，也就不会重新分配。
	// 通常以 2 倍容量重新分配底层数组。在大批量添加数据时，建议一次性分配足够大的空间，
	// 以减少内存分配和数据复制开销。或初始化足够长的 len 属性，改用索引号进行操作。
	// 及时释放不再使用的 slice 对象，避免持有过期数组，造成 GC 无法回收。
}

// 1.1.6. slice中cap重新分配规律：
func section6(){

	// 该方法中的代码主要是为了展示 每当slice重新分配底层数组后 cap的变化
	// cap数 成倍增加
	s := make([]int, 0 ,1)
	c := cap(s)

	for i := 0; i <= 50; i++  {
		s = append(s, i)

		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}
}

// 1.1.7. 切片拷贝
func section7(){

	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice s1 : %v\n", s1)
	s2 := make([]int, 10)
	fmt.Printf("slice s2 : %v\n", s2)

	copy(s1, s2)
	fmt.Printf("copied slice s1 : %v\n", s1)
	fmt.Printf("copied slice s2 : %v\n", s2)

	s3 := []int{1, 2, 3}
	fmt.Printf("slice s3 : %v\n", s3)
	s3 = append(s3, s2...)
	fmt.Printf("appended slice s3 : %v\n", s3)
	s3 = append(s3, 4, 5, 6)
	fmt.Printf("last slice s3 : %v\n", s3)


	// copy ：函数 copy 在两个 slice 间复制数据，复制长度以 len 小的为准。两个 slice 可指向同一底层数组，允许元素区间重叠。
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("array data : ", data)
	s6 := data[8:]
	s7 := data[:5]
	fmt.Printf("slice s6 : %v\n", s6)
	fmt.Printf("slice s7 : %v\n", s7)

	//copy(s7, s6)
	copy(s6, s7)
	fmt.Printf("copied slice s6 : %v\n", s6)
	fmt.Printf("copied slice s7 : %v\n", s7)
	fmt.Println("last array data : ", data)

	// 以下为输出内容：
	// array data :  [0 1 2 3 4 5 6 7 8 9]
	// slice s6 : [8 9]
	// slice s7 : [0 1 2 3 4]
	// copied slice s6 : [8 9]
	// copied slice s7 : [8 9 2 3 4]
	// last array data :  [8 9 2 3 4 5 6 7 8 9]

	// 解释：
	// 上面说了 在两个 slice 间复制数据，复制长度以 len 小的为准
	// 也就是说以 s6的长度为准 2位
	// 将 s6 内容复制给 s7 所以会将 s7 的前两位复制成 s6 的内容
	// 反之如果将 s7 复制给 s6还是会以长度最短的 s6长度 为复制长度
	// 取 s7 前两位复制给 s6 前两位

	// 注意：应及时将所需数据 copy 到较小的 slice，以便释放超大号底层数组内存。
}

// 1.1.8. slice遍历：
func section8(){

	// 内容很简单 使用 for range
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := data[:]
	for index, value := range slice {
		fmt.Printf("inde : %v , value : %v\n", index, value)
	}
}

// 1.1.9. 切片resize（调整大小）
func section9(){

	// 该小结内容也很简单易懂 就是使用 slice.jpg 中的操作方式 来对 切片 进行 切片
	// 得到的结果是切片后的切片 len可能会发生改变 cap可能也会发生改变
	var a = []int{1, 3, 4, 5}
	fmt.Printf("slice a : %v , len(a) : %v\n", a, len(a))
	b := a[1:2]
	fmt.Printf("slice b : %v , len(b) : %v\n", b, len(b))
	c := b[0:3]
	fmt.Printf("slice c : %v , len(c) : %v\n", c, len(c))

	// 这里 b的cap 变成了3 是因为 切片cap是通过 开始切片的位置计算的
	// 如果是 a[0:2] 那么cap 就会变成4
	// 切片c 的cap 也是3 是因为 他是从 b切片 上 切片 的 切片 （有点绕 切片这个词 我个人认为可以是一个"动词" 也可以是一个"名词"）
	fmt.Printf("slice a : %v , cap(a) : %v\n", a, cap(a))
	fmt.Printf("slice b : %v , cap(b) : %v\n", b, cap(b))
	fmt.Printf("slice c : %v , cap(c) : %v\n", c, cap(c))
}


// 1.1.10. 字符串和切片（string and slice）
func section10(){

	// 很像 PHP中的 substr() 函数
	str := "hello world"
	s1 := str[0:5]
	fmt.Println(s1)

	s2 := str[6:]
	fmt.Println(s2)


	// string本身是不可变的，因此要改变string中字符。需要如下操作： 英文字符串：
	str1 := "Hello world"
	s := []byte(str1) //中文字符需要用[]rune(str)
	s[6] = 'G'
	s = s[:8]
	s = append(s, '!')
	str = string(s)
	fmt.Println(str)

	// 中文测试
	str2 := "攻击敌方水晶"
	s3 := []rune(str2) //中文字符需要用[]rune(str)
	s3[2] = '我'
	// string 是为了编码 如果不使用输出内容则是一个切片：[25915 20987 25932 26041 27700 26230]
	// string 整形字符串输出为unicode码点的utf8字符串
	// 这里可以理解成 PHP中的 implode() 函数
	str2 = string(s3)
	fmt.Println(str2)

	// 中文例子二
	str3 := "你好，世界！hello world！"
	s6 := []rune(str3)
	s6[3] = '够'
	s6[4] = '浪'
	s6[12] = 'g'
	s6 = s6[:14]
	str3 = string(s6)
	fmt.Println(str3)

	// golang slice data[:6:8] 两个冒号的理解
	// 常规slice , data[6:8]，从第6位到第8位（返回6， 7），长度len为2， 最大可扩充长度cap为4（6-9）
	// 另一种写法： data[:6:8] 每个数字前都有个冒号， slice内容为data从0到第6位，长度len为6，最大扩充项cap设置为8
	// a[x:y:z] 切片内容 [x:y] 切片长度: y-x 切片容量:z-x

	// 数组or切片转字符串：
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	array := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	str5 := strings.Replace(strings.Trim(fmt.Sprint(slice), "[]"), " ", ",", -1)
	str6 := strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
	fmt.Println(str5)
	fmt.Println(str6)

}