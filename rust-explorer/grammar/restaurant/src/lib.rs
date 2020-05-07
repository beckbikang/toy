use std::io::Result as IoResult;

#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }
}

mod front_of_house {

    fn test_super(){}

    pub mod hosting {
        pub fn add_to_waitlist() {
            super::test_super();
        }

        fn seat_at_table() {}
    }

    pub mod serving {
        fn take_order() {}

        pub fn server_order() {}

        fn take_payment() {}
    }
}

mod back_of_home{
    #[derive(Debug)]
    pub struct Shape{
        pub name:String,
        color:String,
    }
    impl Shape {
        pub fn aa(n1:&str)->Shape{
            Shape{
                name: String::from(n1),
                color: String::from("color"),
            }
        }
    }
}

mod back_of_house {
    pub enum Appetizer {
        Soup,
        Salad,
    }
}

use crate::front_of_house::serving;
use crate::front_of_house::serving::server_order;

pub fn eat_at_restaurant(){
    //absolute 
    crate::front_of_house::hosting::add_to_waitlist();
    //relate
    front_of_house::serving::server_order();

    serving::server_order();
    server_order();

    let s1 = back_of_house::Appetizer::Soup;
    let s2 = back_of_house::Appetizer::Salad;



    let sp = back_of_home::Shape::aa("a123");
    println!("{:?}", sp);
}