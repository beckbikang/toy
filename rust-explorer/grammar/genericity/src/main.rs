use std::fmt::Display;

fn main() {
    println!("Hello, world!");
    let p1 = Point { x: 5, y: 10.4 };
    let p2 = Point { x: "Hello", y: 'c'};
    let p3 = p1.mixup(p2);
    println!("{:?}",p3);

    let list1 = vec![1,313,56];
    let largRet = lager2(&list1);

    println!("largRet={}", largRet);

    //longest data
    let string1 = String::from("xxx00");
    let string2 = "1231";

    let result = longestStr(string1.as_str(), string2);
    println!("{}",result);

}

fn longestStr<'a>(x :&'a str, y :&'a str)->&'a str{
    if x.len() > y.len(){
        x
    }else {
        y
    }
}


fn lager2<T>(list:&[T]) -> T
    where T:PartialOrd+Copy,
{
    let mut largest = list[0];

    for &item in list.iter() {
        if item > largest {
            largest = item;
        }
    }

    largest
}


pub trait Name {
    fn genName(&self)->String{
        String::from("beginning...")
    }
}



struct Pair<T> {
    x: T,
    y: T,
}

impl<T> Pair<T> {
    fn new(x: T, y: T) -> Self {
        Self {
            x,
            y,
        }
    }
}

impl<T: Display + PartialOrd> Pair<T> {
    fn cmp_display(&self) {
        if self.x >= self.y {
            println!("The largest member is x = {}", self.x);
        } else {
            println!("The largest member is y = {}", self.y);
        }
    }
}

pub struct Twitter {
    pub username: String,
}

impl Twitter{
    fn NewTwitter(&self, s1:String)->Twitter{
        Twitter{
            username:s1,
        }
    }
}
impl Name  for Twitter {

    fn genName(&self)->String{
        format!("(Read more from {}...)", self.username)
    }
}

fn getName()->impl Name{
    Twitter{
        username:String::from("abc"),
    }
}


fn lager1<T:PartialOrd+Copy>(list:&[T])->T{
    let mut largest = list[0];

    for &item in list.iter() {
        if item > largest {
            largest = item;
        }
    }

    largest
}

#[derive(Debug)]
struct Point<T,U> {
    x:T,
    y:U,
}

impl<T,U> Point<T,U>{
    fn mixup<V, W>(self, other: Point<V, W>) -> Point<T, W> {
        Point {
            x: self.x,
            y: other.y,
        }
    }
}