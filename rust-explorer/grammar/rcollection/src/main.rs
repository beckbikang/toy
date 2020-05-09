use std::collections::HashMap;

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
    s3.push_str("a");
    println!("{}", &s2[0..4]);

    let s11 = String::from("tic");
    let s12 = String::from("tac");
    let s13 = String::from("toe");

    let s14 = s11 + "-" + &s12 + "-" + &s13;

    let s15 = format!("{}-{}-{}", s14, s12, s13);
    println!("s15-{}",s15);
    for c in s15.bytes(){
        //println!("{}",c);
    }

    //hash map
    let mut scores = HashMap::new();
    scores.insert(String::from("abcc"), 10);
    scores.insert(String::from("def"), 10);

    let teams  = vec![String::from("Blue"), String::from("Yellow")];
    let initial_scores = vec![10, 50];
    
    let scores2: HashMap<_, _> = teams.iter().zip(initial_scores.iter()).collect();

    println!("scores2:{:?}", scores2);

    let teamname = String::from("Blue");
    let ts = scores2.get(&teamname);
    println!("{:?}", ts);
    
    let mut scores22 = HashMap::new();

    scores22.insert(String::from("Blue"), 10);
    scores22.insert(String::from("Yellow"), 50);
    let score_int = 222;
    let blue1 = String::from("Blue");
    scores22.insert(blue1, score_int);

    scores22.entry(String::from("Blue")).or_insert(5555);

    for (key, value) in &scores22{
        println!("{}:{}", key, value);
    }


}