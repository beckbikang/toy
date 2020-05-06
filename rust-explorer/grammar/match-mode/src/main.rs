fn main() {
    rust_enum();
}

#[derive(Debug)]
enum IpAddr {
    V4(u8,u8,u8,u8),
    V6(String),
}
#[derive(Debug)] // 这样可以可以立刻看到州的名称
enum UsState {
    Alabama,
    Alaska,
}
enum Base1 {
    A,
    B,
    C,
    D,
    E(UsState),
}

//add one
fn plus_one(x:Option<i32>) -> Option<i32>{
    match x{
        None => None,
        Some(i) => Some(i+1),
    }
}

fn val_match(b1:Base1) -> u8 {
    match b1{
        Base1::A => {
            println!("luck !");
            1
        },
        Base1::B => 2,
        Base1::C => 5,
        Base1::D => 10,
        Base1::E(UsState) => {
            println!("uc state {:?}", UsState);
            35
        },
    }
}

fn rust_enum(){

    let home = IpAddr::V4(127,0,0,1);
    let loopback = IpAddr::V6(String::from("::1"));
    println!("{:?}", home);
    println!("{:?}", loopback);

    let op1:Option<i32> = Some(5);
    let op2:Option<i32> = None;
    println!("{:?}", op1);
    println!("{:?}", op2);

    let base1 :Base1 = Base1::A;
    println!("{}",val_match(base1));

    let op1p = plus_one(op1);
    let op2p = plus_one(op2);
    println!("{:?}", op1p);
    println!("{:?}", op2p);

    let i = 12;
    match i {
        1=> println!("1"),
        2=> println!("2"),
        3=> println!("3"),
        _=> println!("all"),
    }

    let mut count = 1;
    if let i = 12 {
        count = count + 1;
    }else {
        println!("not matche")
    }
    println!("count={}", count);

}