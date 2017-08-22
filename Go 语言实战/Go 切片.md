# Go 切片

> 切片也是一种数据结构，它和数组非常相似，因为他是围绕动态数组的概念设计的，可以按需自动改变大小，使用这种结构，可以更方便的管理和使用数据集合。

### 内部实现

切片是基于数组实现的，它的底层是数组，它自己本身非常小，可以理解为对底层数组的抽象。切片对象非常小，是因为它是只有3个字段的数据结构：一个是指向底层数组的指针，一个是切片的长度，一个是切片的容量。

### 声明和初始化

```go
// 切片的长度和容量均为5
slice := make([]string, 5)

// 切片长度为3，容量为5
// 切片的“长度<=容量”
slice := make([]string, 3, 5)

// 用字面量声明切片
slice := []string{"a", "b", "c", "d"}
slice := []int{1, 2, 3}

// 使用索引声明切片
slice := []string{99: ""}

// 切片和数组
slice := []int{1, 2, 3}
array := [3]int{1, 2, 3}

// nil切片表示不存在的切片，空切片表示一个空集合
// 创建nil切片
var slice []int
// 创建空切片
slice := make([]int, 0)
```

### 切片使用

```go
// 使用切片创建切片
// 该创建的切片[1:3]包前不包后，也就是说只有两个值
slice := []int{1, 2, 3, 4, 5}
newslice := slice[1:3]

// 修改切片内容导致的结果
// 该程序修改了新切片的第1个值
// 同时也修改了原切片的第2个值
slice := []int{1, 2, 3, 4, 5}
newSlice := slice[1:3]
newSlice[0] = 10
	
fmt.Println(slice)// 输出(1, 10, 3, 4, 5)
fmt.Println(newSlice)//  输出(10, 3)

// append函数可以为一个切片追加一个元素，至于如何增加、返回的是原切片还是一个新切片、长度和容量如何改变这些细节，append函数都会帮我们自动处理

// 使用append向切片增加一个或多个元素
slice := []int{1, 2, 3, 4, 5}
newslice := slice[1:3]
newslice = append(newslice, 6)
newslice = append(newslice, 6, 7, 8)

// 使用append同时增加切片长度和容量
slice := []int{1, 2, 3, 4}
newslice := append(slice, 5)// append操作完，newslice拥有全新的底层数组，容量是原来的两倍

slice := []int{1, 2, 3, 4, 5}
newslice := slice[1:3]
newslice = append(newslice, 10)

fmt.Println(newslice)// 输出(2, 3, 10)
fmt.Println(slice)// 输出(1, 2, 3, 10, 5)
// slice和newslice共用一个底层数组，append操作newslice时同样修改slice的值

// 使用for-range迭代切片
slice := []int{1, 2, 3, 4}
for key, value := range slice{
  fmt.Printf("Key: %d, Value: %d\n", key, value)
}

// 使用传统的for循环进行迭代
for key := 0; key < len(slice); key++{
  fmt.Printf("Key: %d, Value: %d\n", key, slice[key])
} 
/*output
Key: 0, Value: 10
Key: 1, Value: 20
Key: 2, Value: 30
Key: 3, Value: 40
*/

slice := []int{1, 2, 3, 4}
for _, value := range slice{
  fmt.Printf("Value: %d\n", value)
}
/*output
Value: 10
Value: 20
Value: 30
Value: 40
*/
```

### 多维切片

>  与数组多维对应一维数组基本一样

### 函数间传递切片

在函数间传递切片就是要在函数间以值的方式传递切片。由于切片的尺寸很小，在函数间复 制和传递切片成本也很低。在传递复制切片的时候，其底层数组不会被复制，也不会受影响，复制只是复制的切片本身，不涉及底层数组。

```go
func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Printf("%p\n", &slice)
	modify(slice)
	fmt.Println(slice)
}

func modify(slice []int) {
	fmt.Printf("%p\n", &slice)
	slice[1] = 10
}

/* output:
0xc42000a280
0xc42000a2a0
[1 10 3 4 5]
*/
```

两个切片的地址不一样，所以可以确认切片在函数间传递是复制的。而我们修改一个索引的值后，发现原切片的值也被修改了，说明它们共用一个底层数组。

在函数间传递切片非常高效，而且不需要传递指针和处理复杂的语法，只需要复制切片，然后自己修改，最后传递回一个新的切片副本。