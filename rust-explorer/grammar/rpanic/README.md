## 错误处理

 可恢复错误（recoverable）和 不可恢复错误（unrecoverable）


Rust 并没有异常。相反，对于可恢复错误有 Result<T, E> 值，以及 panic!，它在遇到不可恢复错误时停止程序执行。这一章会首先介绍 panic! 调用，接着会讲到如何返回 Result<T, E>。



```
enum Result<T, E> {
    Ok(T),
    Err(E),
}
```






panic场景

    1 示例，代码原型，测试

我们知道错误是可以使用Result

错误校验




