#[cfg(test)]
mod tests {
    #[test]
    fn it_works() {
        assert_eq!(2 + 2, 4);
    }

    #[test]
    fn test_other(){
        println!("tt");
    }
}

#[derive(Debug)]
struct point{
    x:i32,
    y:i32,
}
