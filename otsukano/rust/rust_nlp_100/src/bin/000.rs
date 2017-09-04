#!rust run

use std::env;

fn print_rev_text(text: &str) {
    println!("{:?}", text);
    println!("{:?}", text.chars());
    println!("{:?}", text.chars().rev());
    println!("{:?}", text.chars().rev().collect::<String>());
    let rev_text = text.chars().rev().collect::<String>();
}

fn main() {
    let args: Vec<String> = env::args().collect();
    println!("{:?}", args);
    match args.get(1) {
        Some(text) => print_rev_text(text),
        None => println!("pass one argument"),
    }
}
