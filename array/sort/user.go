package main

import (
	"fmt"
	"sort"
)

type User struct {
	UserName string
	Age int
}

// user数组根据age升序
type UserAsc []User

// 获取长度
func (u UserAsc) Len() int {
	return len(u)
}

// 排序判断
func (u UserAsc) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}
// 位置调换
func (u UserAsc) Swap(i, j int) {
	u[i],u[j] = u[j],u[i]
}

// 自定义结构体排序
func sturtSort()  {
	var userList = UserAsc{}
	for i:=20;i>0;i--{
		var userInfo = User{
			UserName:fmt.Sprintf("user%d",i),
			Age:i,
		}
		userList = append(userList,userInfo)

	}

	sort.Sort(userList)
	fmt.Println("user数组根据age升序:",userList)
	// Reverse 逆向操作
	sort.Sort(sort.Reverse(userList))
	fmt.Println("user数组根据age降序:",userList)

}

func sortMethod()  {
	var intSlice = []int{0, 33, 20, -23, 1, 40}

	sort.Sort(sort.Reverse(sort.IntSlice(intSlice)))
	fmt.Println(intSlice)

	var float64Slice = []float64{1.2, 4.2, -2.2, 8.8, 5.8}

	sort.Sort(sort.Reverse(sort.Float64Slice(float64Slice)))
	fmt.Println(float64Slice)

	var stringSlice = []string{"hello", "golang", "world", "bar", "foo"}

	sort.Sort(sort.Reverse(sort.StringSlice(stringSlice)))
	fmt.Println(stringSlice)
}

func main()  {
	// 自定义结构体排序
	sturtSort()
	// 自带排序方法
	sortMethod()

}



