package main

import "fmt"

func main() {
	//切片的扩容机制
	//在Go语言中，当你向切片添加元素，但切片的容量不足以容纳所有元素时，切片会自动扩容。
	//扩容的规则是：如果新容量小于256个元素，则扩大为原来的两倍；如果新容量大于256个元素，则扩大为原来的1.25倍

	num := make([]int, 0, 3)
	fmt.Printf("地址%p，长度%d，容量%d\n", num, len(num), cap(num))

	num = append(num, 1, 2)
	fmt.Printf("地址%p，长度%d，容量%d\n", num, len(num), cap(num))

	num = append(num, 3, 4, 5, 6)
	fmt.Printf("地址%p，长度%d，容量%d\n", num, len(num), cap(num))

	//每个切片都引用了一个底层数组
	//切片本身不能存储任何数据，都是底层数组存储数据，所以修改切片的时候修改的是地底层数组中的数据
	//如果新容量小于256个元素，则扩大为原来的两倍；如果新容量大于256个元素，则扩大为原来的1.25倍
	//切片一旦扩容，指向一个新的底层数组内存地址也就随之改变

}