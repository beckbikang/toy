fn main() {

    tstruct();
}


struct User {
    username: String,
    email: String,
    sign_in_count: u64,
    active: bool,
}

#[derive(Debug)]
struct Rectangle{
    width: u32,
    height: u32, 
}
impl Rectangle {
    fn area(&self)->u32{
        self.width * self.height
    }

    fn bigger(&self, r :&Rectangle)->bool{
        self.area() > r.area()
    }

    fn square(size :u32)->Rectangle{
        Rectangle{
            width:size,
            height:size,
        }
    }
}




fn tstruct(){
    let a1:(i32,i32) = (12,34);
    println!("{},{}",a1.0,a1.1);

    let mut user1 = User{
        username: String::from("tom1"),
        email: String::from("tom@126.com"),
        sign_in_count: 1,
        active: true,
    };
    user1.username = String::from("tim");

    struct coler1(i32,i32,i32);
    let black = (0,0,0);

    let rect1 = Rectangle { width: 30, height: 50 };
    let rect2 = Rectangle { width: 40, height: 50 };

    println!("rect1 is {:?}, area={}, {}", rect1, rect1.area(), rect1.bigger(&rect2));
    let rect3 = Rectangle::square(11);
    println!("square={}",rect3.area());

}

fn build_user(email: String, username:String) -> User{
    User {
        email,
        username,
        active: true,
        sign_in_count: 1,
    }
}

