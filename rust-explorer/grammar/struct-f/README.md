# 结构体

1 结构体

```
struct User {
    username: String,
    email: String,
    sign_in_count: u64,
    active: bool,
}
```

2 元组

```
struct Color(i32, i32, i32);
struct Point(i32, i32, i32);
```

3 类单元结构体（unit-like structs）因为它们类似于 ()，即 unit 类型。

注意整个实例必须是可变的；Rust 并不允许只将某个字段标记为可变。另外需要注意同其他任何表达式一样，我们可以在函数体的最后一个表达式中构造一个结构体的新实例，来隐式地返回这个实例

4 有趣的例子，计算长方形的面积
    4.1 函数直接计算
    4.2 使用元组
    4.3 使用结构体
    4.4 trait

5 方法的语法

方法 与函数类似：它们使用 fn 关键字和名称声明，可以拥有参数和返回值，同时包含在某处调用该方法时会执行的代码。不过方法与函数是不同的，因为它们在结构体的上下文中被定义并且它们第一个参数总是 self，它代表调用该方法的结构体实例。

self 前面加上 &，就像 &Rectangle 一样。方法可以选择获取 self 的所有权，或者像我们这里一样不可变地借用 self，或者可变地借用 self，就跟其他参数一样

如果想要在方法中改变调用方法的实例，需要将第一个参数改为 &mut self。通过仅仅使用 self 作为第一个参数来使方法获取实例的所有权是很少见的；这种技术通常用在当方法将 self 转换成别的实例的时候，这时我们想要防止调用者在转换之后使用原始的实例

```
#[derive(Debug)]
struct Rectangle {
    width: u32,
    height: u32,
}

impl Rectangle {
    fn area(&self) -> u32 {
        self.width * self.height
    }
}
```

关联函数

    impl 块的另一个有用的功能是：允许在 impl 块中定义 不 以 self 作为参数的函数。这被称为 关联函数（associated functions），因为它们与结构体相关联。它们仍是函数而不是方法，因为它们并不作用于一个结构体的实例。
    使用::语法，Strings::from("abc")

多个impl块





