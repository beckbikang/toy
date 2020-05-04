fn main() {

    let s = "abc123";
    {
        let s = "abc";
    }
    println!("s={}",s);

    let mut str1 = String::from("hello");
    str1.push_str(" world!");
    println!("str1={}", str1);

    let str2 = str1;
    println!("str2={}", str2);
    //println!("str1={}", str1);//str1被移动了

    let str3 = str2.clone();
    println!("str3={}", str3);
    
    let x1 = 5;
    let x2 = x1;
    println!("{},{}", x1,x2);

    takes_ownership(str2);
    //println!("str2={}", str2);

    let str11 = test_own1();
    println!("str11={}", str11);
    let str12 = String::from("abc");
    let str13 = test_own2(str12);
    //println!("str12={}", str12);
    println!("str13={}", str13);

    let str13_len = take_ref(&str13);
    println!("str13={}, len={}", str13,str13_len);

    let mut str14 = String::from("abc");
    take_ref3(&mut str14);
    println!("str14={}", str14);

    let str14_r1 = &mut str14;
    let str14_r2 = &mut str14;
    //println!("{},{}", str14_r1,str14_r2);

    let mut str15 = String::from("abc");
    let r15_1 = &str15; 
    let r15_2 = &str15;
    //let r15_3 = &mut str15; 
    println!("{},{}", r15_1,r15_2);
    //println!("{},{},{}", r15_1,r15_2,r15_3);

    let r15_3 = &mut str15; 
    println!("{}",r15_3);


    let mut str16 = String::from("hello world!");
    let sl1 = &str16[0..5];
    let sl2 = &str16[6..11];
    println!("{},{}", sl1,sl2);

    let word1 = first_word(&str16);
    println!("{}",word1);

    let str16 = "abc ef";
    println!("{}", word1);

    println!("{}", first_word2(&str16[..]));

}

fn first_word2(s :&str)-> &str{
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }

    &s
}

fn first_word(s: &String) -> &str {
    let bytes = s.as_bytes();

    for (i, &item) in bytes.iter().enumerate() {
        if item == b' ' {
            return &s[0..i];
        }
    }

    &s[..]
}

fn take_ref3(s: &mut String) {
    s.push_str("ok ok");
}

/*
fn take_ref2(s: &String) {
    s.push_str("ok ok");
}
*/

//get len
fn take_ref(s: &String) ->usize{
    s.len()
}

fn takes_ownership(some_string: String) { // some_string 进入作用域
    println!("{}", some_string);
} 

fn test_own1() -> String{
    let str1 = String::from("abc");
    str1
}

fn test_own2(str2 :String) -> String{
    str2
}