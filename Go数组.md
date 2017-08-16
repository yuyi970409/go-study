# Go数组

### 内部实现

在 Go 语言里，数组是一个长度固定的数据类型，用于存储一段具有相同的类型的元素的连 续块。数组存储的类型可以是内置类型，如整型或者字符串，也可以是某种结构类型。因为是连续的，所以索引比较好计算，可以很快得到数组中的任何数据。

### 声明和初始化

1. 声明数组并初始化数组

```go
var array [5]int
array = [5]int{1, 2, 3, 4, 5}
```

2. 使用 ":="直接初始化数组

```go
array := [5]int{1, 2, 3, 4, 5}
```

3. Go自动计算数组程度

```go
array := [...]int{1, 2, 3, 4, 5}
```

4. 声明并指定特定元素值

```go
array := [5]int{1: 3, 2: 5}
```

### 使用数组

1. 索引访问

```go
fmt.Printf("%d", array[1])
```

2. 修改元素

 ```go
array := [5]int{1: 3; 4: 5}
fmt.Printf{"%d\n", array[1]}

array[1]=1
fmt.Printf{"%d\n", array[1]}
 ```

3. 访问指针数组的元素

```go
array := [5]*int{0: new(int), 1: new(int)}

*array[0] = 1
*array[1] = 2
```

4. 同一类型的数组复制给另一数组

```go
var array1 [5]string

array2 := [5]string{"a", "b", "c", "d", "e"}
array1 = array2
```

5. 指针数组赋值给另一个

```go
var array1 [3]*string
array2 := [3]*string{new(string), new(string), new(string)}

*array2[0] = "a"
*array2[1] = "b"
*array2[2] = "c"

array1 = array2
```

6. 使用for-range遍历

```go
func main() {
	array := [5]int{1: 1, 3: 4}
	
	for i, v := range array {
	    fmt.Printf("索引:%d,值:%d\n", i, v)
  }
}
```



### 多维数组

声明&初始化&指定特定元素值&访问二维数组

```go
var array [4][2]int

array := [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}

array := [4][2]int{1: {20, 21}, 3: {40, 41}}

// 声明初始化外层和内层数组
array := [4][2]int{1: {0: 20}, 3: {1: 41}}

array[0][0] = 1
array[0][1] = 2
...
array[3][1] = 8
```

### 函数间传递数组

在函数间传递变量时，总是以值的方式，如果变量是个数组，那么就会整个复制，并传递给函数，如果数组非常大，比如长度100多万，那么这对内存是一个很大的开销。

```go
func main() {
	array := [5]int{1: 2, 3: 4}
	modify(array)

	fmt.Println(array)
}

func modify(a [5]int){
	a[1] = 3
	fmt.Println(a)
}
```

通过上面的例子，可以看到，数组是复制的，原来的数组没有修改。我们这里是5个长度的数组还好，如果有几百万怎么办，有一种办法是传递数组的指针，这样，复制的大小只是一个数组类型的指针大小。

```go
func main() {
	array := [5]int{1: 2, 3: 4}
	modify(&array)
	fmt.Println(array)
}

func modify(a *[5]int){
	a[1] = 3
	fmt.Println(*a)
}
```

这是传递数组的指针的例子，会发现数组被修改了。所以这种情况虽然节省了复制的内存，但是要谨慎使用，因为一不小心，就会修改原数组，导致不必要的问题。





