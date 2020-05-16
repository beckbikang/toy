use std::env;

fn main() {
   simple_io();
}

fn simple_io(){
    let args: Vec<String> = env::args().collect();

    if args.len() < 3{
        panic!("length < 3");
    }
    let query = &args[1];
    let filename = &args[2];

    println!("query:{}", query);

    println!("filename:{}", filename);

}
