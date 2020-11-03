package main

import "fmt"

func main(){

	// 例子一
	example1()

	// 例子二
	example2()

	// 例子三
	example3()

}

// 例子一
func example1(){

	// TODO 数组 Array
	var arr = [5]int{1, 2, 3, 4, 5}
	// 通过输出我们可以看到 因为数据是固定长度 所以 arr的 len=cap=5
	fmt.Printf("arr %v len=%v cap=%v\n", arr, len(arr), cap(arr))

	// 及时只初始化 3个值 数组的 len 和 cap 依然相等
	var arr1 = [5]int{1, 3, 5}
	fmt.Printf("arr1 %v len=%v cap=%v\n", arr1, len(arr1), cap(arr1))

	// TODO 切片 Slice
	var data = [...]int{1, 2, 3, 4, 5, 6, 7, 10:11}
	var slice = data[:3]
	// 通过输出可以看到 slice的 len=3 cap=11
	// 之前说过 切片是数组的一个引用 所以
	fmt.Printf("slice %v len=%v cap=%v\n", slice, len(slice), cap(slice))
}

// 例子二
func example2(){

	var team = [10]int {1,2,3,4,5,6,7,8,9,10}
	var teamLeader = team[:]

	teamLeader[0] = 100

	fmt.Println(teamLeader)
}

// 例子三
func example3(){

	var team = [10]int {1,2,3,4,5,6,7,8,9,10}
	var teamLeader = team[:]

	teamLeader[0] = 100
	fmt.Printf("teamLeader %p %v len=%v cap=%v\n", teamLeader, teamLeader, len(teamLeader), cap(teamLeader))

	teamLeader = append(teamLeader, 111, 222, 333)
	// 追加新内容后 如果超过cap限制 切片会重新分配底层数组 cap会在原有基础上乘2
	// 内存切片内存地址也会改变
	fmt.Printf("teamLeader %p %v len=%v cap=%v\n", teamLeader, teamLeader, len(teamLeader), cap(teamLeader))

}
