# Go Map

### 内部实现

Map是给予散列表来实现，就是常说的Hash表，所以我们每次迭代Map的时候，打印的Key和Value是无序的，每次迭代的都不一样，即使我们按照一定的顺序存在也不行。

Map的散列表包含一组桶，在存储、删除或者查找键值对的时候，所有操作都要先选择一个 桶。把操作映射时指定的键传给映射的散列函数，就能选中对应的桶。

这种方式的好处在于，存储的数据越多，索引分布越均匀，所以我们访问键值对的速度也就越快，这里我们只要记住**Map存储的是无序的键值对集合**。

### 创建&初始化

```go
// 使用make声明
dict := make(map[string]int)

// 使用键值对初始化
dict := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}

// 使用字符串切片作为映射
dict := map[[]string]int{}

// 使用字符串切片作为值
dict := map[int][]string{}
```

### 使用

```go
// 为映射赋值
dict := make(map[string]int)
dict["张三"] = 43

// 获取map值
age := dict["张三"]

// 判断键是否存在
age, exists := dict["李四"]
if exists {
  fmt.Println(age)
}

age := dict["李四"]
if value != "" {
  fmt.Println(age)
}

// 删除键值
delete(dict,"张三")

// 使用range遍历
dict := map[string]int{"张三": 43}
for key, value := range dict {
	fmt.Println(key, value)
}

```

### 在函数间传递映射

函数间传递Map是不会拷贝一个该Map的副本的，也就是说如果一个Map传递给一个函数，该函数对这个Map做了修改，那么这个Map的所有引用，都会感知到这个修改。

```go
func main() {
  dict := map[string]int{"王五": 60, "张三": 43}
  modify(dict)
  fmt.Println(dict["张三"])
}

func modify(dict map[string]int) {
  dict["张三"] = 10
}
```

上面这个例子输出的结果是`10`,也就是说已经被函数给修改了，可以证明传递的并不是一个Map的副本。这个特性和切片是类似的，这样就会更高，因为复制整个Map的代价太大了。