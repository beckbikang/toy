rust的模块系统

```
包（Packages）： Cargo 的一个功能，它允许你构建、测试和分享 crate。
Crates ：一个模块的树形结构，它形成了库或二进制项目。
模块（Modules）和 use： 允许你控制作用域和路径的私有性。
路径（path）：一个命名例如结构体、函数或模块等项的方式

crate 是一个二进制项或者库。crate root 是一个源文件，Rust 编译器以它为起始点，并构成你的 crate 的根模块（我们将在 “Defining Modules to Control Scope and Privacy” 一节深入解读）。包（package） 是提供一系列功能的一个或者多个 crate。一个包会包含有一个 Cargo.toml 文件，阐述如何去构建这些 crate。

我们有了一个只包含 src/main.rs 的包，意味着它只含有一个名为 my-project 的二进制 crate。如果一个包同时含有 src/main.rs 和 src/lib.rs，则它有两个 crate：一个库和一个二进制项，且名字都与包相同。通过将文件放在 src/bin 目录下，一个包可以拥有多个二进制 crate：每个 src/bin 下的文件都会被编译成一个独立的二进制 crate
```

包的规则

```
用来将路径引入作用域的 use 关键字；以及使项变为公有的 pub 关键字。我们还将讨论 as 关键字、外部包和 glob 运算符。
```

包的创建
```
cargo new --lib restaurant
```

如何在模块树中找到一个项的位置

```
1 绝对路径（absolute path）从 crate 根开始，以 crate 名或者字面值 crate 开头。
2 相对路径（relative path）从当前模块开始，以 self、super 或当前模块的标识符开头。
```
绝对路径和相对路径都后跟一个或多个由双冒号（::）分割的标识符


Rust 中默认所有项（函数、方法、结构体、枚举、模块和常量）都是私有的。父模块中的项不能使用子模块中的私有项，但是子模块中的项可以使用他们父模块中的项。这是因为子模块封装并隐藏了他们的实现详情，但是子模块可以看到他们定义的上下文。

我们还可以使用 super 开头来构建从父模块开始的相对路径


use crate::front_of_house::serving;
use crate::front_of_house::serving::server_order;
use std::io::Result as IoResult;


[dependencies]
rand = "0.5.5"


节省空间
use std::{cmp::Ordering, io};

所有的
use std::collections::*;


声明使用模块
```
xxx.rs

pub mod xxx{
    pub fn doxxx(){
        println!("do xxx");
    }
}




main.rs

mod xx_xx;
pub use crate::xx_xx::xxx;
fn main() {
    println!("Hello, world!");
    xxx::doxxx();
}

```


Rust 的多层模块遵循如下两条规则：

1 优先查找xxx.rs 文件
main.rs、lib.rs、mod.rs中的mod xxx; 默认优先查找同级目录下的 xxx.rs 文件；
其他文件yyy.rs中的mod xxx;默认优先查找同级目录的yyy目录下的 xxx.rs 文件；
2 如果 xxx.rs 不存在，则查找 xxx/mod.rs 文件，即 xxx 目录下的 mod.rs 文件。


Re-exporting

```
pub mod b;
pub use b::c::d;


我们在 main.rs 中，就可以使用 use a::d; 来调用了
```


加载外部 crate

```
extern crate xxx;


extern crate xxx;

use xxx::yyy::zzz;
```











