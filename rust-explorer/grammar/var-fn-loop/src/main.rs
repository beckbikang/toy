


fn main() {
    var_use();
}

//a function to use val
fn var_use(){
    let x = 12;
    //x = 13
    println!("x={}", x);
    let mut x1 = 22;
    x1 = 23;
    println!("x1={}", x1);
    
    let f1:f32 = 12.13;
    println!("f1={}", f1);

    let z = 'z';
    println!("z={}", z);

    let tup1:(i32, f64, u8) = (500,12.1,3);
    println!("{},{},{}", tup1.0,tup1.1,tup1.2);

    let months = ["January", "February", "March", "April", "May", "June", "July",
              "August", "September", "October", "November", "December"];

    let arr1 = [1,2,3,4,5];
    let arr2:[i32;5] = [1,2,3,4,5];
    println!("arr2[1]={}",arr2[1]);

    athorFunc();
    athorFunc2(12);

    let y = {
        let x = 3;
        x+1
    };
    println!("y={},five={}", y,five());

    if y > 10{
        println!("y is bigger than 10");
    }else if y > 5{
        println!("y is bigger than 5");
    }else{
        println!("y is smaller than 5");
    }

    let cond = true;
    let x22 = if cond {
        5
    }else{
        6
    };
    println!("x22={}",x22);

    let mut counter  = 1;
    let max_counter = 3;
    loop {
        counter +=1;
        println!("counter={}",counter);
        if counter >= max_counter{
            break;
        }
    }
    counter  = 1;
    while counter < max_counter {
        counter +=1;
        println!("counter={}",counter);
    }
    for a in arr1.iter() {
        println!("a={}",a);
    }

    for a in (1..3).rev() {
        println!("a={}",a);
    }


}

//a func with param
fn athorFunc2(x :i32){
    println!("func1={}",x);
}

//a func
fn athorFunc(){
    println!("func1={}",1);
}

//a func will return int
fn five()->i32{
    5
}