use std::fs::File;
use std::io::ErrorKind;
use std::io;
use std::io::Read;
use std::fs;

fn main() {
    println!("Hello, world!");

    tpanicc();

}

fn tpanicc(){
    //panic!("crash and burn");
    let v = vec![1,2,3];
    //v[111];

    let f = File::open("hello12.txt");
    
    let f = match f{
        Ok(file) => file,
        Err(error) => match error.kind(){
            ErrorKind::NotFound => match File::create("hello2.txt"){
                Ok(file) => file,
                Err(error) => {
                    panic!("error{:?}", error);
                },
            },
            other_error => panic!("other error {:?}",other_error),
        },
    };

    let f = File::open("hello3.txt").unwrap_or_else(|error| {
        if error.kind() == ErrorKind::NotFound {
            File::create("hello3.txt").unwrap_or_else(|error| {
                panic!("Problem creating the file: {:?}", error);
            })
        } else {
            panic!("Problem opening the file: {:?}", error);
        }
    });

    let f = File::open("hello.txt").unwrap();

    let f = File::open("hello.txt").expect("Failed to open hello.txt");

    /*
    let f = match f {
        Ok(file) => file,
        Err(error) => {
            panic!("error{:?}", error);
        },
    };
    */
    
    let str12 = read_username_from_file().unwrap_or_else(
        |error|{
            panic!("Problem creating the file: {:?}", error);
    });
    println!("str12:{}",str12);

    let ret1 = read_username_from_file2();
    let ret1 = ret1.unwrap();
    println!("ret1 {}", ret1);


}

fn read_username_from_file4() -> Result<String, io::Error> {
    fs::read_to_string("hello.txt")
}


fn read_username_from_file3() -> Result<String, io::Error> {
    let mut s = String::new();

    File::open("hello.txt")?.read_to_string(&mut s)?;

    Ok(s)
}

fn read_username_from_file2() -> Result<String, io::Error> {
/**
Result 值之后的 ? 被定义为与示例 9-6 中定义的处理 Result 值的 
match 表达式有着完全相同的工作方式。如果 Result 的值是 Ok，
这个表达式将会返回 Ok 中的值而程序将继续执行。
如果值是 Err，Err 中的值将作为整个函数的返回值，
就好像使用了 return 关键字一样，这样错误值就被传播给了调用者。
**/
    let mut f = File::open("hello.txt")?;
    let mut s = String::new();
    f.read_to_string(&mut s)?;
    Ok(s)
}


fn read_username_from_file() -> Result<String, io::Error> {
    let f = File::open("hello.txt");

    let mut f = match f {
        Ok(file) => file,
        Err(e) => return Err(e),
    };

    let mut s = String::new();

    match f.read_to_string(&mut s) {
        Ok(_) => Ok(s),
        Err(e) => Err(e),
    }
}