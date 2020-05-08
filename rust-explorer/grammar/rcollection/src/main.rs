fn main() {
    tcolloct();
}


fn tcolloct(){
    //vector
    let ve1 :Vec<i32> = Vec::new();
    let ve2  = vec![1,23];
    let mut ve3 = Vec::new();
    ve3.push(1);ve3.push(2);ve3.push(3);

    let third: &i32 = &ve3[2];
    println!("The third element is {}", third);

    match ve3.get(2) {
        Some(third) => println!("The third element is {}", third),
        None => println!("There is no third element."),
    }

    for i in &ve3{
        println!("i={}",i);
    }

    enum Spe{
        A(i32),
        B(f32),
        C(String),
    };
    let ve4 = vec![
        Spe::A(1),
        Spe::B(31.1),
        Spe::C(String::from("blue")),
    ];

    //string
    let mut s1 = String::new();
    let s2 = "hello rust";
    println!("s2={}", s2.to_string());
    let mut s3 = String::from("abc");
    s3.push("a");



}