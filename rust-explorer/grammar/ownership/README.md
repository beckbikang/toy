所有权是个很关键的概念。

所有权（系统）是 Rust 最为与众不同的特性，它让 Rust 无需垃圾回收（garbage collector）即可保障内存安全。

一些语言中具有垃圾回收机制，在程序运行时不断地寻找不再使用的内存；在另一些语言中，程序员必须亲自分配和释放内存。Rust 则选择了第三种方式：通过所有权系统管理内存，编译器在编译时会根据一系列的规则进行检查。

所有权的规则

    1 Rust 中的每一个值都有一个被称为其 所有者（owner）的变量。
    2 值有且只有一个所有者。
    3 当所有者（变量）离开作用域，这个值将被丢弃。

分配内存，String::from("hello") ，离开作用域drop

## 移动

移动：赋值会导致移动
Rust 永远也不会自动创建数据的 “深拷贝”。因此，任何 自动 的复制可以被认为对运行时性能影响较小

## 克隆

clone

## 栈上数据拷贝

如果一个类型拥有 Copy trait，一个旧的变量在将其赋值给其他变量后仍然可用
Rust 不允许自身或其任何部分实现了 Drop trait 的类型使用 Copy trait

```
所有整数类型，比如 u32。
布尔类型，bool，它的值是 true 和 false。
所有浮点数类型，比如 f64。
字符类型，char。
元组，当且仅当其包含的类型也都是 Copy 的时候。比如，(i32, i32) 是 Copy 的，但 (i32, String) 就不是。
```

## 所有权和函数

1 函数传递参数会拥有所有权

2 返回值也可以转移所有权


## 引用与借用

### 引用的规则

1 在任意给定时间，要么 只能有一个可变引用，要么 只能有多个不可变引用。

2 引用必须总是有效的。



fn(xx :&String)

&获取使用值但不获取其所有权， *解除引用

将获取引用作为函数参数称为 借用（borrowing）。 

引用分为：
    不可变引用 &String
    可变引用   &mut String

不过可变引用有一个很大的限制：在特定作用域中的特定数据有且只有一个可变引用。

优点
```
1 两个或更多指针同时访问同一数据。
2 至少有一个指针被用来写入数据。
3 没有同步数据访问的机制。
```

不能在拥有不可变引用的同时拥有可变引用


一个引用的作用域从声明的地方开始一直持续到最后一次使用为止。


悬垂指针,编译器自动检测


## slice

另一个没有所有权的数据类型是 slice。slice 允许你引用集合中一段连续的元素序列，而不用引用整个集合



可以使用一个由中括号中的 [starting_index..ending_index] 指定的 range 创建一个 slice，其中 starting_index 是 slice 的第一个位置，ending_index 则是 slice 最后一个位置的后一个值。

"字符串 slice" 的类型声明写作 &str

字符串字面值是slice















